package auth

import (
	"time"

	"github.com/Aritra640/ConnectSphere/server/Database/db"
)

type AuthService struct {
	Queries *db.Queries
	Rts     *RefreshTokenService
	Expiry  time.Duration
}

type SignupRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type SigninRequest struct {
	Email          string `json:"email" validate:"required,email"`
	HashedPassword string `json:"hashed_password" validate:"required"`
}

// NewAuthService creates a new instance of AuthService with queries , refresh-token-service(rts) and expiry
func NewAuthService(
	queries *db.Queries,
	rts *RefreshTokenService,
	expiry time.Duration,
) *AuthService {

	return &AuthService{
		Queries: queries,
		Rts:     rts,
		Expiry:  expiry,
	}
}
