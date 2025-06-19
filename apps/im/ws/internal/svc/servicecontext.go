package svc

import (
	"github.com/ljp-lachouchou/chan_xin/apps/im/ws/internal/config"
	"github.com/ljp-lachouchou/chan_xin/apps/task/mq/client"
)

type ServiceContext struct {
	config.Config
	client.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Client: client.NewMsgChatTransferClient(c.MsgChatTransfer.Addrs, c.MsgChatTransfer.Topic),
	}
}
