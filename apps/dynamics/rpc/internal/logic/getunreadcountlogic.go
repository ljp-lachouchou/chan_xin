package logic

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/pkg/lerr"

	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/dynamics"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUnreadCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUnreadCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUnreadCountLogic {
	return &GetUnreadCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 新增：获取未读通知数量
func (l *GetUnreadCountLogic) GetUnreadCount(in *dynamics.GetUnreadCountRequest) (*dynamics.GetUnreadCountResponse, error) {
	notRead, err := l.svcCtx.NotificationsModel.FindByUserIdAndIsNotRead(l.ctx, in.UserId)
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "dynamics-rpc GetUnreadCount ", in.UserId)
	}
	return &dynamics.GetUnreadCountResponse{
		UnreadCount: notRead,
	}, nil
}
