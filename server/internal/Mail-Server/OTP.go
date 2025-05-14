package mail

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

//GenerateOTP generate otp concurrently and check if it is not repeated
func (ms *MailService) GenerateOTP(ctx context.Context) (string, error) {

	otpFound := false
	var otp string

	for !otpFound {

		randomBytes := make([]byte, 6)
		_, err := rand.Read(randomBytes)
		if err != nil {
			log.Println("Error: could not generate otp: ", err)
			return "", err
		}

		otp = base64.RawURLEncoding.EncodeToString(randomBytes)[:6]

		foundCh := make(chan bool)
		defer close(foundCh)

		go func() {

			found := ms.Ots.CheckOTP(ctx, otp)
			foundCh <- found
		}()

		found := <-foundCh
		if !found {
			otpFound = true
			break
		}
	}

	return otp, nil
}

type GetOTPHandlerParam struct {
	Email string `json:"email" validate:"required,email"`
}

//GetOTPHandler : HTTP handler to get otp
func (ms *MailService) GetOTPHandler(c echo.Context) error {

	var req GetOTPHandlerParam
	err := c.Bind(&req)
	if err != nil {
		log.Println("Error: Cannot get request in GetOTPHandler")
		return c.JSON(http.StatusBadRequest, "Invalid Request")
	}

	if err := c.Validate(&req); err != nil {
		log.Println("Error: Validation failed in GetOTPHandler")
		return c.JSON(http.StatusBadRequest, "Invalid Type")
	}

	otpCh := make(chan string)
	errCh := make(chan error)
	defer close(otpCh)
	defer close(errCh)

	go func() {
		otp, err := ms.GenerateOTP(c.Request().Context())
		if err != nil {
			errCh <- err

		} else {
			otpCh <- otp
		}
	}()

	select {
	case <-c.Request().Context().Done():
		log.Println("OTP generator timed out")
		return c.JSON(http.StatusConflict, "Timed out")

	case otp := <-otpCh:
		errCh := make(chan error)
		defer close(errCh)
		go func() {

			ms.Ots.AddOtp(c.Request().Context(), req.Email, otp)
		}()

		err := <-errCh
		if err != nil {
			log.Println("Error: Failed to add otp : ", err)
			return c.JSON(http.StatusInternalServerError, "Something Went Wrong")
		} else {

			log.Println("OTP generation successfull!")
			return c.JSON(http.StatusOK, map[string]interface{}{
				"otp": otp,
			})
		}
  
  case <-c.Request().Context().Done():
    log.Println("Error: OTPHandler timed out")
    return c.JSON(http.StatusInternalServerError , "Timed Out")
	}
}
