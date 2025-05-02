package ws

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

func (gcs *GroupChatService) AddUserInGroup(gid uuid.UUID, socket *websocket.Conn) error {

	gcs.Mu.Lock()
	gcs.Groups[gid].AddClient(socket)
	gcs.Mu.Unlock()

	return nil
}

func (gcs *GroupChatService) DeleteUserInGroup(gid uuid.UUID, socket *websocket.Conn) error {

	gcs.Mu.Lock()
	gcs.Groups[gid].DeleteClient(socket)
	gcs.Mu.Unlock()

	return nil
}

func (gcs *GroupChatService) DeleteGroup(gid uuid.UUID) error {

  gcs.Mu.Lock()
  delete(gcs.Groups , gid)
  gcs.Mu.Unlock()

  return nil
}


