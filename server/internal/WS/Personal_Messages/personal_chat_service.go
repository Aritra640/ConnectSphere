package ws

import (
	"database/sql"

	"github.com/Aritra640/ConnectSphere/server/Database/db"
)

type PersonalChatService struct {
	Queries *db.Queries
}

func NewPersonalChatService(dbObj *sql.DB) *PersonalChatService {

  dbq := db.New(dbObj)
  return &PersonalChatService{
    Queries: dbq,
  }
}


