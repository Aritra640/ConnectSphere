package utils

import (
	"encoding/json"
	"log"
)

type ResponseGroup struct {
	GroupID string     `json:"group_id"`
	UserID  int        `json:"user_id"`
	Content string     `json:"content"`
	Type    TypeStruct `json:"type"`
}

func GetResponseGroup_JSON(str []byte) (ResponseGroup , error) {

  var res ResponseGroup 
  err := json.Unmarshal(str , &res); if err != nil {

    log.Println("Error: cannot unmarshall in group message: " ,err)
    log.Println("String: " , string(str))
    return ResponseGroup{} , err
  }

  return res, nil
}
