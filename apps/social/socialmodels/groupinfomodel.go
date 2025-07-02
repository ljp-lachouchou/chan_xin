package socialmodels

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ GroupInfoModel = (*customGroupInfoModel)(nil)

type (
	// GroupInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGroupInfoModel.
	GroupInfoModel interface {
		groupInfoModel
	}

	customGroupInfoModel struct {
		*defaultGroupInfoModel
	}
)

// NewGroupInfoModel returns a model for the database table.
func NewGroupInfoModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) GroupInfoModel {
	return &customGroupInfoModel{
		defaultGroupInfoModel: newGroupInfoModel(conn, c, opts...),
	}
}
