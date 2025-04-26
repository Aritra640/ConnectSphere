package config

import (
	"context"
	"database/sql"

	"github.com/Aritra640/ConnectSphere/server/Database/db"
	gcs "github.com/Aritra640/ConnectSphere/server/internal/WS/Group_Messages"
	pcs "github.com/Aritra640/ConnectSphere/server/internal/WS/Personal_Messages"
)

type Config struct {
	DB       *sql.DB
	QueryObj *db.Queries
	CTX      context.Context
  JWT      []byte
  PCS      *pcs.PersonalChatService
  GCS      *gcs.GroupChatService
}

