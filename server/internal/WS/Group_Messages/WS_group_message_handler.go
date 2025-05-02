package ws

import (
	"log"
	"net/http"

	"github.com/Aritra640/ConnectSphere/server/internal/utils"
	"github.com/google/uuid"
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

  gid := c.QueryParam("id")
  guid,_ := uuid.Parse(gid)

  ws,err := upgrader.Upgrade(c.Response() , c.Request() , c.Response().Header())
  if err != nil {
    log.Println("Error: cannot upgrade to websocket in GroupMessageHandler: " , err)
    return c.JSON(http.StatusConflict , "Socket refused")
  }
  defer ws.Close()

  gcs.AddUserInGroup(guid , ws)

  for {
    _,msg,err := ws.ReadMessage()
    if err != nil {
      //TODO: remove the socket 
      log.Println("Error: socket disconnected in FroupMessageHandler: " , err)
      return c.JSON(http.StatusConflict , "Websocket terminated")
    }

    req,err := utils.GetRequestGroup_JSON(msg)
    if err != nil {
      log.Println("Error: cannot get request in GroupMessageHandler: " , err)
      ws.WriteMessage(websocket.TextMessage , []byte("Cannot "))
    }

    if req.RequestType == utils.Join {
      // add a request to join the group
      //if group is not restriceted then add user to Group 

      continue
    }
    
    //check if user is a part of the group 
    //send the message to the group

    
  }
}
