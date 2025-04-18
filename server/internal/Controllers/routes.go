package controllers

import (
	ws "github.com/Aritra640/ConnectSphere/server/internal/WS/test_chat_room"
	"github.com/labstack/echo/v4"
)

func RoutesSetup(e *echo.Echo) {

  e.GET("/hello" , func(c echo.Context) error {

    return c.JSON(200 , "hi hello there")
  })


  e.Any("/chat_room_test" , ws.TestChatRoom)
  e.GET("/join_group")
  e.Any("/chat_to_user")
  e.Any("/chat_to_group")
  e.DELETE("/leave_group")
  e.POST("/create_group")
  e.DELETE("/delete_a_group")
  e.PUT("/block_a_user")
  e.DELETE("/delete_account")
  e.PUT("/change_user_name")
  e.PUT("/change_user_pic")
  e.PUT("/change_group_pic")
  e.PUT("/change_group_info")
  e.PUT("/change_group_perm")
}
