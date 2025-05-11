package ws

import (
	"context"
	"log"

	"github.com/gorilla/websocket"
)

func deleteclientAll(ctx context.Context , socket *websocket.Conn) {
  
  GroupChatMessageSetup.DeleteClientFromAllGroups(ctx , socket)
  log.Println("Error: Deleted client from all groups in GroupHandlerWS (helper)")
}

func (gc *Group) Run(ctx context.Context) {
	for {
		select {

		case msg := <-gc.MessageCh:
			gc.Mu.Lock()
			for socket := range gc.Clients {
				err := socket.WriteMessage(websocket.TextMessage, []byte(msg.Content))
				if err != nil {
					log.Println("Error in group.Run :", err)
					//TODO: delete socket
          gc.DeleteClient(socket)
          deleteclientAll(ctx , socket)
					socket.Close()
				}
			}
			gc.Mu.Unlock()

		case <-ctx.Done():
			log.Println("WS group run stopped")
			return
		}
	}
}
