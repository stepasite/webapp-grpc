package common

import (
	"time"
	"sync"
)

var Sessions = sync.Map{}

type Session struct {
        Email string
        ExpiresAt time.Time
}

func (s Session) IsExpired() bool {
        return s.ExpiresAt.Before(time.Now())
}

