package tcp

import (
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gogo/protobuf/proto"
	"github.com/totoview/centrex/pb"
)

type inState int

const (
	statLength inState = iota
	statData
)

type inBuf struct {
	buf    []byte
	stat   inState
	datLen int
	offset int
	size   int
}

func newInBuf(initCap int) *inBuf {
	return &inBuf{buf: make([]byte, initCap), stat: statLength}
}

func (s *inBuf) Read(c net.Conn) (int, error) {
	if s.offset+s.size >= cap(s.buf) {
		buf := make([]byte, cap(s.buf)+4096)
		copy(buf, s.buf[s.offset:(s.offset+s.size)])
		s.buf, s.offset = buf, 0
	}

	buf := s.buf[s.offset+s.size:]
	n, err := c.Read(buf)
	if err == nil {
		s.size += n
	}
	return n, err
}

// return nil if not enough data
func (s *inBuf) Decode() (*pb.CentrexMsg, error) {
	var msg pb.CentrexMsg
	data := s.buf[s.offset:(s.offset + s.size)]

	if s.stat == statLength {
		cnt, n := proto.DecodeVarint(data)
		if n == 0 {
			return nil, nil
		}

		s.datLen, s.stat = int(cnt), statData
		s.offset += n
		data = data[n:]
	}

	if s.size < s.datLen {
		return nil, nil
	}

	if err := proto.Unmarshal(data, &msg); err != nil {
		return nil, err
	}

	s.size, s.stat = s.size-s.datLen, statLength
	if s.size > 0 {
		copy(s.buf, s.buf[s.offset:(s.offset+s.size)])
	}
	s.offset = 0
	return &msg, nil
}

// ClientSession manages a client connection over TCP.
type ClientSession struct {
	ID     uint64
	conn   net.Conn
	logger log.Logger
	stop   chan struct{}
	done   *sync.WaitGroup
}

// NewClientSession creates a new instance of ClientSession.
func NewClientSession(id uint64, conn net.Conn, stop chan struct{}, done *sync.WaitGroup, logger log.Logger) *ClientSession {
	return &ClientSession{ID: id, conn: conn, stop: stop, done: done, logger: logger}
}

// Start runs ClientSession.
func (s *ClientSession) Start() {

	level.Info(s.logger).Log("op", "start", "Clientsession", s.ID, "from", s.conn.RemoteAddr())

	s.done.Add(1)
	closed, outDone := make(chan struct{}), make(chan struct{})

	// outgoing traffic
	go func() {
		for {
			select {
			case <-s.stop:
				s.conn.Close()
			case <-closed:
				outDone <- struct{}{}
				return
			}
		}
	}()

	// incoming traffic and control
	go func() {
		defer func() { level.Info(s.logger).Log("op", "stop", "Clientsession", s.ID) }()
		defer s.done.Done()
		inputBuf := newInBuf(4096)
		for {
			_, err := inputBuf.Read(s.conn)
			if err != nil {
				if netErr, ok := err.(net.Error); ok && netErr.Temporary() {
					<-time.After(100 * time.Millisecond)
					continue
				}
				closed <- struct{}{}
				<-outDone
				return
			}
			for {
				msg, err := inputBuf.Decode()
				if msg != nil {
					fmt.Printf("==== [in] %v\n", msg)
					switch msg.Type {
					case pb.CentrexMsg_Login:
						s.handleLogin(msg.Login)
					}
				} else {
					if err != nil {
						level.Error(s.logger).Log("err", fmt.Sprintf("Failed to parse input: %s", err.Error()))
						s.conn.Close()
					}
					break
				}
			}
		}
	}()
}

func (s *ClientSession) handleLogin(req *pb.Login) {

}
