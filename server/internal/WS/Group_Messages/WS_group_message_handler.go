// TODO: CheckOrigin should not be always true in prod
package ws

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
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

	ws, err := upgrader.Upgrade(c.Response(), c.Request(), c.Response().Header())
	if err != nil {
		log.Println("WS connection failed in WS group message handler: ", err)
		return c.JSON(http.StatusInternalServerError, "Something went wrong")
	}
	defer ws.Close()


  for{

  }

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
