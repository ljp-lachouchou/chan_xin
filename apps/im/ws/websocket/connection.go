package websocket

import (
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
	"time"
)

type Connection struct {
	*websocket.Conn
	idleMu      sync.Mutex
	s           *Server
	idle        time.Time //当前空闲时间
	maxConnIdle time.Duration
	done        chan struct{}
}

func (c *Connection) readMessage() (messageType int, p []byte, err error) {
	messageType, p, err = c.Conn.ReadMessage()
	c.idleMu.Lock()
	defer c.idleMu.Unlock()
	c.idle = time.Time{}
	return
}
func (c *Connection) WriteMessage(messageType int, data []byte) error {
	c.idleMu.Lock()
	defer c.idleMu.Unlock()
	c.idle = time.Now()
	return c.Conn.WriteMessage(messageType, data)
}
func NewConnection(s *Server, w http.ResponseWriter, r *http.Request) *Connection {
	c, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		s.Errorf("upgrade.Upgrade err:%v", err)
		return nil
	}
	conn := &Connection{
		Conn:        c,
		s:           s,
		idle:        time.Now(),
		maxConnIdle: s.opt.maxConnIdle,
		done:        make(chan struct{}),
	}
	go conn.KeepAlive()
	return conn
}
func (c *Connection) KeepAlive() {
	idleTimer := time.NewTimer(c.maxConnIdle)
	defer idleTimer.Stop()
	for {
		select {
		case <-idleTimer.C:
			c.idleMu.Lock()
			idle := c.idle
			if idle.IsZero() {
				c.idleMu.Unlock()
				idleTimer.Reset(c.maxConnIdle)
				continue
			}
			val := c.maxConnIdle - time.Since(idle)
			c.idleMu.Unlock()
			if val <= 0 {
				// The connection has been idle for a duration of keepalive.MaxConnectionIdle or more.
				// Gracefully close the connection.
				c.s.Close(c)
				return

			}
			idleTimer.Reset(val)
		case <-c.done:
			return
		}
	}
}
func (c *Connection) Close() error {
	select {
	case <-c.done:
	default:
		close(c.done)
	}
	return c.Conn.Close()
}
