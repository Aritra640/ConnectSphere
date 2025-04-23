package auth

import (
	"testing"
)

func TestGenerateAndParseJWT(t *testing.T) {

  userId := 12345

  token,err := CreateAuthToken(userId) 
  if err != nil {
    t.Fatalf("Failed to generate jwt: %v" , err)
  }

  claims_user_id,err := VerifyToken(token)
  if err != nil {
    t.Fatalf("Failed to parse JWT: %v" , err )
  }

  if claims_user_id != userId {
    t.Errorf("Expected user id %d, got %d" , userId , claims_user_id)
  }
}
