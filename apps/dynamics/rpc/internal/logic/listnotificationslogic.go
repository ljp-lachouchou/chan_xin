package logic

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/dynamicsmodels"
	"github.com/ljp-lachouchou/chan_xin/pkg/ctxdata"
	"github.com/ljp-lachouchou/chan_xin/pkg/lerr"
	"github.com/pkg/errors"
	"time"

	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/dynamics"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListNotificationsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListNotificationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListNotificationsLogic {
	return &ListNotificationsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取通知列表（分页）
func (l *ListNotificationsLogic) ListNotifications(in *dynamics.ListNotificationsRequest) (*dynamics.ListNotificationsResponse, error) {
	if in.Pagination == nil {
		return nil, errors.WithStack(PaginationNilErr)
	}
	var startLocation int32 = 0
	if in.Pagination.PageToken == "" {
		return l.listNotifications(in.Pagination.PageSize, startLocation, in.UserId)
	}
	secretKey := []byte(l.svcCtx.Config.JwtAuth.AccessSecret)
	parse, err := ctxdata.ParseByTokenString(secretKey, in.Pagination.PageToken)
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NewCOMMONError(), err, "dynamic-rpc ListNotifications ctxdata.ParseByTokenString")
	}
	v, ok := parse.Claims.(jwt.MapClaims)
	if !ok {
		return nil, lerr.NewWrapError(lerr.NewCOMMONError(), err, "dynamic-rpc ListNotifications claims error")
	}
	startLocation = int32(v[ctxdata.PageIdentify].(float64))

	return l.listNotifications(in.Pagination.PageSize, startLocation, in.UserId)
}
func (l *ListNotificationsLogic) listNotifications(pageSize, startLocation int32, userId string) (*dynamics.ListNotificationsResponse, error) {
	endLocation := pageSize + startLocation
	token, err := ctxdata.GenPageToken(l.svcCtx.Config.JwtAuth.AccessSecret, time.Now().Unix(), l.svcCtx.Config.JwtAuth.AccessExpire, endLocation)
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NewCOMMONError(), err, "dynamics-rpc ListNotifications ctxdata.GenPageToken")
	}
	notifications, err := l.svcCtx.NotificationsModel.FindByUserIdWithPage(l.ctx, userId, pageSize, startLocation)
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "dynamics-rpc listNotifications FindByUserId")
	}
	notificationsDynamics := copyNotifications(notifications)
	return &dynamics.ListNotificationsResponse{
		Notifications: notificationsDynamics,
		NextPageToken: token,
	}, nil
}
func copyNotifications(notifications []*dynamicsmodels.Notifications) []*dynamics.Notification {
	var notificationsDynamics []*dynamics.Notification
	for _, v := range notifications {
		notification := &dynamics.Notification{
			Id:            v.Id,
			Type:          dynamics.NotificationType(dynamics.NotificationType_value[v.Type]),
			TriggerUserId: v.TriggerUserId,
			PostId:        v.PostId,
			CommentId:     v.CommentId.String,
			IsRead:        v.IsRead,
		}
		notificationsDynamics = append(notificationsDynamics, notification)
	}
	return notificationsDynamics
}
