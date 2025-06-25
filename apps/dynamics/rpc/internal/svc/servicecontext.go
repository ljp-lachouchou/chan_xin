package svc

import (
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/dynamicsmodels"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	dynamicsmodels.CommentsModel
	dynamicsmodels.CommentRepliesModel
	dynamicsmodels.PostLikesModel
	dynamicsmodels.NotificationsModel
	dynamicsmodels.PostsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:              c,
		CommentsModel:       dynamicsmodels.NewCommentsModel(conn, c.Cache),
		CommentRepliesModel: dynamicsmodels.NewCommentRepliesModel(conn, c.Cache),
		PostLikesModel:      dynamicsmodels.NewPostLikesModel(conn, c.Cache),
		NotificationsModel:  dynamicsmodels.NewNotificationsModel(conn, c.Cache),
		PostsModel:          dynamicsmodels.MustNewPostsModel(c.Mongo.Url, c.Mongo.Db),
	}
}
