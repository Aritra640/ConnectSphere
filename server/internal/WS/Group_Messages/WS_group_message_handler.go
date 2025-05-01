package ws

import (
	"log"
	"net/http"

	"github.com/Aritra640/ConnectSphere/server/internal/utils"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var upgrader = websocket.Upgrader{
  ReadBufferSize: 1024,
  WriteBufferSize: 1024,
  CheckOrigin: func(r *http.Request) bool {
    return true
  },
}

func (gcs *GroupChatService) GroupMessageHandler(c echo.Context) error {

  ws,err := upgrader.Upgrade(c.Response() , c.Request() , c.Response().Header())
  if err != nil {
    log.Println("Error: cannot upgrade to websocket in GroupMessageHandler: " , err)
    return c.JSON(http.StatusConflict , "Socket refused")
  }
  defer ws.Close()

  for {
    _,msg,err := ws.ReadMessage()
    if err != nil {
      log.Println("Error: socket disconnected in FroupMessageHandler: " , err)
      return c.JSON(http.StatusConflict , "Websocket terminated")
    }

    res,err := utils.GetResponseGroup_JSON(msg); if err != nil {
      log.Println("Error: cannot unmarshall json")
      return c.JSON(http.StatusBadRequest , "Invalid Request")
    }


  }
}
