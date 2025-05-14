package mail

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"
)


type OTPstore struct {

  Mu sync.Mutex
  Store map[string]OtpStoreParam 
}

type OtpStoreParam struct {

  Email string 
  TimeStamp time.Time
}

//Check if otp exists in the otp store 
func (ots *OTPstore) CheckOTP(ctx context.Context , otp string) bool {

  ots.Mu.Lock()
  _,found := ots.Store[otp]
  ots.Mu.Unlock()

  return found
}

//Add otp 
func (ots *OTPstore) AddOtp(ctx context.Context , email string , otp string) error {

  ots.Mu.Lock()
  ots.Store[otp] = OtpStoreParam{
    Email: email,
    TimeStamp: time.Now(),
  }
  ots.Mu.Unlock()

  return nil
}

//Verify OTP 
func (ots *OTPstore) VerifyOTP(ctx context.Context , email string , otp string) (bool , error) {

  ots.Mu.Lock()
  otp_store,found := ots.Store[otp]
  ots.Mu.Unlock()

  if !found {
    return false , errors.New("Error: could not fetch email-otp pair")
  }

  timeSince := time.Since(otp_store.TimeStamp)
  if timeSince > 5*time.Minute {
    log.Println("Time expired in VerifyOTP")
    ots.DeleteOTP(ctx, otp)
    return false, errors.New("Time Expired")
  }

  if otp_store.Email != email {
    log.Println("In verify otp email not matched")
    return false, errors.New("Email did not match")
  }

  return true , nil
}


//Delete otp pair from store 
func (ots *OTPstore) DeleteOTP(ctx context.Context , otp string) {

  ots.Mu.Lock()
  delete(ots.Store , otp)
  ots.Mu.Unlock()
}


//Refresh OTP stores
func (ots *OTPstore) RefreshStore(ctx context.Context) {

  doneCh := make(chan bool)
  defer close(doneCh)

  go func() {
    ots.Mu.Lock()
    
    for otp,store := range ots.Store {

      if time.Since(store.TimeStamp) > 5*time.Minute {
        delete(ots.Store , otp)
      }
    }
    ots.Mu.Unlock()

    doneCh <- true
  }()


  <-doneCh
}
