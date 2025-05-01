package ws

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Aritra640/ConnectSphere/server/Database/db"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type NewPersonalMsg struct {
  UserID1 int `json:"user_id1"`
  UserID2 int  `json:"user_id_2"`
}

func (pcs *PersonalChatService) GetPersonalMessageHandler(c echo.Context) error {

  var req NewPersonalMsg
  err := c.Bind(&req); if err != nil {
    log.Println("Error: cannot get proper input in CreateNewPersonalMessageHandler")
    return c.JSON(http.StatusBadRequest , "Invalid Request")
  }

  var pid uuid.UUID


  found,_ := pcs.Queries.AreFriends(c.Request().Context() , db.AreFriendsParams{
    UserID: int32(req.UserID1),
    FriendID: int32(req.UserID2),
  })
  if !found {
    pcs.Queries.AdduserFriendsBothWays(c.Request().Context() , db.AdduserFriendsBothWaysParams{
      UserID: int32(req.UserID1),
      FriendID: int32(req.UserID2),
    })

    pid = uuid.New()

    pcs.Queries.CreatePersonalWS(c.Request().Context() , db.CreatePersonalWSParams{
      ID: pid,
      Usera: sql.NullInt32{
        Int32: int32(req.UserID1),
        Valid: true,
      },
      Userb: sql.NullInt32{
        Int32: int32(req.UserID2),
        Valid: true,
      },
    })
  } else {

    obj,err := pcs.Queries.GetPersonalWSbyUsers(c.Request().Context() , db.GetPersonalWSbyUsersParams{
      Usera: sql.NullInt32{
        Int32: int32(req.UserID1),
        Valid: true,
      },
      Userb: sql.NullInt32{
        Int32: int32(req.UserID2),
        Valid: true,
      },
    })
    if err != nil {
      log.Println("Error: could not get ws id in Create new personal massages: " , err)
      return c.JSON(http.StatusConflict , "could not find ws id")
    }else{

      pid = obj.ID
    }
  }

  pcs.CreateChat(req.UserID1 , req.UserID2 , pid)
  return c.JSON(http.StatusOK , map[string]interface{}{
    "pid": pid,
  })
}


