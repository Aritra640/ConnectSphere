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
}

type Group struct {
	GroupID   uuid.UUID
	Clients   map[*websocket.Conn]bool
	Mu        sync.Mutex
	MessageCh chan GroupMessage
}

type GroupMessage struct {
	UserID  int              `json:"user_id"`
	ChatID  uuid.UUID        `json:"chat_id"`
	Content string           `json:"content"`
	TypeMsg utils.TypeStruct `json:"type_msg"`
	Owner   *websocket.Conn  `json:"owner"`
}
