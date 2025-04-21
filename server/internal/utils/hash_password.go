package utils

import "golang.org/x/crypto/bcrypt"

//HashPassword hashes the password with bcrypt
func HashPassword(password string) (string , error) {

  hsp,err := bcrypt.GenerateFromPassword([]byte(password) , bcrypt.DefaultCost)
  return string(hsp) , err
}

//VerifyHashedPassword verifies a plain text password against the bcrypt hash
func VerifyHashedPassword(password , hash string) bool {

  err := bcrypt.CompareHashAndPassword([]byte(hash) , []byte(password))
  return err == nil
}
