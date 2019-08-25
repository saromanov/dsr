package dsr

import (
	"context"
	"sync"

	"github.com/sirupsen/logrus"
)

type DSR struct {
	locker    *lock
	hasher    Hasher
	config    *Config
	wg        sync.WaitGroup
	server    *Server
	entry     *logrus.Entry
	discovery *discovery
	ctx       context.Context
	cancel    context.CancelFunc
	name      string
}

// New provides initialization of the dsr app
func New(conf *Config) (*DSR, error) {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	if conf == nil {
		conf = DefaultConfig()
	}

	ctx, cancel := context.WithCancel(context.Background())
	dsr := &DSR{
		config: conf,
		entry:  &logrus.Entry{},
		ctx:    ctx,
		cancel: cancel,
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

	return dsr.listenAndServe()
}

func (s *Server) listenAndServe() error {
	close(s.ch)

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			select {
			case <-s.ctx.Done():
				return nil
			default:
			}
			s.logger.Printf("[DEBUG] Failed to accept TCP connection: %v", err)
			continue
		}
	}
}
