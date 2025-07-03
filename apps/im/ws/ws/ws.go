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
		Id             string            `mapstructure:"id"`
		ConversationId string            `mapstructure:"conversationId"`
		SendId         string            `mapstructure:"sendId"`
		RecvId         string            `mapstructure:"recvId"`
		RecvIds        []string          `mapstructure:"recvIds"`
		ChatType       constant.ChatType `mapstructure:"chatType"`
		MsgType        constant.MType    `mapstructure:"msgType"`
		MsgContent     string            `mapstructure:"msgContent"`
		SendTime       int64             `mapstructure:"sendTime"`
	}
)
