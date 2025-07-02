package ws

import "github.com/ljp-lachouchou/chan_xin/deploy/constant"

type (
	Msg struct {
		MsgType    constant.MType `mapstructure:"msgType"`
		MsgContent string         `mapstructure:"msgContent"`
	}
	Chat struct {
		ConversationId string            `mapstructure:"conversationId"`
		ChatType       constant.ChatType `mapstructure:"chatType"`
		SendId         string            `mapstructure:"sendId"`
		RecvId         string            `mapstructure:"recvId"`
		Msg            Msg               `mapstructure:"msg"`
		SendTime       int64             `mapstructure:"sendTime"`
	}
	Push struct {
		ConversationId string            `json:"conversationId"`
		SendId         string            `json:"sendId"`
		RecvId         string            `json:"recvId"`
		ChatType       constant.ChatType `json:"chatType"`
		MsgType        constant.MType    `json:"msgType"`
		MsgContent     string            `json:"msgContent"`
		SendTime       int64             `json:"sendTime"`
	}
)
