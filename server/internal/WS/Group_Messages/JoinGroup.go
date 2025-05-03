package ws

import (
	"context"
	"log"

	"github.com/google/uuid"
)

func (gcs *GroupChatService) RequestJoinGroup(ctx context.Context, uid int, gid uuid.UUID) error {

	found, err := gcs.IsGroupMember(ctx, gid, int32(uid))
	if err != nil {
		log.Println("Error in searching through database")
		log.Println("Attempt to add user in database failed")
		return err
	}

	if found {
		log.Println("User in group")
		return nil
	}

	//User not in group
	grp, err := gcs.GetGroupDetails(ctx, gid)
	if err != nil {
		log.Println("Error: could not get data from database: ", err)
		return err
	}

	if !grp.RequiredPermission {
		err := gcs.AddMemberToGroup(ctx, AddMemberParams{
			UserID:  int32(uid),
			GroupID: gid,
			IsAdmin: false,
		})

		if err != nil {
			log.Println("Error: cannot add member to group")
			return err
		}
		return nil
	}

	//Grp is restricted -- add user to the queue for membership to the group
	err = gcs.AddToGroupQueue(ctx, AddToGroupQueueParam{
		UserID:  uid,
		GroupID: gid,
	})
	if err != nil {

		log.Println("Error: could not add user to queue: ", err)
		return err
	}

	return nil
}
