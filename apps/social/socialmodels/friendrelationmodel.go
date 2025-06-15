package socialmodels

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FriendRelationModel = (*customFriendRelationModel)(nil)

type (
	// FriendRelationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFriendRelationModel.
	FriendRelationModel interface {
		friendRelationModel
	}

	customFriendRelationModel struct {
		*defaultFriendRelationModel
	}
)

// NewFriendRelationModel returns a model for the database table.
func NewFriendRelationModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) FriendRelationModel {
	return &customFriendRelationModel{
		defaultFriendRelationModel: newFriendRelationModel(conn, c, opts...),
	}
}
