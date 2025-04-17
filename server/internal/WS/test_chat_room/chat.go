package ws

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//for test a message is a normal string and does not need a
//owner

func TestChatRoom(c echo.Context) error {

	ws, err := upgrader.Upgrade(c.Response(), c.Request(), c.Response().Header())
	if err != nil {
		log.Println("could not upgrade websocker: ", err.Error())
		return c.JSON(500, "something went wrong")
	}
	defer ws.Close()

  mygroup.clients[ws] = true

	for {

		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Println("Error: client disconnected: ", err)
			delete(mygroup.clients, ws)
			break
		}

		//send message
		mygroup.MessageCh <- string(msg)
	}

	return c.JSON(200, "socket ended")
}

type Group struct {
	clients   map[*websocket.Conn]bool
	RoomMu    sync.Mutex
	MessageCh chan string
}

func newGroup() *Group {
	return &Group{
		clients:   make(map[*websocket.Conn]bool),
		MessageCh: make(chan string),
	}
}

func (g *Group) run() {
  for msg := range mygroup.MessageCh {

		g.RoomMu.Lock()
		for client := range g.clients {

			err := client.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				log.Println("Error: ", err)
				client.Close()
			}
		}

		g.RoomMu.Unlock()
	}
}

var mygroup = newGroup()

func Start_test_group() {

	go mygroup.run()
}
