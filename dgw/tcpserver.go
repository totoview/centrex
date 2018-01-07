package dgw

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

// TCPServer implements data source gateway over TCP
type TCPServer struct {
	logger    log.Logger
	listener  net.Listener
	sessionID uint64

	stop chan struct{}
	done chan struct{}
}

const serverName = "DgwTcpServer"

// NewTCPServer creates a new TCPServer over regular TCP.
func NewTCPServer(addr string) (*TCPServer, error) {
	listener, err := net.Listen("tcp4", addr)
	if err != nil {
		return nil, err
	}
	return &TCPServer{listener: listener, logger: clog.Logger()}, nil
}

// NewSecureTCPServer creates a new TCPServer over TLS.
func NewSecureTCPServer(addr string, cert, key string) (*TCPServer, error) {
	kp, err := tls.LoadX509KeyPair(cert, key)
	if err != nil {
		return nil, err
	}

	tlsCfg := &tls.Config{Certificates: []tls.Certificate{kp}, CipherSuites: []uint16{tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA}}

	listener, err := tls.Listen("tcp4", addr, tlsCfg)
	if err != nil {
		return nil, err
	}
	return &TCPServer{listener: listener, logger: clog.Logger()}, nil
}

// Start runs TCPServer
func (s *TCPServer) Start() error {
	s.stop, s.done = make(chan struct{}), make(chan struct{})

	var wg sync.WaitGroup

	var wgSessions sync.WaitGroup
	stopSessions := make(chan struct{})

	go func() {
		level.Info(s.logger).Log("start", serverName)
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
			session := NewTCPSession(id, conn, stopSessions, &wgSessions, s.logger)
			session.Start()
		}
	}()
	return nil
}

// Stop shuts down TCPServer
func (s *TCPServer) Stop() {
	close(s.stop)
	<-s.done
	level.Info(s.logger).Log("stop", serverName)
}
