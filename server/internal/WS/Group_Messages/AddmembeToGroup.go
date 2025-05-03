package ws

import (
	"context"
	"database/sql"
	"log"

	"github.com/Aritra640/ConnectSphere/server/Database/db"
	"github.com/google/uuid"
)


type AddMemberParams struct {

  UserID  int32
  GroupID uuid.UUID
  IsAdmin bool
}

//Add member to group
func (gcs *GroupChatService) AddMemberToGroup(ctx context.Context , req AddMemberParams) error {

  errChan := make(chan error)

  go func() {
    
    err := gcs.Queries.AddUserToGroup(ctx , db.AddUserToGroupParams{
      UserID: req.UserID,
      GroupID: req.GroupID,
      IsAdmin: sql.NullBool{
        Bool: req.IsAdmin,
        Valid: true,
      },
    })

    errChan <- err
  }()

  err := <-errChan 
  if err != nil {
    log.Println("Error: cannot add user to group in database: " , err)
  }
  return err
}
