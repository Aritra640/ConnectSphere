package ws

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/Aritra640/ConnectSphere/server/Database/db"
	"github.com/Aritra640/ConnectSphere/server/internal/utils"
	"github.com/google/uuid"
)

type EditGroupMessageParam struct {
	ChatID      uuid.UUID        `json:"chat_id"`
	UserID      int32            `json:"user_id"`
	GroupID     uuid.UUID        `json:"group_id"`
	Content     string           `json:"content"`
	TypeContent utils.TypeStruct `json:"type_content"`
}

// EditGroupMessage edits a message in the group
func (gcs *GroupChatService) EditGroupMessage(ctx context.Context, req EditGroupMessageParam) error {
	//Check if the edited message type and type of the previous message is same
	//If all checks up, edit chat and then edit group message

	errChan := make(chan error)
	chatChan := make(chan db.Chat)
	defer close(errChan)
	defer close(chatChan)

	go func() {
		chat, err := gcs.Queries.GetChatByID(ctx, req.ChatID)

		if err != nil {
			errChan <- err
		}
		chatChan <- chat
	}()

	select {
	case <-ctx.Done():
		log.Println("Error: Timed out")
		return errors.New("Timed out")

	case err := <-errChan:
		log.Println("Error: Error found in EditGroupMessage: ", err)
		return err

	case chat := <-chatChan:

		if chat.Type != req.TypeContent {
			log.Println("Chat content type not matched in EditGroupMessage")
			return errors.New("Error: chat content type not matched")
		}

		userid := sql.NullInt32{
			Int32: req.UserID,
			Valid: true,
		}

		if chat.UserID != userid {
			log.Printf("In EditGroupMessage user with userid %v does not maches with userid %v of the chat", req.UserID, chat.UserID.Int32)
			return errors.New("Error: userids not matched")
		}
	}

	errCh := make(chan error)
	defer close(errCh)

	go func() {

		err := gcs.Queries.UpdateGroupMessageContent(ctx, db.UpdateGroupMessageContentParams{
			ID:      req.ChatID,
			Content: req.Content,
		})

		errCh <- err
	}()

	err := <-errCh
	if err != nil {
		log.Println("Error: In EditGroupMessage , could not update messsage with error: ", err)
		return err
	}

	return nil
}

func (gcs *GroupChatService) GroupHistoryByID(ctx context.Context, gid uuid.UUID)

type DeleteGroupChatParam interface{}

func (gcs *GroupChatService) DeleteGroupChat(ctx context.Context, req DeleteGroupChatParam)

//------------------------Group utility-----------------------------------------

type CreateNewGroupParam interface{}

func (gcs *GroupChatService) CreateNewGroup(ctx context.Context, req CreateNewGroupParam)

func (gcs *GroupChatService) GetGroupDesc(ctx context.Context, gid uuid.UUID)

type EditGroupDetailParam interface{}

func (gcs *GroupChatService) EditGroupDetail(ctx context.Context, req EditGroupDetailParam)

type DeleteGroupParam interface{}

func (gcs *GroupChatService) DeleteGroupByID(ctx context.Context, req DeleteGroupParam)
