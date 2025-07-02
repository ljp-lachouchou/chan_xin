package svc

import (
	"github.com/ljp-lachouchou/chan_xin/apps/im/immodels"
	"github.com/ljp-lachouchou/chan_xin/apps/im/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	immodels.ChatLogModel
	immodels.ConversationModel
	immodels.ConversationsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:             c,
		ChatLogModel:       immodels.MustNewChatLogModel(c.Mongo.Url, c.Mongo.Db),
		ConversationModel:  immodels.MustNewConversationModel(c.Mongo.Url, c.Mongo.Db),
		ConversationsModel: immodels.MustNewConversationsModel(c.Mongo.Url, c.Mongo.Db),
	}
}
