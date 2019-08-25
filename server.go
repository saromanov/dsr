package dsr

import (
	"context"
	"net"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

type Server struct {
	addr            string
	keepAlivePeriod time.Duration
	logger          *log.Logger
	wg              sync.WaitGroup
	listener        net.Listener
	ctx             context.Context
	cancel          context.CancelFunc
}

func NewServer(addr string, logger *log.Logger, keepalivePeriod time.Duration) *Server {
	ctx, cancel := context.WithCancel(context.Background())
	return &Server{
		addr:            addr,
		keepAlivePeriod: keepalivePeriod,
		logger:          logger,
		ctx:             ctx,
		cancel:          cancel,
	}
}
