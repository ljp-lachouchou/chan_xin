package websocket

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"sync"
)

type Server struct {
	addr   string
	patten string
	Authentication
	sync.RWMutex
	routes     map[string]HandlerFunc
	upgrader   websocket.Upgrader
	opt        serverOption
	connToUser map[*Connection]string
	userToConn map[string]*Connection
	logx.Logger
}

func NewServer(addr string, opts ...ServerOption) *Server {
	opt := newServerOption(opts...)
	return &Server{
		addr:           addr,
		patten:         opt.patten,
		Authentication: opt.Authentication,
		routes:         make(map[string]HandlerFunc),
		upgrader:       websocket.Upgrader{},
		opt:            opt,
		connToUser:     make(map[*Connection]string),
		userToConn:     make(map[string]*Connection),
		Logger:         logx.WithContext(context.Background()),
	}
}
func (s *Server) Send(msg interface{}, conns ...*Connection) error {
	if len(conns) == 0 {
		return nil
	}
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	for _, conn := range conns {
		if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
			return err
		}
	}
	return nil
}
func (s *Server) GetConn(uid string) *Connection {
	s.RWMutex.RLock()
	defer s.RWMutex.RUnlock()
	return s.userToConn[uid]
}
func (s *Server) AddRoutes(r []Route) {
	for _, route := range r {
		s.routes[route.Method] = route.Handler
	}
}
func (s *Server) WsServer(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			s.Errorf("server handler ws recover: %v", r)
		}
	}()
	conn := NewConnection(s, w, r)
	if conn == nil {
		s.Errorf("coon为空")
		return
	}
	if !s.Authentication.Auth(w, r) {
		s.Errorf("此链接不具备执行权限")
		s.Send(&Message{FrameType: FrameData, Data: fmt.Sprintf("此ws链接不具备执行权限")}, conn)
		conn.Close()
		return
	}
	s.addConn(conn, r)
	go s.handleConn(conn)

}
func (s *Server) handleConn(conn *Connection) {
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			s.Errorf("websocket conn read message err: %v msg %v", err, string(msg))
			s.Close(conn)
			return
		}
		var message Message
		if err := json.Unmarshal(msg, &message); err != nil {
			s.Errorf("json unmarshal err: %v,msg %v", err, string(msg))
			s.Close(conn)
			return
		}
		switch message.FrameType {
		case FrameData:
			handler := s.routes[message.Method]
			if handler == nil {
				s.Send(NewErrMessage(errors.New(fmt.Sprintf("找不到方法 %v", message.Method))))
			} else {
				handler(s, conn, &message)
			}
		case FramePing:
			s.Send(NewPingMessage(), conn)
		}
	}
}
func (s *Server) Start() {
	http.HandleFunc(s.patten, s.WsServer)
	s.Info(http.ListenAndServe(s.addr, nil))
}
func (s *Server) addConn(conn *Connection, r *http.Request) {
	uid := s.Authentication.GetUid(r)
	s.RWMutex.Lock()
	defer s.RWMutex.Unlock()
	if c, ok := s.userToConn[uid]; ok {
		c.Close()
	}
	s.connToUser[conn] = uid
	s.userToConn[uid] = conn
}
func (s *Server) Close(conn *Connection) {
	s.RWMutex.Lock()
	defer s.RWMutex.Unlock()
	uid, ok := s.connToUser[conn]
	if !ok {
		return
	}
	delete(s.connToUser, conn)
	delete(s.userToConn, uid)
	conn.Close()
}

func (s *Server) Stop() {
	s.Info("停止服务")
}
