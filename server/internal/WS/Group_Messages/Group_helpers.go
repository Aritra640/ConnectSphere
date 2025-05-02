package ws

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

func NewGroup() *Group {

	return &Group{
		GroupID:   uuid.UUID{},
		Clients:   make(map[*websocket.Conn]bool),
		MessageCh: make(chan GroupMessage),
	}
}

func (gc *Group) AddClient(socket *websocket.Conn) {

	gc.Mu.Lock()
	gc.Clients[socket] = true
	gc.Mu.Unlock()
}

func (gc *Group) DeleteClient(socket *websocket.Conn) {

	gc.Mu.Lock()
	delete(gc.Clients, socket)
	gc.Mu.Unlock()
}
