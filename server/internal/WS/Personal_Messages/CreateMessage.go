//TODO: fix uuid
package ws

import (
	"context"
	"database/sql"
	"log"

	"github.com/Aritra640/ConnectSphere/server/Database/db"
	"github.com/google/uuid"
)

type CreatePersonalMassageInput struct {
	UserID     int32
	Content    string
	Type       string
	SenderID   int32
	ReceiverID int32
}

//CreatePersonalMassage checks if sender and receiver are friends and if not make them friends , create a new chat and create a personal massage 
func (pcs *PersonalChatService) CreatePersonalMassage(ctx context.Context , req CreatePersonalMassageInput) (int , error) {

  //Check if sender and receiver are friends and if not make them friends 
  areFriends,err := pcs.Queries.AreFriends(ctx , db.AreFriendsParams{
    UserID: req.SenderID,
    FriendID: req.ReceiverID,
  })
  if err != nil {
    log.Println("Error: Failed to check friendship: " , err)
    return -1,err
  }

  if !areFriends {
    err := pcs.Queries.AdduserFriendsBothWays(ctx , db.AdduserFriendsBothWaysParams{
      UserID: req.SenderID,
      FriendID: req.ReceiverID,
    }) 

    if err != nil {
      log.Println("Error: cannot add friends both ways: " , err)
      return -1,err
    }
  }

  //Create a chat int database (id is uuid)
  chatID := uuid.New()
  _,err = pcs.Queries.CreateChat(ctx , db.CreateChatParams{
    ID: chatID,
    UserID: sql.NullInt32{
      Int32: req.SenderID,
      Valid: true,
    },
    Content: req.Content,
    Type: req.Type,
  }) 
  if err != nil {
    log.Println("Error: cannot create chat: " , err)
    return -1,err
  }

  //Create a personal massage 
  _,err = pcs.Queries.CreatePersonalMessage(ctx , db.CreatePersonalMessageParams{
    ChatID: chatID,
    SenderID: req.SenderID,
    ReceiverID: req.ReceiverID,
  })

  if err != nil {
    log.Println("Error: could not create personal massage: " , err)
    return -1,err 
  }

  return int(chatID.ID()), nil
}
