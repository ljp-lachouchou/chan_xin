package socialmodels

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ GroupMemberModel = (*customGroupMemberModel)(nil)

type (
	// GroupMemberModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGroupMemberModel.
	GroupMemberModel interface {
		groupMemberModel
	}

	customGroupMemberModel struct {
		*defaultGroupMemberModel
	}
)

// NewGroupMemberModel returns a model for the database table.
func NewGroupMemberModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) GroupMemberModel {
	return &customGroupMemberModel{
		defaultGroupMemberModel: newGroupMemberModel(conn, c, opts...),
	}
}
