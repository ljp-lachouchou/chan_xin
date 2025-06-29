package logic

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/pkg/lerr"

	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/dynamics"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListNotificationsByUserIdAndTypeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListNotificationsByUserIdAndTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListNotificationsByUserIdAndTypeLogic {
	return &ListNotificationsByUserIdAndTypeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据type和userid查找
func (l *ListNotificationsByUserIdAndTypeLogic) ListNotificationsByUserIdAndType(in *dynamics.ListNotificationsByUserIdAndTypeReq) (*dynamics.ListNotificationsByUserIdAndTypeReqResponse, error) {
	notifications, err := l.svcCtx.NotificationsModel.FindByUserIdAndType(l.ctx, in.UserId, in.Type.String())
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "dynamics-rpc ListNotificationsByUserIdAndType FindByUserIdAndType", in.UserId, in.Type.String())
	}
	notificationDynamics := copyNotifications(notifications)
	return &dynamics.ListNotificationsByUserIdAndTypeReqResponse{
		Notifications: notificationDynamics,
	}, nil
}
