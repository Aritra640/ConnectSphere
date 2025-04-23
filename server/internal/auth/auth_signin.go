package auth

import (
	"log"
	"net/http"
	"time"

	"github.com/Aritra640/ConnectSphere/server/Database/db"
	"github.com/Aritra640/ConnectSphere/server/internal/utils"
	"github.com/labstack/echo/v4"
)

func (as *AuthService) SigninHandler(c echo.Context) error {
  var req SigninRequest 

  //Bind and validate input 
  if err := c.Bind(&req); err != nil {
    return c.JSON(404 , "Invalid Request")
  }

  if err := c.Validate(&req); err != nil {
    return c.JSON(http.StatusBadRequest , "Validation failed")
  }

  errChan := make(chan error)
  userChan := make(chan db.User)
  //Fetch User
  go func() {
    user,err := as.Queries.GetUserbyEmail(c.Request().Context() , req.Email)
    if err != nil {
      errChan <- err
    }
    userChan <- user
  }()
  
  var user db.User
  select{
  case err := <-errChan: 
    log.Println("Invalid Email in sign in :" , err)
    return c.JSON(http.StatusUnauthorized , "Invalid Email")
  case user = <-userChan : 
    log.Println("User found in sign in")
  }

  //Verify password 
  verifyCh := make(chan bool)
  go func() {
    match := utils.VerifyHashedPassword(req.Password , user.PasswordHashed)
    verifyCh <- match
  }()
  found := <-verifyCh
  if !found {
    log.Println("Could not found or match password")
    return c.JSON(http.StatusUnauthorized , "Invalid Password")
  }

  //Generate jwt token 
  jwtCh := make(chan string)
  errCh := make(chan error)
  go func() {
    token,err := CreateAuthToken(int(user.ID))
    if err != nil {
      log.Println("Error in creating jwt token: " , err)
      errCh <- err
    }
    jwtCh <- token
  }()
  
  var jwt string
  select {
  case jwt = <-jwtCh:
  case err := <-errCh: 
    log.Println("Error : could not create token: " , err)
    return c.JSON(http.StatusInternalServerError , "something went wrong")
  }

  //Genrate refresh token 
  refreshToken,err := as.Rts.CreateRefreshToken(c.Request().Context() , int(user.ID) , 7*24*time.Hour) 
  if err != nil {
    log.Println("Refresh token creation: " , err )
    return c.JSON(http.StatusInternalServerError , "could not generate refresh token")
  }

  //Return token 
  return c.JSON(http.StatusOK , map[string]interface{}{
    "access_token": jwt,
    "refresh_token": refreshToken,
    "user": map[string]interface{}{
      "id": user.ID,
      "username": user.Username,
      "email": user.Email,
    },
  })
 }
