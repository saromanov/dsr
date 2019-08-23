package dsr

import (
	"net/http"
	"sync"

	"github.com/sirupsen/logrus"
)

type DSR struct {
	locker *lock
	hasher Hasher
	config *Config
	wg     sync.WaitGroup
	server *http.Server
	entry  *logrus.Entry
}

// New provides initialization of the dsr app
func New(conf *Config) (*DSR, error) {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	if conf == nil {
		conf = DefaultConfig()
	}
	dsr := &DSR{
		config: conf,
		entry:  &logrus.Entry{},
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

	select {
	case err := <-errCh:
		return err
	default:
	}

	return nil
}
