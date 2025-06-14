package svc

import (
	"github.com/ljp-lachouchou/chan_xin/apps/user/rpc/internal/config"
	"github.com/ljp-lachouchou/chan_xin/apps/user/usermodels"
	"github.com/ljp-lachouchou/chan_xin/pkg/ctxdata"
	"github.com/ljp-lachouchou/chan_xin/pkg/ldefault"
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
	systemToken, err := ctxdata.GetToken(svc.Config.Jwt.AccessSecret, time.Now().Unix(), ldefault.DEFAULT_EXP, ldefault.SYSTEM_REDIS_UID)
	if err != nil {
		return err
	}
	//写入redis
	return svc.Redis.Setex(ldefault.SYSTEM_REDIS_TOPNKEN_KEY, systemToken, ldefault.DEFAULT_EXP)
}
