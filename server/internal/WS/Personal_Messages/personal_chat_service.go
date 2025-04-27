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

//Delete id pair if exists
func (store *PersonalChatStore) DeleteCons(id uuid.UUID) {
  
  store.Mu.Lock()
  if _,ok := store.Store[id]; ok {
    delete(store.Store , id)
  }
  store.Mu.Unlock()
}


