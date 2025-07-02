package svc

import (
	"github.com/ljp-lachouchou/chan_xin/apps/user/api/internal/config"
	"github.com/ljp-lachouchou/chan_xin/apps/user/rpc/userservice"
	"github.com/ljp-lachouchou/chan_xin/pkg/lmiddleware"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	userservice.UserService
	*redis.Redis
	LimitMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		UserService:     userservice.NewUserService(zrpc.MustNewClient(c.UserRpc)),
		Redis:           redis.MustNewRedis(c.Redisx),
		LimitMiddleware: lmiddleware.NewLimitMiddleware(c.Redisx).TokenLimitHandler(1, 100),
	}
}
