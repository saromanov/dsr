package dsr

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
)

const (
	protocol = "tcp"
)

// Server defines struct for the server
type Server struct {
	address         string
	ch              chan struct{}
	keepAlivePeriod time.Duration
	logger          *log.Logger
	wg              sync.WaitGroup
	listener        net.Listener
	ctx             context.Context
	cancel          context.CancelFunc
}

// NewServer provides starting of the server
func NewServer(addr string, logger *log.Logger, keepalivePeriod time.Duration) *Server {
	ctx, cancel := context.WithCancel(context.Background())
	return &Server{
		address:         addr,
		keepAlivePeriod: keepalivePeriod,
		logger:          logger,
		ctx:             ctx,
		cancel:          cancel,
		ch:              make(chan struct{}),
	}
}

// ListenAndServe provides starting of the server
func (s *Server) ListenAndServe() error {
	defer func() {
		select {
		case <-s.ch:
			return
		default:
		}
		close(s.ch)
	}()

	l, err := net.Listen(protocol, s.address)
	if err != nil {
		return xerrors.New(fmt.Sprintf("unable to listen server by the address: %s %v", s.address, err))
	}
	s.listener = l
	return nil
}
