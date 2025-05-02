package ws

import (
	"context"

	"github.com/Aritra640/ConnectSphere/server/Database/db"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

func NewGCS() *GroupChatService {

  return &GroupChatService{
    Queries: &db.Queries{},
    Groups: make(map[uuid.UUID]*Group),
  }
}

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



//------------------------WS------------------------------

//Run all groups concurrently 
func (gcs *GroupChatService) RunAll(ctx context.Context) {

  for _,grp := range gcs.Groups {
    go grp.Run(ctx)
  }
}

