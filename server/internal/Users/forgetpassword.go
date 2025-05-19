package users

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/Aritra640/ConnectSphere/server/Database/db"
	mail "github.com/Aritra640/ConnectSphere/server/internal/Mail-Server"
	"github.com/Aritra640/ConnectSphere/server/internal/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type ForgetPasswordParam struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}

// Send a mail to the user with a otp
func (u *UserService) ForgetPasswordHandler(c echo.Context) error {

	var req ForgetPasswordParam
	if err := c.Bind(&req); err != nil {
		log.Println("Error: ForgetPasswordHandler -> invalid request: ", err)
		return c.JSON(http.StatusBadRequest, "Invalid Request")
	}
	if err := c.Validate(&req); err != nil {
		log.Println("Error: ForgetPasswordHandler -> validation failed: ", err)
		return c.JSON(http.StatusBadRequest, "Invalid Type")
	}

	foundCh := make(chan bool)
	go func() {
		_, err := u.Queries.GetUserByUsername(c.Request().Context(), req.Username)
		if err != nil {
			foundCh <- false
		}
		foundCh <- true
	}()

	jwtCh := make(chan string)
	defer close(jwtCh)

	go func() {
		jwt, _ := generateJWT(req.Email)
		jwtCh <- jwt
	}()

	jwttoken := <-jwtCh
	if jwttoken == "" {
		log.Println("Error: cannot generate jwt in ForgetPasswordHandler")
		return c.JSON(http.StatusInternalServerError, "Something went wrong")
	}
	resentlink := u.Resentlink + "?token=" + jwttoken
	wg := &sync.WaitGroup{}
	retCh := make(chan error, 1)
	defer close(retCh)

	ctx, ctxClose := context.WithCancel(c.Request().Context())
	defer ctxClose()

	wg.Add(1)
	go mail.MailSetup.ForgetPasswordMail(ctx, req.Username, req.Email, resentlink, wg, retCh)
	//Lets not do the error in mail now

	found := <-foundCh
	if !found {
		ctxClose()
		log.Println("Error: user not found in database in ForgetPasswordHandler with email id: ", req.Email)
		return c.JSON(http.StatusBadRequest, "Invalid Request")
	}

	wg.Wait()
	err := <-retCh
	if err != nil {
		log.Println("Error: sending resent password mail: ", err)
		return c.JSON(http.StatusConflict, "Failed to send email!")
	}

	log.Println("Resent mail successfully sent")
	return c.JSON(http.StatusOK, "Resent mail sent!")
}

//Send a otp and verify
//if verified then proceed

type MyCustomClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

// generate jwt token
func generateJWT(email string) (string, error) {

	claims := MyCustomClaims{
		email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(6 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, err := token.SignedString(config.App.JWT)
	if err != nil {
		log.Println("Error: could not generate jwt in ForgetPasswordHandler: ", err)
		return "", err
	}

	return ss, nil
}

func verifyJWT(tokenString string) (string, error) {

	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method : %v", token.Header["alg"])
		}
		return config.App.JWT, nil

	})
	if err != nil {

		log.Println("Error: error in parsing token: ", err)
		return "", err
	}

	claims, ok := token.Claims.(*MyCustomClaims)
	if !ok || !token.Valid {
		log.Println("Invalid token in ForgetPasswordHandler checker: ", err)
		return "", err
	}

	return claims.Email, nil
}

type ChangePasswordReq struct {
	NewPassword string `json:"new_password" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Token       string `json:"token" validate:"required"`
}

// Change users password to a new password
func (u *UserService) ChangeUserPasswordHandler(c echo.Context) error {
	
	var req ChangePasswordReq
	if err := c.Bind(&req); err != nil {
		log.Println("Error: invalid request in ChangeUserPasswordHandler: " , err)
		return c.JSON(http.StatusBadRequest , "Invalid Request")
	}

	if err := c.Validate(&req); err != nil {
		log.Println("Error: validation failed in ChangeUserPasswordHandler: " , err)
		return c.JSON(http.StatusBadRequest , "Invalid Type")
	}

	verrChan := make(chan error)
	vemailChan := make(chan string)
	defer close(verrChan)
	defer close(vemailChan)

	go func() {

		emailstr,err := verifyJWT(req.Token)
		if err != nil {
			verrChan <- err
		}
		vemailChan <- emailstr
	}()

	select{
	case <-c.Request().Context().Done(): 
		log.Println("Error: ChangeUserPasswordHandler timed out")
		return c.JSON(http.StatusInternalServerError , "timed out")

	case err := <-verrChan:
		log.Println("Error: Failed to verify token in ChangeUserPasswordHandler: " ,err)
		return c.JSON(http.StatusConflict , "Process failed!")

	case email := <-vemailChan: 
		if email != req.Email {
			log.Println("Error: email id not matched with resent link with email id: " , req.Email)
			return c.JSON(http.StatusBadRequest , "Invalid Request")
		}
	}

	errCh := make(chan error)
	defer close(errCh)
	go func() {
		err := u.Queries.UpdateUserPasswordByEmail(c.Request().Context() , db.UpdateUserPasswordByEmailParams{
			Email: req.Email,
			PasswordHashed: req.NewPassword,
		})
		errCh <- err
	}()

	err := <-errCh
	
	if err != nil {
		log.Println("Error: Update user password failed: " , err)
		return c.JSON(http.StatusInternalServerError , "Something went wrong")
	}

	return c.JSON(http.StatusOK , "Password changed! login with your new password")
}
