package group

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/social"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/socialservice"

	"github.com/ljp-lachouchou/chan_xin/apps/social/api/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/social/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ManageGroupMemberLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 设置管理员
func NewManageGroupMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ManageGroupMemberLogic {
	return &ManageGroupMemberLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ManageGroupMemberLogic) ManageGroupMember(req *types.GroupMemberManage) error {
	in := &socialservice.GroupMemberManage{
		OperatorId: req.OperatorId,
		GroupId:    req.GroupId,
		TargetId:   req.TargetId,
		Action:     social.GroupPermission(req.Action.Action),
	}
	_, err := l.svcCtx.SocialService.ManageGroupMember(l.ctx, in)
	return err
}
