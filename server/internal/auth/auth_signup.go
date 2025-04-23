package auth

import (
	"log"
	"net/http"

	"github.com/Aritra640/ConnectSphere/server/Database/db"
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
    hash_password , err_hash := utils.HashPassword(req.Password)
    if err_hash != nil {
      errCh <- err_hash
    }

    hashCh <- hash_password
  }()

  var hashed_password string 
  select {
  case hashed_password = <- hashCh : 
  case err := <-errCh : 
    log.Println("Error hashing password: " , err)
    return c.JSON(http.StatusInternalServerError , "internal error")
  }

  //save the user in the database 
  _,err = as.Queries.AddUser(c.Request().Context() , db.AddUserParams{
    Username: req.Username,
    PasswordHashed: hashed_password,
    Email: req.Email,
  })
  if err != nil {
    log.Println("Error creating user: " , err)
    return c.JSON(http.StatusInternalServerError , "Could not create user")
  }

  return c.JSON(http.StatusCreated , map[string]interface{}{
    "message": "User created successfully",
  })
}
