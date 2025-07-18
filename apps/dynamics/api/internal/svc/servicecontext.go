package svc

import (
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/api/internal/config"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/dynamicsclient"
	"github.com/ljp-lachouchou/chan_xin/pkg/lmiddleware"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

var retryPolicy = `{
	"methodConfig" : [{
		"name":[{
			"service":"dynamics.Dynamics"
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
	dynamicsclient.Dynamics
	LimitMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		Dynamics:        dynamicsclient.NewDynamics(zrpc.MustNewClient(c.DynamicsRpc, zrpc.WithDialOption(grpc.WithDefaultServiceConfig(retryPolicy)))),
		LimitMiddleware: lmiddleware.NewLimitMiddleware(c.Redisx).TokenLimitHandler(1, 100),
	}
}
