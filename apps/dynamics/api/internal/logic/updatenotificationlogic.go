package logic

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/dynamics"

	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/api/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateNotificationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新通知状态
func NewUpdateNotificationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateNotificationLogic {
	return &UpdateNotificationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateNotificationLogic) UpdateNotification(req *types.UpdateNotificationReq) (*types.Empty, error) {

	_, err := l.svcCtx.UpdateNotification(l.ctx, &dynamics.UpdateNotificationReq{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, err
	}
	return &types.Empty{}, nil
}
