package controllers

import (
	"github.com/labstack/echo/v4"
)

func RoutesSetup(e *echo.Echo) {

  e.GET("/hello" , func(c echo.Context) error {

    return c.JSON(200 , "hi hello there")
  })

  e.POST("/signup"  )
}
