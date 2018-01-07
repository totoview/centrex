package tcp

import (
	"crypto/tls"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	clog "github.com/totoview/centrex/log"
)

// Server implements data source gateway over TCP
type Server struct {
	name      string
	logger    log.Logger
	listener  net.Listener
	sessionID uint64

	stop chan struct{}
	done chan struct{}
}

// NewServer creates a new Server over regular TCP.
func NewServer(name, addr string) (*Server, error) {
	listener, err := net.Listen("tcp4", addr)
	if err != nil {
		return nil, err
	}
	return &Server{name: name, listener: listener, logger: clog.Logger()}, nil
}

// NewSecureServer creates a new Server over TLS.
func NewSecureServer(addr string, cert, key string) (*Server, error) {
	kp, err := tls.LoadX509KeyPair(cert, key)
	if err != nil {
		return nil, err
	}

	tlsCfg := &tls.Config{Certificates: []tls.Certificate{kp}, CipherSuites: []uint16{tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA}}

	listener, err := tls.Listen("tcp4", addr, tlsCfg)
	if err != nil {
		return nil, err
	}
	return &Server{listener: listener, logger: clog.Logger()}, nil
}

// Start runs Server
func (s *Server) Start() error {
	s.stop, s.done = make(chan struct{}), make(chan struct{})

	var wg sync.WaitGroup

	var wgSessions sync.WaitGroup
	stopSessions := make(chan struct{})

	go func() {
		level.Info(s.logger).Log("start", s.name)
		for {
			select {
			case <-s.stop:
				s.listener.Close()
				wg.Wait()
				s.done <- struct{}{}
				return
			}
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		level.Info(s.logger).Log("listen", s.listener.Addr())
		for {
			conn, err := s.listener.Accept()
			if err != nil {
				if netErr, ok := err.(net.Error); ok && netErr.Temporary() {
					<-time.After(100 * time.Millisecond)
					continue
				}
				close(stopSessions)
				wgSessions.Wait()
				<-s.stop
				return
			}
			id := atomic.AddUint64(&s.sessionID, 1)
			session := NewClientSession(id, conn, stopSessions, &wgSessions, s.logger)
			session.Start()
		}
	}()
	return nil
}

// Stop shuts down Server
func (s *Server) Stop() {
	close(s.stop)
	<-s.done
	level.Info(s.logger).Log("stop", s.name)
}
