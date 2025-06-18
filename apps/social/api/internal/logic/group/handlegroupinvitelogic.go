package group

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/socialservice"

	"github.com/ljp-lachouchou/chan_xin/apps/social/api/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/social/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HandleGroupInviteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 被邀请者处理群申请
func NewHandleGroupInviteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HandleGroupInviteLogic {
	return &HandleGroupInviteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HandleGroupInviteLogic) HandleGroupInvite(req *types.GroupInviteAction) error {
	in := &socialservice.GroupInviteAction{
		InviteId:   req.InviteId,
		IsAccepted: req.IsAccepted,
	}
	_, err := l.svcCtx.SocialService.HandleGroupInvite(l.ctx, in)

	return err
}
