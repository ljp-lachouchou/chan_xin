package logic

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/pkg/lerr"

	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/dynamics"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateNotificationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateNotificationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateNotificationLogic {
	return &UpdateNotificationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新通知
func (l *UpdateNotificationLogic) UpdateNotification(in *dynamics.UpdateNotificationReq) (*dynamics.Empty, error) {
	err := l.svcCtx.NotificationsModel.UpdateByUserId(l.ctx, in.UserId)
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "dynamics-rpc UpdateNotification", in.UserId)
	}
	return &dynamics.Empty{}, nil
}
