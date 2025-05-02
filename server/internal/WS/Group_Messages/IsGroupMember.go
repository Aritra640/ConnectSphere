package ws

import (
	"context"
	"errors"
	"log"

	"github.com/Aritra640/ConnectSphere/server/Database/db"
	"github.com/google/uuid"
)

func (gcs *GroupChatService) IsGroupMember(ctx context.Context , gid uuid.UUID , uid int32) (bool, error) {

  foundCh := make(chan bool)
  errCh := make(chan error)

  go func() {

    found,err := gcs.Queries.IsUserInGroup(ctx, db.IsUserInGroupParams{
      UserID: uid,
      GroupID: gid,
    })

    if err != nil {
      errCh <- err
    }
    foundCh <- found
  }()
  
  var found bool
  select{
  case err := <-errCh:
    log.Println("Error: something went wrong in IsGroupMember: " , err)
    return false, err

  case found = <-foundCh: 
    return found, nil

  case <-ctx.Done() :
    log.Println("Error: context timed out in IsGroupMember")
    return false , errors.New("timed out")
  }
}
