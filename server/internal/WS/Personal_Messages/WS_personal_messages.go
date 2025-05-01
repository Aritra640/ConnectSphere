package ws

import (
	"encoding/json"
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
		//should not be all true in prod
		return true
	},
}

// PersonalMessagesHandler is the ws handler for personal messages
func (pcs *PersonalChatService) PersonalMessagesHandler(c echo.Context) error {

  pid := c.QueryParam("pid")
  puid,_ := uuid.Parse(pid)

	ws, err := upgrader.Upgrade(c.Response(), c.Request(), c.Response().Header())
	if err != nil {
		log.Println("Error: cannot upgrade to websocket: ", err)
		return c.JSON(http.StatusInternalServerError, "something went wrong")
	}
	defer ws.Close()

	pcs.WS_store.AddClient(puid, ws)

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			//remove client
			pcs.WS_store.DeleteConn(puid)
			log.Println("Error: client disconnected in PersonalMessagesHandler: ", err)
			return c.JSON(http.StatusConflict, "WS connection terminated")
		}

		res, err := utils.GetPersonalPersonal_JSON(msg)
		if err != nil {
			log.Println("Error: data invalid: ", err)
			str, _ := StringReturn(res.SenderID, uuid.New(), "error", "Invalid Format")
			ws.WriteMessage(websocket.TextMessage, []byte(str))
			continue
		}

    if res.TypeMsg == utils.TypeMessage(Join) {
      
    }

		//Create Personal Message
		cidCh := make(chan uuid.UUID)
		errCh := make(chan error)
		go func() {
			cid, err := PersonalMessageSetup.CreatePersonalMassage(c.Request().Context(), CreatePersonalMassageInput{
				UserID:     int32(res.SenderID),
				SenderID:   int32(res.SenderID),
				ReceiverID: int32(res.ReceiverID),
				Content:    res.Content,
				Type:       string(res.Type),
			})

			if err != nil {
				errCh <- err
			}
			cidCh <- cid
		}()

		select {
		case cid := <-cidCh:
			sendStr, _ := StringReturn(res.SenderID, cid, "message", res.Content)
			//send the same message to both
			pcs.WS_store.SendMesssage(sendStr , ws , puid , Chat)

		case err = <-errCh:
			log.Println("Error: cannot create personal message: ", err)
			errStr, _ := StringReturn(res.SenderID, uuid.New(), "error", "Something went wrong")
			ws.WriteMessage(websocket.TextMessage, []byte(errStr))
		}

	}
}

func StringReturn(userID int, chatID uuid.UUID, typeMsg string, content string) (string, error) {

	data := map[string]interface{}{
		"user":    userID,
		"type":    typeMsg,
		"content": content,
		"chatID":  chatID,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println("Error: cannot Marshal data: ", err)
		return "", err
	}

	return string(jsonData), nil
}

func StringReturnHandler(c echo.Context) error {

	str, _ := StringReturn(123, uuid.New(), "test", "this is an example return handler")
	return c.JSON(http.StatusOK, map[string]string{
		"message": str,
	})
}
