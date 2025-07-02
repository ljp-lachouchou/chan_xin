package dynamicsmodels

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ NotificationsModel = (*customNotificationsModel)(nil)

type (
	// NotificationsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customNotificationsModel.
	NotificationsModel interface {
		notificationsModel
	}

	customNotificationsModel struct {
		*defaultNotificationsModel
	}
)

// NewNotificationsModel returns a model for the database table.
func NewNotificationsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) NotificationsModel {
	return &customNotificationsModel{
		defaultNotificationsModel: newNotificationsModel(conn, c, opts...),
	}
}
