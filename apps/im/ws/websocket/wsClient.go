package websocket

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/url"
	"time"
)

type Client interface {
	Close() error
	Send(v any) error
	Read(v any) error
}
type client struct {
	*websocket.Conn
	host  string
	opt   dailOption
	done  chan struct{}
	timer *time.Timer
}

func NewClient(host string, opts ...DailOption) Client {
	o := newDailOptions(opts...)
	c := &client{
		host:  host,
		opt:   o,
		done:  make(chan struct{}),
		timer: time.NewTimer(10 * time.Second),
	}
	dail, err := c.dail()
	if err != nil {
		panic(err)
	}
	c.Conn = dail
	fmt.Println("conn::", c.Conn)
	go c.pingServer(c.Conn)
	return c
}
func (c *client) dail() (*websocket.Conn, error) {
	u := url.URL{
		Scheme: "ws",
		Host:   c.host,
		Path:   c.opt.patten,
	}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), c.opt.header)
	return conn, err
}
func (c *client) pingServer(conn *websocket.Conn) {
	defer c.timer.Stop()
	for {
		select {
		case <-c.timer.C:
			msg := Message{
				FrameType: FrameData,
				Method:    "user.online",
			}
			data, err := json.Marshal(msg)
			fmt.Println("pingserver start:", string(data))
			if err != nil {
				fmt.Println("pingServer json marshal message error:", err)
				return
			}

			err = conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				fmt.Println("pingServer write ping message error:", err)
				c.Close()
				dail, err := c.dail()
				if err != nil {
					panic(err)
				}
				c.Conn = dail
				continue
			}
			c.timer.Reset(3 * time.Second)
		case <-c.done:
			return
		}
	}
}
func (c *client) Close() error {
	select {
	case <-c.done:
	default:
		close(c.done)
	}
	return c.Conn.Close()
}
func (c *client) Send(v any) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	err = c.Conn.WriteMessage(websocket.TextMessage, data)
	if err == nil {
		return nil
	}
	//todo: 增加一个重连发送
	conn, err := c.dail()
	if err != nil {
		panic(err)
	}
	c.Conn = conn
	return c.Conn.WriteMessage(websocket.TextMessage, data)
}
func (c *client) Read(v any) error {
	_, msg, err := c.Conn.ReadMessage()
	if err != nil {
		return err
	}
	return json.Unmarshal(msg, v)
}
