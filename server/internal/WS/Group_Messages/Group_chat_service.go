package ws

import (
	"sync"

	"github.com/Aritra640/ConnectSphere/server/Database/db"
	"github.com/Aritra640/ConnectSphere/server/internal/utils"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type GroupChatService struct {
	Queries *db.Queries
	Groups  map[uuid.UUID]*Group
	Mu      sync.Mutex
  JWT     []byte
}

type Group struct {
	GroupID   uuid.UUID
	Clients   map[*websocket.Conn]bool
	Mu        sync.Mutex
	MessageCh chan GroupMessage
}

type GroupMessage struct {
	Content string           `json:"content"`
	Owner   *websocket.Conn  `json:"owner"`
}
