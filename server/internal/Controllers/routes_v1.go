package controllers

import (
	wsp "github.com/Aritra640/ConnectSphere/server/internal/WS/Personal_Messages"
	"github.com/Aritra640/ConnectSphere/server/internal/auth"
	"github.com/Aritra640/ConnectSphere/server/internal/config"
	"github.com/labstack/echo/v4"
)

func RoutesSetupV1(e *echo.Echo) {

	apiv1 := e.Group("/api/v1")

	apiv1.GET("/template_signup", func(c echo.Context) error {

		return c.JSON(200, auth.SignupRequest{
			Username: "testUser",
			Password: "1333",
			Email:    "test@test.com",
		})
	})
	apiv1.GET("/template_signin", func(c echo.Context) error {

		return c.JSON(200, auth.SigninRequest{
			Email:    "test@test.com",
			Password: "12333",
		})
	})

	apiv1.POST("/signup", auth.AuthSetup.SignupHandler)
	apiv1.POST("/signin", auth.AuthSetup.SigninHandler)

	apiv1.GET("/protected", ProtectedHandler, auth.AuthSetup.AuthMiddleware)

  //Personal Messages Routes 
	apiv1.GET("/auth/get_personal_chat_history", config.App.PCS.GetPersonalChatHistoryHandler)

	apiv1.GET("/auth/create_NewPmsg", config.App.PCS.CreateNewPersonalMessageHandler, auth.AuthSetup.AuthMiddleware)
	apiv1.Any("/ws/join_PMsg", config.App.PCS.PersonalMessagesHandler)

	apiv1.GET("/StringPersonalMessageEx", wsp.StringReturnHandler)
	apiv1.PUT("/auth/edit_personal_msg", config.App.PCS.EditPersonalMessageHandler, auth.AuthSetup.AuthMiddleware)
	apiv1.DELETE("/auth/delete_personal_msg", config.App.PCS.DeletePersonalMassageHandler, auth.AuthSetup.AuthMiddleware)
	apiv1.POST("/auth/mark_personal_msg_as_seen", config.App.PCS.PersonalMessageMarkAsSeenHandler, auth.AuthSetup.AuthMiddleware)

	apiv1.GET("/auth/get_friends_personal", config.App.PCS.GetAllFriendsHandler, auth.AuthSetup.AuthMiddleware)
	apiv1.GET("/sample_get_friends_personal", wsp.SampleGetAllFriendsHandler)

  // ---- 


}
