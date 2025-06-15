package socialmodels

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FriendApplyModel = (*customFriendApplyModel)(nil)

type (
	// FriendApplyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFriendApplyModel.
	FriendApplyModel interface {
		friendApplyModel
	}

	customFriendApplyModel struct {
		*defaultFriendApplyModel
	}
)

// NewFriendApplyModel returns a model for the database table.
func NewFriendApplyModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) FriendApplyModel {
	return &customFriendApplyModel{
		defaultFriendApplyModel: newFriendApplyModel(conn, c, opts...),
	}
}
