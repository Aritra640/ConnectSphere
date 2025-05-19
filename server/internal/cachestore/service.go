package cachestore

import (
	"sync"
	"time"
)

type CacheStore struct {
	Mu    sync.Mutex
	Store sync.Map
}

type UnverifiedUserStore struct {
	UserName       string `json:"user_name"`
	HashedPassword string `json:"hashed_password"`
	Email          string `json:"email"`
	TimeStamp      time.Time
}

type OtpStore struct {
	Otp       string `json:"otp"`
	Email     string `json:"email"`
	TimeStamp time.Time
}


var CacheService = &CacheStore{}
