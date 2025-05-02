package ws

import (
	"context"

	"github.com/google/uuid"
)

func (gcs *GroupChatService) RequestJoinGroup(ctx context.Context , uid int , gid uuid.UUID) {

  //TODO: check if the user is already a member 
  //TODO: check if the group is restrictive -- if the group is restricteive add it to the request queue else add user to group 


}
