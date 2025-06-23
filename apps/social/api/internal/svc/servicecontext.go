package svc

import (
	"github.com/ljp-lachouchou/chan_xin/apps/im/rpc/imclient"
	"github.com/ljp-lachouchou/chan_xin/apps/social/api/internal/config"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/socialservice"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	socialservice.SocialService
	imclient.Im
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		SocialService: socialservice.NewSocialService(zrpc.MustNewClient(c.SocialRpc)),
		Im:            imclient.NewIm(zrpc.MustNewClient(c.ImRpc)),
	}
}
