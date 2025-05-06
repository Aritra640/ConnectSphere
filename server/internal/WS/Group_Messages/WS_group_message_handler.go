// TODO: CheckOrigin should not be always true in prod
package ws

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/Aritra640/ConnectSphere/server/internal/utils"
	"github.com/golang-jwt/jwt/v5"
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

func (gcs *GroupChatService) WSGroupMessageHandler(c echo.Context) error {

	authToken := c.QueryParam("access_token")
	uidChan := make(chan int)
	errChan := make(chan error)

	go func() {

		uid, err := gcs.WSverify_access_token(authToken)
		if err != nil {
			errChan <- err
		}

		uidChan <- uid
	}()

	var uid int

	select {
	case <-errChan:
		log.Println("WS authentication failed")
		return c.JSON(http.StatusBadRequest, "Invalid Request")

	case uid = <-uidChan:
		log.Println("WS authentcation successfull with user id: ", uid)

	case <-c.Request().Context().Done():
		log.Println("Request cancelled (WS group) ...")
	}

  gid := c.QueryParam("gid")
  guid,err := uuid.Parse(gid); if err != nil {
    log.Println("Group id invalid in WS group messsage handler: " , err)

    return c.JSON(http.StatusBadRequest, "Invalid request")
  }

	ws, err := upgrader.Upgrade(c.Response(), c.Request(), c.Response().Header())
	if err != nil {
		log.Println("WS connection failed in WS group message handler: ", err)
		return c.JSON(http.StatusInternalServerError, "Something went wrong")
	}
	defer ws.Close()


  for{

    _,msg,err := ws.ReadMessage()
    if err != nil {

      //delete this socket 
      log.Println("WS client disconected in WS_group_message_handler: " , err)
      break
    }

    req,err := utils.GetRequestGroup_JSON(msg)
    if err != nil {
      str,_ := GroupMessageString(uid , guid , true , "Message type invalid warning" , utils.Text)
      log.Println("Error: request message not of valid type: " , err)
      ws.WriteMessage(websocket.TextMessage , []byte(str))
    }

    if req.RequestType == utils.Join {
      //Create a join request in the queue if group is restricted else add user to group 
      errChan := make(chan error)

      go func() {

        err := gcs.AddMemberToGroup(c.Request().Context() , AddMemberParams{
          UserID: int32(uid),
          GroupID: guid,
        })

        errChan <- err
      }()

      err := <-errChan; if err != nil {
        log.Println("Error: join request failed in ws group handler: " , err)
        str,_ := GroupMessageString(uid , guid , true , "Join Request failed" , utils.Text)
        ws.WriteMessage(websocket.TextMessage , []byte(str))
      }

    }else {
      //create a group message (through database) and write back (ws)
      errChan := make(chan error)

      go func() {

        err := gcs.CreateGroupMessage(c.Request().Context() , guid , req.Payload.Content , uid , string(req.Payload.TypeMsg))
        errChan <- err
      }()
      
      err := <-errChan
      if err != nil {

        log.Println("Error: message creatiob failed in WS group message: " ,err)
        str,_ := GroupMessageString(uid , guid , true , "Message failed!" , utils.Text)
        ws.WriteMessage(websocket.TextMessage , []byte(str))
      }else {

        str,_ := GroupMessageString(uid , guid , false , req.Payload.Content , utils.Text)
        //send str to group socket 

      }
    }
    
  }

  log.Println("WS group handler has shutdown for uid: " , uid)
  return c.JSON(http.StatusOK , "ws connection ended")

}


func GroupMessageString(userId int , groupID uuid.UUID , isError bool , content string , type_content utils.TypeStruct) (string , error) {

  data := map[string]interface{}{
    "user_id": userId,
    "group_id": groupID,
    "is_error": isError,
    "content": content,
  }

  JSONdata,err := json.Marshal(data)
  if err != nil {

    log.Println("Error: cannot marsha data in GroupMessageString: " , err)
    return "", err
  }

  return string(JSONdata) , nil
}


func GroupMessageResponseStringHandler(c echo.Context) error {

  str,err := GroupMessageString(123 , uuid.New() , false , "This is an example of gropup message response" , utils.Text)

  if err != nil {
    log.Println("Error: GroupMessageResponseStringHandler failed! :" , err)
    return c.JSON(http.StatusInternalServerError , "Error in getting string res")
  }

  return c.JSON(http.StatusOK , str)
}



type myCustomClaims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

func (gcs *GroupChatService) WSverify_access_token(tokenString string) (int, error) {

	token, err := jwt.ParseWithClaims(tokenString, &myCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		//Validate the signing method (HMAC)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method : %v", token.Header["alg"])
		}
		return gcs.JWT, nil
	})

	if err != nil {
		log.Println("Error in parsing token: ", err.Error())
		return -1, err
	}

	claims, ok := token.Claims.(*myCustomClaims)
	if !ok || !token.Valid {
		log.Println("invalid token or claims")
		return -1, errors.New("error cannot verify token")
	}

	log.Println("Token verified successfully for user-id: ", claims.UserID)
	return claims.UserID, nil
}
