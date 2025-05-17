package ws

import (
	"net/http"

	"github.com/Aritra640/ConnectSphere/server/internal/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (gcs *GroupChatService) ExampleGroupRequestHandler(c echo.Context) error {

	return c.JSON(http.StatusOK, utils.RequestGroup{
		UserID:      123,
		RequestType: utils.Chat,
		Payload: utils.Payload{
			Content: "This is the content of the chat, it can be url to a image , etc",
			TypeMsg: utils.Text,
		},
	})
}

func (gcs *GroupChatService) ExampleGroupMessageResponse(c echo.Context) error {

  str,_ := GroupMessageString(1234, uuid.New() , false , "this is a average group messsage response", utils.Text )

  return c.JSON(http.StatusOK , str)
}
