package controllers

import (
	ws "github.com/Aritra640/ConnectSphere/server/internal/WS/test_chat_room"
	"github.com/labstack/echo/v4"
)

func RoutesSetup(e *echo.Echo) {

  e.GET("/hello" , func(c echo.Context) error {

    return c.JSON(200 , "hi hello there")
  })


}
