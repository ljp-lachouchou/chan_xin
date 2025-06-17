package socialmodels

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ GroupApplyModel = (*customGroupApplyModel)(nil)

type (
	// GroupApplyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGroupApplyModel.
	GroupApplyModel interface {
		groupApplyModel
	}

	customGroupApplyModel struct {
		*defaultGroupApplyModel
	}
)

// NewGroupApplyModel returns a model for the database table.
func NewGroupApplyModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) GroupApplyModel {
	return &customGroupApplyModel{
		defaultGroupApplyModel: newGroupApplyModel(conn, c, opts...),
	}
}
