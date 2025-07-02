package svc

import (
	"github.com/ljp-lachouchou/chan_xin/apps/im/api/internal/config"
	"github.com/ljp-lachouchou/chan_xin/apps/im/rpc/imclient"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

var retryPolicy = `{
	"methodConfig" : [{
		"name":[{
			"service":"im.Im"
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
	Config config.Config
	imclient.Im
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Im:     imclient.NewIm(zrpc.MustNewClient(c.ImRpc, zrpc.WithDialOption(grpc.WithDefaultServiceConfig(retryPolicy)))),
	}
}
