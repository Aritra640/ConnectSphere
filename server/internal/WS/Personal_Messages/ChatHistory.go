//TODO: Fix uuid 
package ws

import (
	"log"
	"net/http"

	"github.com/Aritra640/ConnectSphere/server/Database/db"
	"github.com/labstack/echo/v4"
)

type ChatHistoryRequest struct {
	User1 int `json:"user_1"`
	User2 int `json:"user_2"`
}

//GetPersonalChatHistoryHandler get all chat history of the two users
func (pcs *PersonalChatService) GetPersonalChatHistoryHandler(c echo.Context) error {

	reqCh := make(chan ChatHistoryRequest)
	errCh := make(chan error)
	go func() {
		var req ChatHistoryRequest
		if err := c.Bind(&req); err != nil {
			log.Println("Error: could not get chat history request")
			errCh <- err
		}
		reqCh <- req
	}()

	var req ChatHistoryRequest
	select {
	case req = <-reqCh:
	case err := <-errCh:
		log.Println("could not get request in chat history handler: ", err)
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}
	//TODO: make this concurrent
	messages, err := pcs.Queries.GetMessagesBetweenTwoUsers(c.Request().Context(), db.GetMessagesBetweenTwoUsersParams{
		SenderID:   int32(req.User1),
		ReceiverID: int32(req.User2),
	})
	if err != nil {
		log.Println("Error failed to get any message : ", err)
		return c.JSON(http.StatusConflict, "Something went wrong")
	}

	returnCh := make(chan []ChatHistoryReturn)
	go func() {
		returnReq := make([]ChatHistoryReturn, 0)
		for _, message := range messages {
			returnReq = append(returnReq, ChatHistoryReturn{
				Content: message.Content,
				Type:    message.Type.(string),
				IsSeen:  message.IsSeen,
				UserID:  int(message.UserID.Int32),
				ChatID:  message.ID.String(),
			})
		}

		returnCh <- returnReq
	}()

	returnReq := <-returnCh
	return c.JSON(http.StatusOK, returnReq)
}

type ChatHistoryReturn struct {
	ChatID  string    `json:"chat_id"`
	Content string `json:"content"`
	Type    string `json:"type"`
	UserID  int    `json:"user_id"`
	IsSeen  bool   `json:"is_seen"`
}
