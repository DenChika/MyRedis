package commands

import (
	"sync"
	"time"
)

type RedisStorage struct {
	Vocabulary map[string]ExpiryValue
	Mu         *sync.RWMutex
}

var Storage RedisStorage

func init() {
	Storage = RedisStorage{
		Vocabulary: make(map[string]ExpiryValue),
		Mu:         &sync.RWMutex{},
	}
}

type ExpiryValue struct {
	Value      string
	ValidUntil time.Time
	IsDurable  bool
}
