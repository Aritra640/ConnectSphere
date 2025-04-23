package controllers

import (
	"log"
	"strconv"

	"github.com/Aritra640/ConnectSphere/server/internal/config"
	"github.com/labstack/echo/v4"
)

func ProtectedHandler(c echo.Context) error {

  userID := c.Request().Header.Get("UserID")
  log.Println("User id: " , userID)
  user_id ,err := strconv.Atoi(userID)
  if err != nil {
    log.Println("user cannot be converted to suitable format")
    return c.JSON(500 , "something went wrong")
  }
  user,err := config.App.QueryObj.GetUserByID(c.Request().Context() , int32(user_id))
  if err != nil {
    log.Println("Could not find user wuth uid: " , user_id)
    return c.JSON(500 , "something went wrong")
  }


  return c.JSON(200 , map[string]interface{}{
    "Username": user.Username,
    "Email": user.Email,
  })
}
