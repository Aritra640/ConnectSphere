package auth

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/Aritra640/ConnectSphere/server/Database/db"
	mail "github.com/Aritra640/ConnectSphere/server/internal/Mail-Server"
	"github.com/Aritra640/ConnectSphere/server/internal/cachestore"
	"github.com/Aritra640/ConnectSphere/server/internal/utils"
	"github.com/labstack/echo/v4"
)

func (as *AuthService) SignupHandler(c echo.Context) error {

	var req SignupRequest
	err := c.Bind(&req)
	if err != nil {
		log.Println("Invalid signup request")
		return c.JSON(http.StatusBadRequest, "Invalid Request")
	}

	err = c.Validate(&req)
	if err != nil {
		log.Println("Signup request could not be validated")
		return c.JSON(http.StatusBadRequest, "validation failed")
	}

	hashCh, errCh := make(chan string), make(chan error)
	go func() {
		hash_password, err_hash := utils.HashPassword(req.Password)
		if err_hash != nil {
			errCh <- err_hash
		}

		hashCh <- hash_password
	}()

	var hashed_password string
	select {
	case hashed_password = <-hashCh:
	case err := <-errCh:
		log.Println("Error hashing password: ", err)
		return c.JSON(http.StatusInternalServerError, "internal error")
	}

	//save the user in the database
	// _,err = as.Queries.AddUser(c.Request().Context() , db.AddUserParams{
	//   Username: req.Username,
	//   PasswordHashed: hashed_password,
	//   Email: req.Email,
	// })
	// if err != nil {
	//   log.Println("Error creating user: " , err)
	//   return c.JSON(http.StatusInternalServerError , "Could not create user")
	// }
	//
	// return c.JSON(http.StatusCreated , map[string]interface{}{
	//   "message": "User created successfully",
	// })

	//Add user to cache store
	cachestore.CacheService.AddUnverifiedUser(c.Request().Context(), cachestore.UnverifiedUserStore{
		UserName:       req.Username,
		Email:          req.Email,
		HashedPassword: hashed_password,
		TimeStamp:      time.Now(),
	})

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":   "User signup process -- otp verification is required",
		"user_name": req.Username,
		"email":     req.Email,
	})
}

type GetOTPHandlerParam struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}

// Send otp through mail for verification
func (as *AuthService) GetOTPHandler(c echo.Context) error {

	var req GetOTPHandlerParam
	err := c.Bind(&req)
	if err != nil {
		log.Println("Error: GetOTPHandler -> Invalid request: ", err)
		return c.JSON(http.StatusBadRequest, "Invalid Request")
	}

	err = c.Validate(&req)
	if err != nil {
		log.Println("Error: Validation failed in GetOTPHandler: ", err)
		return c.JSON(http.StatusBadRequest, "Invalid Type")
	}

	otpChan := make(chan string)
	defer close(otpChan)

	go func() {
		otp := cachestore.CacheService.GetNewOTP(c.Request().Context(), req.Email)
		otpChan <- otp
	}()

	otp := <-otpChan
	errChan := make(chan error)
	defer close(errChan)

	go func() {

		err := mail.MailSetup.OTPmail(c.Request().Context(), req.Username, req.Email, otp)
		errChan <- err
	}()

	err = <-errChan
	if err != nil {
		log.Println("Error: Failed to send otp mail : ", err)
		return c.JSON(http.StatusInternalServerError, "Something went wrong")
	}

	log.Println("1 OTP mail sent!")
	return c.JSON(http.StatusOK, "OTP has been sent in mail!")
}

type VerifyOTPHandlerParam struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Otp      string `json:"otp" validate:"required"`
}

// Verify otp sent through and if otp matched then add user to the database
func (as *AuthService) VerifyOTPHandler(c echo.Context) error {

	var req VerifyOTPHandlerParam
	err := c.Bind(&req)

	if err != nil {
		log.Println("Error: GetOTPHandler -> Invalid request: ", err)
		return c.JSON(http.StatusBadRequest, "Invalid Request")
	}

	err = c.Validate(&req)
	if err != nil {
		log.Println("Error: Validation failed in GetOTPHandler: ", err)
		return c.JSON(http.StatusBadRequest, "Invalid Type")
	}

	foundCh := make(chan bool)
	defer close(foundCh)
	userdataCh := make(chan cachestore.UnverifiedUserStore)
	defer close(userdataCh)
	found2Ch := make(chan bool)

	go func() {

		found := cachestore.CacheService.VerifyOTP(c.Request().Context(), req.Email, req.Otp)
		foundCh <- found
	}()

	go func() {

		userfound, userdata := cachestore.CacheService.GetUnverifiedUser(c.Request().Context(), req.Email)
		if !userfound {
			found2Ch <- false
		}
		userdataCh <- userdata
	}()

	found := <-foundCh
	if !found {

		log.Println("Error: Could not verify otp with mailid: ", req.Email)
		return c.JSON(200, map[string]interface{}{
			"match":   "false",
			"message": "Otp did not match",
		})
	}

	var hashed_password string
	select {
	case userdata := <-userdataCh:
		hashed_password = userdata.HashedPassword
	case <-found2Ch:
		log.Println("Error: cannot get userdata in cache store with email id: ", req.Email)
		return c.JSON(404, "Something went wrong")
	}
  wg := &sync.WaitGroup{}
  wg.Add(1)

  go mail.MailSetup.WelcomeMail(c.Request().Context(), req.Username , req.Email , wg)

	log.Println("OTP matched for user with mail id: ", req.Email)
	log.Println("Adding User to database with username: ", req.Username)

	errchan := make(chan error)
	defer close(errchan)

	go func() {

		_, err = as.Queries.AddUser(c.Request().Context(), db.AddUserParams{
			Username:       req.Username,
			PasswordHashed: hashed_password,
			Email:          req.Email,
		})
		errchan <- err
	}()
	if err != nil {
		log.Println("Error creating user: ", err)
		return c.JSON(http.StatusInternalServerError, "Could not create user")
	}
  
  wg.Wait()
	return c.JSON(200, map[string]interface{}{
		"match":  "true",
		"mesage": "User successfully signed up!",
	})
}
