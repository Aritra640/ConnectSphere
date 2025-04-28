package ws

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type NewPersonalMsg struct {
  UserID1 int `json:"user_id1"`
  UserID2 int  `json:"user_id_2"`
}

func (pcs *PersonalChatService) CreateNewPersonalMessageHandler(c echo.Context) error {

  var req NewPersonalMsg
  err := c.Bind(&req); if err != nil {
    log.Println("Error: cannot get proper input in CreateNewPersonalMessageHandler")
    return c.JSON(http.StatusBadRequest , "Invalid Request")
  }


  pid := pcs.CreateChat(req.UserID1 , req.UserID2)
  return c.JSON(http.StatusOK , map[string]interface{}{
    "pid": pid,
  })
}


