package utils

import (
	"encoding/json"
	"log"
)

type TypeMessage string

const (
	Chat TypeMessage = "Chat"
	Join TypeMessage = "Join"
)

type ResponsePersonal struct {
	SenderID   int         `json:"sender_id"`
	ReceiverID int         `json:"receiver_id"`
	Content    string      `json:"content"`
	TypeMsg    TypeMessage `json:"type_message"`
	Type       TypeStruct  `json:"type"`
}

type TypeStruct string

const (
	Text  TypeStruct = "text"
	Image TypeStruct = "image"
	Emoji TypeStruct = "emoji"
	PDF   TypeStruct = "pdf"
	Video TypeStruct = "video"
)

func GetPersonalPersonal_JSON(str []byte) (ResponsePersonal, error) {

	var res ResponsePersonal
	err := json.Unmarshal(str, &res)
	if err != nil {
		log.Println("Error: cannot transform data in response: ", err)
		return ResponsePersonal{}, err
	}

	return res, nil
}
