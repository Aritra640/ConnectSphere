package cachestore

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"sync"
	"time"
)

//Get a new otp
func getOTP(OTP chan string) {

  otp := ""
  for i := 0 ; i < 6 ; i ++ {

    num,_ := rand.Int(rand.Reader , big.NewInt(10))
    otp += fmt.Sprint(num)
  }

  OTP <- otp
}

//Get new otp and store it
func (cs *CacheStore) GetNewOTP(ctx context.Context , email string) string {

  cs.Mu.Lock()
  cs.Store.Delete(email)
  cs.Mu.Unlock()

  otpCh := make(chan string)
  defer close(otpCh)

  go getOTP(otpCh)

  otp := <-otpCh

  cs.Mu.Lock()
  cs.Store.Store(email , OtpStore{
    Email: email,
    Otp: otp,
    TimeStamp: time.Now(),
  })

  return otp
}

//Verify otp 
func (cs *CacheStore) VerifyOTP(ctx context.Context, email string , otp string) (found bool) {

  cs.Mu.Lock()
  val,found := cs.Store.Load(email)
  cs.Mu.Unlock()

  if !found {
    log.Println("No user found in verify otp function with email id: " , email)
    return
  }

  if val.(OtpStore).Otp == otp {
    log.Println("User otp matched successfully with email id: " , email)
    return
  }

  found = false
  log.Println("Otp didnot match with email id: ", email)
  return 
}

//Refresh otp cache storage with a 10 minute time limit
func (cs *CacheStore) RefreshOtpStorage(ctx context.Context , wg *sync.WaitGroup) {
  defer wg.Done()

  doneCh := make(chan bool)
  defer close(doneCh)

  go func() {
    var wg *sync.WaitGroup

    cs.Mu.Lock()
    cs.Store.Range(func(key, value any) bool {
      
      wg.Add(1)
      go func(otpStore OtpStore , wg *sync.WaitGroup){
        defer wg.Done()
        if time.Since(otpStore.TimeStamp) > 10 * time.Minute {
          log.Println("OTP deleted from storage with email id: " , otpStore.Email)
          cs.Store.Delete(otpStore.Email)
        }
        
      }(value.(OtpStore) , wg)

      return true
    })
    cs.Mu.Unlock()
    doneCh <- true
  }()

  <-doneCh
}
