package dynamicsmodels

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PostLikesModel = (*customPostLikesModel)(nil)

type (
	// PostLikesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPostLikesModel.
	PostLikesModel interface {
		postLikesModel
	}

	customPostLikesModel struct {
		*defaultPostLikesModel
	}
)

// NewPostLikesModel returns a model for the database table.
func NewPostLikesModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) PostLikesModel {
	return &customPostLikesModel{
		defaultPostLikesModel: newPostLikesModel(conn, c, opts...),
	}
}
