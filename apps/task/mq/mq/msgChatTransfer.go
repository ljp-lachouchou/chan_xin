package mq

import (
	"github.com/ljp-lachouchou/chan_xin/apps/im/ws/ws"
	"github.com/ljp-lachouchou/chan_xin/deploy/constant"
)

type MsgChatTransfer struct {
	ConversationId string            `json:"conversationId"`
	SendId         string            `json:"sendId"`
	RecvId         string            `json:"recvId"`
	ChatType       constant.ChatType `json:"chatType"`
	MsgType        constant.MType    `json:"msgType"`
	MsgContent     ws.MessageContent `json:"msgContent"`
	SendTime       int64             `json:"sendTime"`
}
