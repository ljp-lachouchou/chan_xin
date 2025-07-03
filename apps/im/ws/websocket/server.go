package websocket

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/ljp-lachouchou/chan_xin/pkg/ldefault"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
	"net/http"
	"sync"
	"time"
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
	*threading.TaskRunner
}

func NewServer(addr string, opts ...ServerOption) *Server {
	opt := newServerOption(opts...)
	return &Server{
		addr:           addr,
		patten:         opt.patten,
		Authentication: opt.Authentication,
		routes:         make(map[string]HandlerFunc),
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true }, //websocket解决跨域
		},
		opt:        opt,
		connToUser: make(map[*Connection]string),
		userToConn: make(map[string]*Connection),
		Logger:     logx.WithContext(context.Background()),
		TaskRunner: threading.NewTaskRunner(opt.concurrency),
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
	s.RWMutex.RLock()
	defer s.RWMutex.RUnlock()
	for _, conn := range conns {
		if _, ok := s.connToUser[conn]; !ok {
			continue
		}
		if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
			continue
		}
	}
	return nil
}

func (s *Server) SendPingMessage() error {
	for _, conn := range s.userToConn {
		if conn.Uid == ldefault.SYSTEM_REDIS_UID {
			continue
		}
		go func(c *Connection) {
			s.RWMutex.RLock()
			err := c.WriteMessage(websocket.PingMessage, []byte(c.Uid))
			s.RWMutex.RUnlock()
			if err != nil {
				s.Errorf("Ping 失败: UID=%s, Err=%v", c.Uid, err)
				c.Close()
				return
			}
			return
		}(conn)
	}
	return nil
}
func (s *Server) GetConn(uid string) *Connection {
	s.RWMutex.RLock()
	defer s.RWMutex.RUnlock()
	return s.userToConn[uid]
}
func (s *Server) GetUid(conn *Connection) string {
	s.RWMutex.RLock()
	defer s.RWMutex.RUnlock()
	return s.connToUser[conn]
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
	conn.Uid = s.GetUid(conn)
	go s.handleWrite(conn)
	go s.readAck(conn)
	for {
		_, msg, err := conn.ReadMessage()

		if err != nil {
			s.Errorf(fmt.Sprintf("websocket conn:%s read message err: %v msg %v", conn.Uid, err, string(msg)))
			s.Close(conn)
			return
		}
		var message Message
		if err := json.Unmarshal(msg, &message); err != nil {
			s.Errorf("json unmarshal err: %v,msg %v", err, string(msg))
			s.Close(conn)
			return
		}
		if conn.Uid == ldefault.SYSTEM_REDIS_UID {
			fmt.Println("root dail")
			conn.message <- &message
		} else {
			conn.appendAckMq(&message)
		}
	}
}
func (s *Server) handleWrite(conn *Connection) {
	for {
		select {
		case message := <-conn.message:
			switch message.FrameType {
			case FrameData:
				handler := s.routes[message.Method]
				if handler == nil {
					s.Send(NewErrMessage(errors.New(fmt.Sprintf("找不到方法 %v", message.Method))))
				} else {
					handler(s, conn, message)
					if _, ok := conn.readAckMq[message.Id]; ok {
						conn.messageMu.Lock()
						delete(conn.readAckMq, message.Id)
						conn.messageMu.Unlock()
					}
				}
			case FramePing:
				s.Send(NewPingMessage(), conn)
			}
		case <-conn.done:
			return
		}
	}
}
func (s *Server) readAck(conn *Connection) {
	for {
		select {
		case <-conn.done:
			s.Info("close message ack Uid %v", conn.Uid)
			return
		default:
		}
		conn.messageMu.Lock()
		if len(conn.ackMessages) == 0 {
			conn.messageMu.Unlock()
			time.Sleep(100 * time.Millisecond)
			continue
		}
		message := conn.ackMessages[0]
		s.Send(&Message{
			FrameType: FrameAck,
			Id:        message.Id,
			Seq:       message.Seq + 1,
		}, conn)
		//进行业务处理
		//消息从队列中移除
		conn.ackMessages = conn.ackMessages[1:]
		conn.messageMu.Unlock()
		conn.message <- message
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

func (s *Server) GetConns(ids []string) []*Connection {
	s.RWMutex.RLock()
	defer s.RWMutex.RUnlock()

	var res []*Connection
	if len(ids) == 0 {
		// 获取全部
		res = make([]*Connection, 0, len(s.userToConn))
		for _, uid := range s.userToConn {
			res = append(res, uid)
		}
	} else {
		// 获取部分
		res = make([]*Connection, 0, len(ids))
		for _, conn := range ids {
			res = append(res, s.userToConn[conn])
		}
	}
	return res
}

func (s *Server) GetUsers(conns ...*Connection) []string {
	s.RWMutex.RLock()
	defer s.RWMutex.RUnlock()

	var res []string
	if len(conns) == 0 {
		// 获取全部
		res = make([]string, 0, len(s.connToUser))
		for _, uid := range s.connToUser {
			res = append(res, uid)
		}
	} else {
		// 获取部分
		res = make([]string, 0, len(conns))
		for _, conn := range conns {
			res = append(res, s.connToUser[conn])
		}
	}

	return res
}
