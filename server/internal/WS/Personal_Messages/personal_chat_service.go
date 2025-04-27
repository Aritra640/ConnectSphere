package ws

import (
	"sync"

	"github.com/Aritra640/ConnectSphere/server/Database/db"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type PersonalChatService struct {
	Queries  *db.Queries
  WS_store *PersonalChatStore
  CUID     map[uuid.UUID]PersonalChatID_UIDmap
  Mu       sync.Mutex
}

type PersonalChatID_UIDmap struct {
  User1 int 
  User2 int
}

type Message struct {
  Owner *websocket.Conn
  chat  string
}


type PersonalChatStore struct {

  Store map[uuid.UUID][]*websocket.Conn
  Mu  sync.Mutex
  MessageCh chan Message
}


func NewPersonalChatStore() *PersonalChatStore {

  return &PersonalChatStore{
    Store: make(map[uuid.UUID][]*websocket.Conn),
    MessageCh: make(chan Message),
  }
}


