package svc

import (
	"github.com/ljp-lachouchou/chan_xin/apps/im/ws/internal/config"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/socialservice"
	"github.com/ljp-lachouchou/chan_xin/apps/task/mq/client"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	config.Config
	client.Client
	socialservice.SocialService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		Client:        client.NewMsgChatTransferClient(c.MsgChatTransfer.Addrs, c.MsgChatTransfer.Topic),
		SocialService: socialservice.NewSocialService(zrpc.MustNewClient(c.SocialRpc)),
	}
}
