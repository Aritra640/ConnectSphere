package ws

import (
	"log"
	"net/http"

	"github.com/Aritra640/ConnectSphere/server/Database/db"
	"github.com/Aritra640/ConnectSphere/server/internal/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type EditPersonalMessageParam struct {
	ChatID  string `json:"chat_id"`
	Content string    `json:"content"`
}

func (pcs *PersonalChatService) EditPersonalMessageHandler(c echo.Context) error {

	reqCh := make(chan EditPersonalMessageParam)
	errCh := make(chan error)
	go func() {
		var req EditPersonalMessageParam
		if err := c.Bind(&req); err != nil {
			errCh <- err
		}
		reqCh <- req
	  }()
    
  var chatIdparam uuid.UUID

	var req EditPersonalMessageParam
	select {
	case req = <-reqCh:
    uuid,err := utils.ParseUUID(req.ChatID); if err != nil {
      log.Println("Error: cannot be converted to uuid in edit param")
    }
    chatIdparam = uuid

	case err := <-errCh:
		log.Println("Error: cannot get req EditPersonalMessageParam in EditPersonalMessageHandler: ", err)
		return c.JSON(http.StatusBadRequest, "Invalid Request")
	}

	go func() {
		err := pcs.Queries.EditMessageContent(c.Request().Context(), db.EditMessageContentParams{
			ID:      chatIdparam,
			Content: req.Content,
		})
		errCh <- err
	}()

	err := <-errCh
	if err != nil {
		log.Println("Error: cannot edit message: ", err)
		return c.JSON(http.StatusConflict, "Something went wrong")
	}
	return c.JSON(http.StatusOK, "edit messages successfull!")
}
