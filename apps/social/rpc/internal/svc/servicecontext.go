package svc

import (
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/internal/config"
	"github.com/ljp-lachouchou/chan_xin/apps/social/socialmodels"
	"github.com/ljp-lachouchou/chan_xin/apps/user/usermodels"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	usermodels.UsersModel
	socialmodels.FriendRelationModel
	socialmodels.FriendApplyModel
	socialmodels.GroupInfoModel
	socialmodels.GroupMemberModel
	socialmodels.GroupOperationModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:              c,
		UsersModel:          usermodels.NewUsersModel(conn, c.Cache),
		FriendRelationModel: socialmodels.NewFriendRelationModel(conn, c.Cache),
		FriendApplyModel:    socialmodels.NewFriendApplyModel(conn, c.Cache),
		GroupInfoModel:      socialmodels.NewGroupInfoModel(conn, c.Cache),
		GroupMemberModel:    socialmodels.NewGroupMemberModel(conn, c.Cache),
		GroupOperationModel: socialmodels.NewGroupOperationModel(conn, c.Cache),
	}
}
