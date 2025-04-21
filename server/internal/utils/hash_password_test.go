package utils

import "testing"

func TestHashPasswordAndVerifyHashedPassword(t *testing.T) {
  password := "abcrdedadcwcwqcqcw"

  hashed_password,err := HashPassword(password)
  if err != nil {
    t.Fatalf("HashPassword failed : %v" , err)
  }
  
  if hashed_password == password {
    t.Error("Expected hashed_password to differ from plain password")
  }

  if !VerifyHashedPassword(password , hashed_password) {
    t.Error("Expected password to be verified by VerifyHashedPassword with hashed_password")
  }

  if VerifyHashedPassword("wrong_password" , hashed_password) {
    t.Error("Expected wrong password to not be verified by VerifyHashedPassword with hashed_password")
  }
}
