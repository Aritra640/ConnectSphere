package controllers

import (
	"github.com/Aritra640/ConnectSphere/server/internal/auth"
	"github.com/Aritra640/ConnectSphere/server/internal/config"
	"github.com/labstack/echo/v4"
)

func RoutesSetupV1(e *echo.Echo) {

  apiv1 := e.Group("/api/v1")

  apiv1.GET("/template_signup" , func(c echo.Context) error {

    return c.JSON(200, auth.SignupRequest{
      Username: "testUser",
      Password: "1333",
      Email: "test@test.com",
    })
  })
  apiv1.GET("/template_signin" , func(c echo.Context) error {

    return c.JSON(200, auth.SigninRequest{
      Email: "test@test.com",
      Password: "12333",
    })
  })

  apiv1.POST("/signup" , auth.AuthSetup.SignupHandler)
  apiv1.POST("/signin" , auth.AuthSetup.SigninHandler)

  apiv1.GET("/protected" , ProtectedHandler , auth.AuthSetup.AuthMiddleware)

  apiv1.GET("/get_personal_chat_history" , config.App.PCS.GetPersonalChatHistoryHandler)
}
