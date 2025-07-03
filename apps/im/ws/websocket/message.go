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
	FrameType FrameType `json:"frameType"` //一般为0，与gorilla的frametype一样
	Id        string    `json:"id"`        //消息唯一标识
	FromId    string    `json:"fromId"`    //发送者
	ToId      string    `json:"toId"`      //接受者
	AckTime   time.Time `json:"ackTime"`   //acktime,不需要填写，为不严谨ack

	Method string      `json:"method"` //方法: 填写 conversation.chat
	Seq    int         `json:"seq"`    //
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

/*
	{
	    "frameType":0,
	    "id":"12832a8sdas22as3",//唯一id
	    "toId":"0x0000006000000002",//接收方iD
	    "fromId":"0x0000001000000001",//发送方iD
	    "method":"conversation.chat",
	    "data":{
	        "sendId":"0x0000001000000001",//聊天消息发送方iD
	        "recvId":"0x0000006000000002",//聊天消息接收方iD
	        "chatType":1,//0 -私聊 1-群聊
	        "msg":{
	            "msgType":1, //发送消息类型，0-文本，1-图片，2-大型文件
	            "msgContent":"hello"//依据消息不同，发送不同字符串内容
	        }
	    }
	}
*/
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
