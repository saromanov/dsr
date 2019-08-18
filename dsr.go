package dsr

import (
	"sync"
)

type DSR struct {
	locker *lock
	hasher Hasher
	config *Config
	wg     sync.WaitGroup
}

// New provides initialization of the dsr app
func New(conf *Config) (*DSR, error) {
	if conf == nil {
		conf = DefaultConfig()
	}
	dsr := &DSR{
		config: conf,
	}
	return dsr, nil
}

// Start provides starting of the app
func (dsr *DSR) Start() error {
	errCh := make(chan error, 1)
	dsr.wg.Add(1)
	go func() {
		defer dsr.wg.Done()
		errCh <- dsr.server.ListenAndServe()
	}()

	<-dsr.server.StartCh
}
