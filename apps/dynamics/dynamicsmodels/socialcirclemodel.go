package dynamicsmodels

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SocialCircleModel = (*customSocialCircleModel)(nil)

type (
	// SocialCircleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSocialCircleModel.
	SocialCircleModel interface {
		socialCircleModel
	}

	customSocialCircleModel struct {
		*defaultSocialCircleModel
	}
)

// NewSocialCircleModel returns a model for the database table.
func NewSocialCircleModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) SocialCircleModel {
	return &customSocialCircleModel{
		defaultSocialCircleModel: newSocialCircleModel(conn, c, opts...),
	}
}
