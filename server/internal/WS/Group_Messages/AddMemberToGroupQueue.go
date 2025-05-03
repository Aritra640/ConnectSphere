package ws

import (
	"context"
	"log"

	"github.com/Aritra640/ConnectSphere/server/Database/db"
	"github.com/google/uuid"
)

type AddToGroupQueueParam struct {
	UserID  int
	GroupID uuid.UUID
}

func (gcs *GroupChatService) AddToGroupQueue(ctx context.Context, req AddToGroupQueueParam) error {

	errCh := make(chan error)

	go func() {

		err := gcs.Queries.AddGroupJoinRequest(ctx, db.AddGroupJoinRequestParams{
			UserID:      int32(req.UserID),
			ChatGroupID: req.GroupID,
		})

		errCh <- err
	}()

	err := <-errCh
	if err != nil {
		log.Println("Error: Failed to add member to group: ", err)
		return err
	}

	return nil
}
