package socialmodels

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ GroupOperationModel = (*customGroupOperationModel)(nil)

type (
	// GroupOperationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGroupOperationModel.
	GroupOperationModel interface {
		groupOperationModel
	}

	customGroupOperationModel struct {
		*defaultGroupOperationModel
	}
)

// NewGroupOperationModel returns a model for the database table.
func NewGroupOperationModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) GroupOperationModel {
	return &customGroupOperationModel{
		defaultGroupOperationModel: newGroupOperationModel(conn, c, opts...),
	}
}
