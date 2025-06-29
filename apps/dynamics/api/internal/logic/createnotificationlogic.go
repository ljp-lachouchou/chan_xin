package logic

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/dynamics"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/dynamicsclient"

	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/api/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateNotificationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建通知
func NewCreateNotificationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateNotificationLogic {
	return &CreateNotificationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateNotificationLogic) CreateNotification(req *types.CreateNotificationReq) (*types.Empty, error) {
	_, err := l.svcCtx.CreateNotification(l.ctx, &dynamicsclient.CreateNotificationReq{
		UserId:        req.UserId,
		Type:          dynamics.NotificationType(req.Type),
		TriggerUserId: req.TriggerUserId,
		PostId:        req.PostId,
		CommentId:     req.CommentId,
	})
	if err != nil {
		return nil, err
	}
	return &types.Empty{}, nil
}
