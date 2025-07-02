package svc

import (
	"github.com/ljp-lachouchou/chan_xin/apps/im/ws/internal/config"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/socialservice"
	"google.golang.org/grpc"

	"github.com/ljp-lachouchou/chan_xin/apps/task/mq/client"
	"github.com/zeromicro/go-zero/zrpc"
)

var retryPolicy = `{
	"methodConfig" : [{
		"name":[{
			"service":"im.Ws"
		}],
		"waitForReady":true,
		"retryPolicy": {
			"maxAttempts": 5,
			"initialBackoff": "0.001s",
			"maxBackoff": "0.002s",
			"backoffMultiplier": 1.0,
			"retryableStatusCodes":["DEADLINE_EXCEEDED"]
		}
	}]
}`

type ServiceContext struct {
	config.Config
	client.Client
	socialservice.SocialService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		Client:        client.NewMsgChatTransferClient(c.MsgChatTransfer.Addrs, c.MsgChatTransfer.Topic),
		SocialService: socialservice.NewSocialService(zrpc.MustNewClient(c.SocialRpc, zrpc.WithDialOption(grpc.WithDefaultServiceConfig(retryPolicy)))),
	}
}
