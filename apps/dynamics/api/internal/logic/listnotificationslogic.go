package logic

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/dynamics"

	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/api/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListNotificationsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取通知列表
func NewListNotificationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListNotificationsLogic {
	return &ListNotificationsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListNotificationsLogic) ListNotifications(req *types.ListNotificationsRequest) (*types.ListNotificationsResponse, error) {
	notifications, err := l.svcCtx.Dynamics.ListNotifications(l.ctx, &dynamics.ListNotificationsRequest{
		UserId: req.UserId,
		Pagination: &dynamics.Pagination{
			PageSize:  int32(req.Pagination.PageSize),
			PageToken: req.Pagination.PageToken,
		},
	})
	if err != nil {
		return nil, err
	}
	var notificationList []types.Notification
	for _, notification := range notifications.Notifications {
		notificationList = append(notificationList, types.Notification{
			Id:            notification.Id,
			Type:          int(notification.Type),
			TriggerUserId: notification.TriggerUserId,
			PostId:        notification.PostId,
			CommentId:     notification.CommentId,
			IsRead:        notification.IsRead,
		})
	}
	return &types.ListNotificationsResponse{
		Notifications: notificationList,
		NextPageToken: notifications.NextPageToken,
	}, nil
}
