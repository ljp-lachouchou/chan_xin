package logic

import (
	"context"

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

func (l *UpdateNotificationLogic) UpdateNotification(req *types.UpdateNotificationReq) (resp *types.Empty, err error) {
	// todo: add your logic here and delete this line

	return
}
