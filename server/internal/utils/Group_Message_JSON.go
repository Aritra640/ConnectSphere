package utils

import (
	"encoding/json"
	"log"
)

type RequestGroup struct {
	UserID      int         `json:"user_id"`
	RequestType TypeMessage `json:"type_request"`
	Payload     Payload     `json:"payload"`
}

type Payload struct {
	Content string     `json:"content"`
	TypeMsg TypeStruct `json:"type"`
}

func GetRequestGroup_JSON(str []byte) (RequestGroup, error) {

	var res RequestGroup
	err := json.Unmarshal(str, &res)
	if err != nil {

		log.Println("Error: cannot unmarshall in group message: ", err)
		log.Println("String: ", string(str))
		return RequestGroup{}, err
	}

	return res, nil
}
