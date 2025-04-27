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

func (store *PersonalChatStore) Run(ctx context.Context, id uuid.UUID) {
	for {
		select {

		case msg := <-store.MessageCh:
			store.Mu.Lock()

			for _, socket := range store.Store[id] {
				socket.WriteMessage(websocket.TextMessage, []byte(msg.chat))
			}

			store.Mu.Unlock()
		case <-ctx.Done():
			log.Println("Run looped stopped for: ", id)
			return
		}
	}
}

func (store *PersonalChatStore) DeleteConn(id uuid.UUID) {

	store.Mu.Lock()
	if _, ok := store.Store[id]; ok {
		delete(store.Store, id)
	}
	store.Mu.Unlock()
}


