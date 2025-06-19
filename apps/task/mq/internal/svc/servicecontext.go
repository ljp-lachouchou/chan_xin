package svc

import (
	"fmt"
	"github.com/ljp-lachouchou/chan_xin/apps/im/immodels"
	"github.com/ljp-lachouchou/chan_xin/apps/im/ws/websocket"
	"github.com/ljp-lachouchou/chan_xin/apps/task/mq/internal/config"
	"github.com/ljp-lachouchou/chan_xin/deploy/constant"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"net/http"
)

type ServiceContext struct {
	config.Config
	immodels.ChatLogModel
	immodels.ConversationModel
	websocket.Client
	*redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	svc := &ServiceContext{
		Config:            c,
		ChatLogModel:      immodels.MustNewChatLogModel(c.Mongo.Url, c.Mongo.Db),
		ConversationModel: immodels.MustNewConversationModel(c.Mongo.Url, c.Mongo.Db),
		Redis:             redis.MustNewRedis(c.Redisx),
	}
	token, err := svc.GetRootToken()
	if err != nil {
		panic(err)
	}
	header := http.Header{}
	fmt.Println("token:", token)
	header.Set("Authorization", token)
	svc.Client = websocket.NewClient(c.Ws.Host, websocket.WithClientHeader(header))
	return svc
}
func (c *ServiceContext) GetRootToken() (string, error) {
	return c.Redis.Get(constant.REDIS_SYSTEM_ROOT_TOKEN)

}
