package friend

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/socialservice"

	"github.com/ljp-lachouchou/chan_xin/apps/social/api/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/social/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateFriendStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新我对此好友的状态
func NewUpdateFriendStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFriendStatusLogic {
	return &UpdateFriendStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateFriendStatusLogic) UpdateFriendStatus(req *types.FriendStatusUpdate) error {
	in := &socialservice.FriendStatusUpdate{
		UserId:   req.UserId,
		FriendId: req.FriendId,
		Status: &socialservice.FriendStatus{
			IsMuted:   req.Status.IsMuted,
			IsTopped:  req.Status.IsTopped,
			IsBlocked: req.Status.IsBlocked,
			Remark:    req.Status.Remark,
		},
	}
	_, err := l.svcCtx.SocialService.UpdateFriendStatus(l.ctx, in)

	return err
}
