// TODO: rewrite it
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
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (gcs *GroupChatService) GroupMessageHandler(c echo.Context) error {

	gid := c.QueryParam("id")
	guid, _ := uuid.Parse(gid)

	ws, err := upgrader.Upgrade(c.Response(), c.Request(), c.Response().Header())
	if err != nil {
		log.Println("Error: cannot upgrade to websocket in GroupMessageHandler: ", err)
		return c.JSON(http.StatusConflict, "Socket refused")
	}
	defer ws.Close()

	gcs.AddUserInGroup(guid, ws)

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			gcs.DeleteClientFromAllGroups(c.Request().Context(), ws)
			log.Println("Error: socket disconnected in FroupMessageHandler: ", err)
			return c.JSON(http.StatusConflict, "Websocket terminated")
		}

		req, err := utils.GetRequestGroup_JSON(msg)
		if err != nil {
			log.Println("Error: cannot get request in GroupMessageHandler: ", err)
			ws.WriteMessage(websocket.TextMessage, []byte("Cannot "))
		}

		if req.RequestType == utils.Join {
			// add a request to join the group
			//if group is not restriceted then add user to Group
			err := gcs.RequestJoinGroup(c.Request().Context(), req.UserID, guid)
			if err != nil {
				ws.WriteMessage(websocket.TextMessage, []byte("Something went wrong with websocket connection"))
				gcs.DeleteClientFromAllGroups(c.Request().Context(), ws)

				log.Println("Error: Failed to request join group in WS group handler: ", err)
				return c.JSON(http.StatusInternalServerError, "Invalid Request")
			}
			continue
		}

		//check if user is a part of the group
		//send the message to the group

		found, err := gcs.IsGroupMember(c.Request().Context(), guid, int32(req.UserID))
		if err != nil {
			log.Println("Error: error in IsGroupMember in ws group handler: ", err)

			return c.JSON(http.StatusInternalServerError, "Something went wrong")
		}

		if !found {
			log.Printf("Error: user with userid: %v not a member og group with id: %v", req.UserID, guid)
			return c.JSON(http.StatusBadRequest, "Invalid Request")
		}
	}
}
