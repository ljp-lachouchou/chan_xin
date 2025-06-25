package dynamicsmodels

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CommentRepliesModel = (*customCommentRepliesModel)(nil)

type (
	// CommentRepliesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCommentRepliesModel.
	CommentRepliesModel interface {
		commentRepliesModel
	}

	customCommentRepliesModel struct {
		*defaultCommentRepliesModel
	}
)

// NewCommentRepliesModel returns a model for the database table.
func NewCommentRepliesModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) CommentRepliesModel {
	return &customCommentRepliesModel{
		defaultCommentRepliesModel: newCommentRepliesModel(conn, c, opts...),
	}
}
