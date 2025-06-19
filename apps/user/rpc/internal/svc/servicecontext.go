package svc

import (
	"github.com/ljp-lachouchou/chan_xin/apps/user/rpc/internal/config"
	"github.com/ljp-lachouchou/chan_xin/apps/user/usermodels"
	"github.com/ljp-lachouchou/chan_xin/deploy/constant"
	"github.com/ljp-lachouchou/chan_xin/pkg/ctxdata"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
)

type ServiceContext struct {
	Config config.Config
	usermodels.UsersModel
	*redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		Redis:      redis.MustNewRedis(c.Redisx),
		UsersModel: usermodels.NewUsersModel(conn, c.Cache),
	}
}
func (svc *ServiceContext) SetRootToken() error {
	//生成jwt
	systemToken, err := ctxdata.GetToken(svc.Config.Jwt.AccessSecret, time.Now().Unix(), svc.Config.Jwt.AccessExpire, constant.SYSTEM_ROOT_ID)
	if err != nil {
		return err
	}
	//写入redis
	return svc.Redis.Set(constant.REDIS_SYSTEM_ROOT_TOKEN, systemToken)
}
