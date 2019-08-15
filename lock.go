package dsr

import (
	"sync"
)

type lock struct {
	mu sync.Mutex
}

func (l *lock) lock() {
	l.mu.Lock()
}

func (l *lock) unlock() {
	l.mu.Unlock()
}
