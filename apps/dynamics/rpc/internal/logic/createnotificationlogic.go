package logic

import (
	"context"
	"database/sql"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/dynamicsmodels"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/dynamics"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/pkg/lerr"
	"github.com/ljp-lachouchou/chan_xin/pkg/wuid"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateNotificationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateNotificationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateNotificationLogic {
	return &CreateNotificationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建通知
func (l *CreateNotificationLogic) CreateNotification(in *dynamics.CreateNotificationReq) (*dynamics.Empty, error) {
	_, err := l.svcCtx.NotificationsModel.Insert(l.ctx, &dynamicsmodels.Notifications{
		Id:            wuid.GenUid(l.svcCtx.Config.Mysql.DataSource),
		UserId:        in.UserId,
		Type:          in.Type.String(),
		TriggerUserId: in.TriggerUserId,
		PostId:        in.PostId,
		CommentId: sql.NullString{
			Valid:  true,
			String: in.CommentId,
		},
		IsRead: false,
	})
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "dynamics-rpc CreateNotification")
	}
	return &dynamics.Empty{}, nil
}
