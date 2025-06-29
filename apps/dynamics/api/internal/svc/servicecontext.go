package svc

import (
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/api/internal/config"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/dynamicsclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	dynamicsclient.Dynamics
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		Dynamics: dynamicsclient.NewDynamics(zrpc.MustNewClient(c.DynamicsRpc)),
	}
}
