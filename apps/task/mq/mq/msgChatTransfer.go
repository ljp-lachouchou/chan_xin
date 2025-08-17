package mq

import (
	"github.com/ljp-lachouchou/chan_xin/deploy/constant"
)

type MsgChatTransfer struct {
	Id          string            `json:"id"`
	ConversationId string            `json:"conversationId"`
	SendId         string            `json:"sendId"`
	RecvId         string            `json:"recvId"`
	ChatType       constant.ChatType `json:"chatType"`
	MsgType        constant.MType    `json:"msgType"`
	MsgContent     string            `json:"msgContent"`
	SendTime       int64             `json:"sendTime"`
}
