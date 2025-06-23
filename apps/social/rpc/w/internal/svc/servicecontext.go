package svc

import "github.com/ljp-lachouchou/chan_xin/apps/social/rpc/w/internal/config"

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
