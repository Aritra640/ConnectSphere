//TODO: Fix SendMessageInGroup
package ws

import (
	"context"
	"errors"
	"log"

	"github.com/Aritra640/ConnectSphere/server/Database/db"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

func NewGCS() *GroupChatService {

	return &GroupChatService{
		Queries: &db.Queries{},
		Groups:  make(map[uuid.UUID]*Group),
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
	delete(gcs.Groups, gid)
	gcs.Mu.Unlock()

	return nil
}

// Remove the socket connection from all the groups
func (gcs *GroupChatService) DeleteClientFromAllGroups(ctx context.Context, socket *websocket.Conn) {

	gcs.Mu.Lock()
	for _, grp := range gcs.Groups {
		grp.DeleteClient(socket)
	}
	gcs.Mu.Unlock()
}

//------------------------WS------------------------------

// Run all groups concurrently
func (gcs *GroupChatService) RunAll(ctx context.Context) {

	for _, grp := range gcs.Groups {
		go grp.Run(ctx)
	}
}

//Send Message to group with group id 
func (gcs *GroupChatService) SendMessageInGroup(ctx context.Context, guid uuid.UUID, MessageStr GroupMessage) error {

	errChan := make(chan error)
	defer close(errChan)
  MsgChan := make(chan GroupMessage)
  defer close(MsgChan)

	go func() {

		gcs.Mu.Lock()
		if grp, ok := gcs.Groups[guid]; ok {

			grp.Mu.Lock()
			if _, found := grp.Clients[MessageStr.Owner]; found {
				//Best case
        log.Println("User sending message in group with id: " , guid)
				MsgChan <- MessageStr
			} else {

				log.Println("Error: user attempting to send message not in group or has exited ws group with in SendMessageInGroup")
				errChan <- errors.New("User not found in group in SendMessageInGroup")
			}

			grp.Mu.Unlock()
		} else {

      log.Println("Error: user attempting to send message in group, group not found in SendMessageInGroup")

      errChan <- errors.New("Group not found!")
    }
		gcs.Mu.Unlock()
	}()

  select{
  case <-ctx.Done(): 
    log.Println("Error: Timed out in *gcs.SendMessageInGroup")
    return errors.New("Timed out")

  case err := <-errChan: 
    log.Println("Error: cannot send message in group: " , err)
    return err

  case msg := <-MsgChan: 
    log.Println("Group Message Successfully sent!")
    gcs.Groups[guid].MessageCh <- msg
    return nil

  }
}


