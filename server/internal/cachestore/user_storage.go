package cachestore

import (
	"context"
	"log"
	"sync"
	"time"
)

// Add unverified user data to cachestore with 5 minute timer
func (cs *CacheStore) AddUnverifiedUser(ctx context.Context, userdata UnverifiedUserStore) {

	cs.Mu.Lock()
	cs.Store.Delete(userdata.Email)
	cs.Store.Store(userdata.Email, userdata)
	cs.Mu.Unlock()
}

// Get user data if exists and within time frame of 5 minutes
func (cs *CacheStore) GetUnverifiedUser(ctx context.Context, email string) (bool, UnverifiedUserStore) {

	cs.Mu.Lock()
	val, ok := cs.Store.Load(email)
	cs.Mu.Unlock()

	if !ok {
		log.Println("User cannot be found in cache store with email: ", email)
		return ok, UnverifiedUserStore{}
	}

	res := val.(UnverifiedUserStore)
	if time.Since(res.TimeStamp) > 5*time.Minute {
		log.Println("User data timed out with email id: ", email)
		cs.Mu.Lock()
		cs.Store.Delete(email)
		cs.Mu.Unlock()
		return false, UnverifiedUserStore{}
	}

	return true, val.(UnverifiedUserStore)
}

//Refresh the user cache storage with deleting the timed out data
func (cs *CacheStore) RefreshUnverifiedUserData(ctx context.Context , wg *sync.WaitGroup) {

	doneChan := make(chan bool)
	defer close(doneChan)

	go func() {
		var wg *sync.WaitGroup

    cs.Mu.Lock()
		cs.Store.Range(func(key, value any) bool {

			udb := value.(UnverifiedUserStore)
			wg.Add(1)

			go func(udb UnverifiedUserStore, wg *sync.WaitGroup) {
				defer wg.Done()

				if time.Since(udb.TimeStamp) > 5*time.Minute {
					log.Println("User data timed out with email id: ", udb.Email)
					cs.Store.Delete(udb.Email)
				}

			}(udb, wg)

			return true //Continue iterating
		})
    cs.Mu.Unlock()

		wg.Wait()
		doneChan <- true //Terminate go routine
	}()

	<-doneChan
}
