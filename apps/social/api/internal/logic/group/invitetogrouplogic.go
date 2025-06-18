package group

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/socialservice"

	"github.com/ljp-lachouchou/chan_xin/apps/social/api/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/social/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InviteToGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 邀请某人入群
func NewInviteToGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InviteToGroupLogic {
	return &InviteToGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InviteToGroupLogic) InviteToGroup(req *types.GroupInvitation) error {
	in := &socialservice.GroupInvitation{
		InviterId: req.InviterId,
		GroupId:   req.GroupId,
		TargetIds: req.TargetIds,
	}
	_, err := l.svcCtx.SocialService.InviteToGroup(l.ctx, in)
	return err
}
