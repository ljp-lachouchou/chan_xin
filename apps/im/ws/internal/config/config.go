package config

import (
	"github.com/zeromicro/go-zero/core/service"
)

type Config struct {
	service.ServiceConf
	ListenOn string
	Mongo    struct {
		Url string
		Db  string
	}
	Jwt struct {
		AccessSecret string
		AccessExpire int
	}
	MsgChatTransfer struct {
		Topic string
		Addrs []string
	}
}
