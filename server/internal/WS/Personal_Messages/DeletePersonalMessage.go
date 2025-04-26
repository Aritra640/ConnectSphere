package ws

import (
	"log"
	"net/http"

	"github.com/Aritra640/ConnectSphere/server/internal/utils"
	"github.com/labstack/echo/v4"
)

type DeletePersonalMessageParam struct {
	ChatId string `json:"chat_id"`
}

func (pcs *PersonalChatService) DeletePersonalMassage(c echo.Context) error {

	reqCh := make(chan DeletePersonalMessageParam)
	errCh := make(chan error)
	go func() {
		var req DeletePersonalMessageParam
		err := c.Bind(&req)
		if err != nil {
			errCh <- err
		}

		reqCh <- req
	}()

	var req DeletePersonalMessageParam
	select {
	case req = <-reqCh:
    uuid,erri := utils.ParseUUID(req.ChatId); if erri != nil {
      log.Println("could not convert into uuid")
      return c.JSON(hhtp.StatusBadRequest , "Invalid Request")
    }
	case err := <-errCh:
		log.Println("Error: could not get chat id in delete personal message handler: ", err)
		return c.JSON(http.StatusBadRequest, "Invalid Request")
	}

	err := pcs.Queries.DeletePersonalMessage(c.Request().Context(), req.ChatId)
	if err != nil {
		log.Println("Error: could not delete chat in delete persona message handler: ", err)
		return c.JSON(http.StatusInternalServerError, "Something went wrong")
	}

	return c.JSON(http.StatusOK, "chat deleted successfully")
}
