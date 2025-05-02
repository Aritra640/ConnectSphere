package ws

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/Aritra640/ConnectSphere/server/Database/db"
	"github.com/google/uuid"
)

func (gcs *GroupChatService) createChat(ctx context.Context , userID int , typeContent string , content string) (uuid.UUID , error) {

  cid := uuid.New() 

  _,err := gcs.Queries.CreateChat(ctx , db.CreateChatParams{
    ID: cid,
    Content: content,
    Type: typeContent,
    UserID: sql.NullInt32{
      Int32: int32(userID),
      Valid: true,
    },
  })

  return cid , err
}

func (gcs *GroupChatService) CreateGroupMessage(ctx context.Context, gid uuid.UUID, content string, uid int , typeContent string) error {

  cidCh := make(chan uuid.UUID)
  errCh := make(chan error)

  go func() {
    cid,err := gcs.createChat(ctx , uid , typeContent , content)

    cidCh <- cid 
    if err != nil {
      errCh <- err
    }
  }()

  var cid uuid.UUID 

  select{
  case err := <-errCh :
    if err != nil {
      log.Println("Error: cannot create chat in CreateGroupMessage: " , err)
      return err
    }
  case ChatID := <-cidCh:
    cid = ChatID
    log.Println("chat created in CreateGroupMessage")
  }

	errChan := make(chan error)

	go func() {

		_, err := gcs.Queries.CreateGroupMessage(ctx, db.CreateGroupMessageParams{
			ChatID:      cid,
			ChatGroupID: gid,
			SenderID:    int32(uid),
		})

		errChan <- err

	}()

	select {
	case err := <-errChan:
		if err != nil {
			log.Println("Error: cannot create group message: ", err)
			return err
		}
    log.Println("Group Message Created with userID: " , uid)
		return nil
	case <-ctx.Done():
		log.Println("Error: context done")
		return errors.New("Error: context timed out")
	}
}
