package ws

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type MASinput struct {
	ChatID string `json:"chat_id"`
}

func (pcs *PersonalChatService) PersonalMessageMarkAsSeenHandler(c echo.Context) error {

	var req MASinput
	err := c.Bind(&req)
	if err != nil {
		log.Println("could not get chat id in MasrAsSeen Handler: ", err)
		return c.JSON(http.StatusBadRequest, "Invalid Request")
	}

	id, err := uuid.Parse(req.ChatID)
	if err != nil {
		log.Println("Error: cannot convert chat id to (uuid) int: ", err)
		return c.JSON(http.StatusBadRequest, "Invalid Request")
	}

	err = pcs.Queries.MarkMessageAsSeen(c.Request().Context(), id)
	if err != nil {
		log.Println("Error: cannot mark massage as seen: ", err)
		return c.JSON(http.StatusConflict, "could not mark massage")
	}

	return c.JSON(http.StatusOK, "message successfully marked as seen")
}
