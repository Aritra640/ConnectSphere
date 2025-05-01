package ws

import "github.com/google/uuid"

func (pcs *PersonalChatService) CreateChat(uid1, uid2 int , id uuid.UUID)  {

	pcs.Mu.Lock()
	pcs.CUID[id] = PersonalChatID_UIDmap{
		User1: uid1,
		User2: uid2,
	}
	pcs.Mu.Unlock()
}

func (pcs *PersonalChatService) FindPersonalMessageID(uid1, uid2 int) (uuid.UUID, bool) {

	pcs.Mu.Lock()

	for id, participants := range pcs.CUID {
		if (participants.User1 == uid1 && participants.User2 == uid2) || (participants.User2 == uid1 && participants.User1 == uid2) {
			return id, true
		}
	}

	pcs.Mu.Unlock()
	return uuid.UUID{}, false
}

func (pcs *PersonalChatService) DeleteCUID(id uuid.UUID) {

	pcs.Mu.Lock()
	delete(pcs.CUID, id)
	pcs.Mu.Unlock()
}
