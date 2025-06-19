package websocket

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/url"
)

type Client interface {
	Close() error
	Send(v any) error
	Read(v any) error
}
type client struct {
	*websocket.Conn
	host string
	opt  dailOption
}

func NewClient(host string, opts ...DailOption) Client {
	o := newDailOptions(opts...)
	c := &client{
		host: host,
		opt:  o,
	}
	dail, err := c.dail()
	if err != nil {
		panic(err)
	}
	c.Conn = dail
	fmt.Println("conn::", c.Conn)
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
func (c *client) Close() error {
	return c.Conn.Close()
}
func (c *client) Send(v any) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return c.Conn.WriteMessage(websocket.TextMessage, data)

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
