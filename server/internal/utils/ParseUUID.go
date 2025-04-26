package utils

import (
	"log"

	"github.com/google/uuid"
)

func ParseUUID(s string) (uuid.UUID , error) {
  
  errCh := make(chan error )
  uuidCh := make(chan uuid.UUID)
  go func() {
    
    uuid,err := uuid.Parse(s)
    if err != nil {
      errCh <- err
    }
    uuidCh <- uuid
  }()

  select{
  case uuid := <-uuidCh: 
    return uuid,nil 
  case err := <-errCh:
    log.Println("Error: cannot parse string to uuid: " , err)
    return uuid.UUID{} , err
  }
}
