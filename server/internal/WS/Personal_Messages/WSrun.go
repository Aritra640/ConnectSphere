package ws

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

func (store *PersonalChatStore) AddClient(id uuid.UUID, socket *websocket.Conn) {
 
	store.Mu.Lock()
	store.Store[id] = append(store.Store[id], socket)
	store.Mu.Unlock()
}

func (store *PersonalChatStore) Run(ctx context.Context) {
	for {
		select {

		case msg := <-store.MessageCh:
      
      if msg.TypeMsg == Chat {

          store.Mu.Lock()
          for _,socket := range store.Store[msg.Pid] {
            socket.WriteMessage(websocket.TextMessage ,  []byte(msg.chat))
          }
      }


	  case <-ctx.Done():
			log.Println("Run looped stopped ")
			return
		}
	}
}

func (store *PersonalChatStore) DeleteConn(id uuid.UUID) {

	store.Mu.Lock()
	delete(store.Store, id)
	store.Mu.Unlock()
}

func (store *PersonalChatStore) SendMesssage(msgContent string , socket *websocket.Conn , pid uuid.UUID , t TypeMessage) {
  
  store.Mu.Lock()
  store.MessageCh <- Message{
    Owner: socket,
    chat: msgContent,
    Pid: pid,
    TypeMsg: t,
  } 
  store.Mu.Unlock()
}


func (store *PersonalChatStore) RunWS(ctx context.Context) {

  go store.Run(ctx)
}
