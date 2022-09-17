package db

import (
	"firebase.google.com/go/v4/auth"
	"sync"
)

type TokenUserRecord struct {
	UserRecord *auth.UserRecord
	Token      *auth.Token
}

type UserCache struct {
	m    map[string]TokenUserRecord
	lock sync.RWMutex
}

func NewUserCache() *UserCache {
	return &UserCache{m: map[string]TokenUserRecord{}, lock: sync.RWMutex{}}
}

func (uc *UserCache) Get(accessKey string) *auth.UserRecord {
	uc.lock.RLock()
	defer uc.lock.RUnlock()
	return uc.m[accessKey].UserRecord
}

func (uc *UserCache) Put(accessKey string, token *auth.Token, userRecord *auth.UserRecord) {
	uc.lock.Lock()
	defer uc.lock.Unlock()
	uc.m[accessKey] = TokenUserRecord{Token: token, UserRecord: userRecord}
}
