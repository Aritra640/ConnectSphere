package ws

import (
	"context"
	"errors"
	"log"

	"github.com/Aritra640/ConnectSphere/server/Database/db"
	"github.com/google/uuid"
)

func (gcs *GroupChatService) GetGroupDetails(ctx context.Context, gid uuid.UUID) (db.ChatGroup, error) {

	errChan := make(chan error)
	groupCh := make(chan db.ChatGroup)

	go func() {

		group, err := gcs.Queries.GetGroupByID(ctx, gid)
		if err != nil {
			errChan <- err
		}
		groupCh <- group
	}()

	select {
	case err := <-errChan:
		log.Println("Error in get group by id: ", err)
		return db.ChatGroup{}, err

	case grp := <-groupCh:
		return grp, nil

	case <-ctx.Done():
		log.Println("Context timed out in gcs.GetGroupDetails")
		return db.ChatGroup{}, errors.New("Timed out")
	}
}
