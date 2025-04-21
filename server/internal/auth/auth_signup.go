package auth

import (
	"log"
	"net/http"

	"github.com/Aritra640/ConnectSphere/server/internal/utils"
	"github.com/labstack/echo/v4"
)

func (as *AuthService) SignupHandler(c echo.Context) error {
  
  var req SignupRequest
  //TODO: get proper error code and message from echo.HTTPError
  err := c.Bind(&req); if err != nil {
    log.Println("Invalid signup request")
    return c.JSON(http.StatusBadRequest , "Invalid Request")
  }

  err = c.Validate(&req); if err != nil {
    log.Println("Signup request could not be validated")
    return c.JSON(http.StatusBadRequest , "validation failed")
  }

  hashCh, errCh := make(chan string), make(chan error)
  go func() {
    hash_password , error := utils.HashPassword(req.Password)
  }()
}
