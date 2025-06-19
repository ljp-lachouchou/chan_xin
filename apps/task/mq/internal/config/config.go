package config

import (
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/redis"
)
import "github.com/zeromicro/go-queue/kq"

type Config struct {
	service.ServiceConf
	ListenOn        string
	MsgChatTransfer kq.KqConf
	Mongo           struct {
		Url string
		Db  string
	}
	Ws struct {
		Host string
	}
	Redisx redis.RedisConf
}
