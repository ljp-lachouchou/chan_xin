package svc

import (
	"github.com/ljp-lachouchou/chan_xin/apps/im/api/internal/config"
	"github.com/ljp-lachouchou/chan_xin/apps/im/rpc/imclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	imclient.Im
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Im:     imclient.NewIm(zrpc.MustNewClient(c.ImRpc)),
	}
}
