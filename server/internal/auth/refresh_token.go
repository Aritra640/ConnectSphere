package auth

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"errors"
	"log"
	"time"

	"github.com/Aritra640/ConnectSphere/server/Database/db"
)

func GeneratedRefreshToken() (string, error) {

	b := make([]byte, 64)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

type RefreshTokenService struct {
	Queries *db.Queries
}

// CreateRefreshToken generates and stores a refresh token
func (rts *RefreshTokenService) CreateRefreshToken(ctx context.Context, userID int, expiry time.Duration) (string, error) {

	tokenCh, errChan := make(chan string), make(chan error)

	// Generate refresh token concurrently
	go func() {
		token, err := GeneratedRefreshToken()
		if err != nil {
			errChan <- err
		}

		tokenCh <- token
	}()

	var token string
	select {
	case <-ctx.Done():
		log.Println("In generate refresh token ctx has been ended")
		return "", ctx.Err()
	case err := <-errChan:
		log.Println("In generate refresh token err received")
		return "", err
	case token = <-tokenCh:
		log.Println("token recived from channel")
	}

	expiresAt := time.Now().Add(expiry)

	errChan1 := make(chan error)
	// Concurrently create a new refresh token
	go func() {
		_, err := rts.Queries.CreateNewRefreshToken(ctx, db.CreateNewRefreshTokenParams{
			UserID: sql.NullInt32{
				Int32: int32(userID),
				Valid: true,
			},
			Token:     token,
			ExpiresAt: sql.NullTime{Time: expiresAt, Valid: true},
		})

		errChan1 <- err
	}()
	err := <-errChan1
	if err != nil {
		return "", err
	}
	return token, nil
}

// VerifyRefreshToken verifies the token and the user id as well as the expiry
func (rts *RefreshTokenService) VerifyRefreshToken(ctx context.Context, UserID int, token string) (string, error) {
	//TODO: make the get refresh token function concurrent
	refreshToken, err := rts.Queries.GetRefreshTokenByUserId(ctx, sql.NullInt32{
		Int32: int32(UserID),
		Valid: true,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Println("refresh token not found")
			return "", errors.New("invalid refresh token")
		}

		log.Println("Error: something went wrong!")
		return "", err
	}

	if refreshToken.ExpiresAt.Valid && refreshToken.ExpiresAt.Time.Before(time.Now()) {
    doneChan := make(chan bool)
		go func() {
			_ = rts.Queries.DeleteRefreshTokenByUserID(ctx, sql.NullInt32{
				Int32: int32(UserID),
				Valid: true,
			})

      doneChan <- true
		}()

    <- doneChan
		log.Println("Token expired ...deleted")
		return "", errors.New("refresh token expired")
	}

	if refreshToken.Token != token {
		log.Println("Refresh token not matched")
		return "", errors.New("refresh token not matched")
	}

	return refreshToken.Token, nil
}

func (rts *RefreshTokenService) DeleteRefreshTokenByUserID(ctx context.Context, UserID int) error {

	errChan := make(chan error)
	go func() {

		err := rts.Queries.DeleteRefreshTokenByUserID(ctx, sql.NullInt32{
			Int32: int32(UserID),
			Valid: true,
		})

		errChan <- err

	}()

	err := <-errChan
	if err != nil {
		log.Println("Error in deleting token")
		return err
	}

	return nil
}
