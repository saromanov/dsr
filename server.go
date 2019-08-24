package dsr

import (
	"time"
	"context"
	log "github.com/sirupsen/logrus"
	"sync"
	"net"
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
	if logger == nil {
		logger = log.New(os.Stderr, "", log.LstdFlags)
	}
	return &Server{
		operations:      operations{m: make(map[protocol.OpCode]protocol.Operation)},
		addr:            addr,
		keepAlivePeriod: keepalivePeriod,
		logger:          logger,
		StartCh:         make(chan struct{}),
		ctx:             ctx,
		cancel:          cancel,
	}
}
