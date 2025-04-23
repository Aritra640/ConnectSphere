package config

import (
	"context"
	"database/sql"

	"github.com/Aritra640/ConnectSphere/server/Database/db"
)

type Config struct {
	DB       *sql.DB
	QueryObj *db.Queries
	CTX      context.Context
  JWT      []byte
}

