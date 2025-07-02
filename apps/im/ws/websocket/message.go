package websocket

import (
	"github.com/google/uuid"
	"time"
)

type FrameType int

const (
	FrameData  FrameType = 0x0
	FramePing  FrameType = 0x1
	FrameAck   FrameType = 0x2
	FrameNoAck FrameType = 0x3
	FrameErr   FrameType = 0x9
)

type Message struct {
	FrameType FrameType `json:"frameType"`
	Id        string    `json:"id"`
	FromId    string    `json:"fromId"`
	ToId      string    `json:"toId"`
	AckTime   time.Time `json:"ackTime"`

	Method string      `json:"method"`
	Seq    int         `json:"seq"`
	Data   interface{} `json:"data"`
}

func NewMessage(fromId string, data interface{}) *Message {
	id := uuid.New().String()
	return &Message{
		FrameType: FrameData,
		Id:        id,
		FromId:    fromId,
		Data:      data,
	}
}
func NewAckMessage(fromId string, ackTime time.Time) *Message {
	return &Message{
		FrameType: 0,
		Id:        "",
		FromId:    "",
		ToId:      "",
		AckTime:   time.Time{},
		Method:    "",
		Seq:       0,
		Data:      nil,
	}
}
func NewMessageWithId(id, fromId, toId string, data interface{}) *Message {
	return &Message{
		FrameType: FrameData,
		Id:        id,
		FromId:    fromId,
		ToId:      toId,
		Data:      data,
	}
}
func NewPingMessage() *Message {
	return &Message{
		FrameType: FramePing,
	}
}
func NewErrMessage(err error) *Message {
	return &Message{
		FrameType: FrameErr,
		Data:      err.Error(),
	}
}
