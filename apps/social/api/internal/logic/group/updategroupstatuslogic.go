package group

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/socialservice"

	"github.com/ljp-lachouchou/chan_xin/apps/social/api/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/social/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateGroupStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新某人对此群的状态
func NewUpdateGroupStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGroupStatusLogic {
	return &UpdateGroupStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateGroupStatusLogic) UpdateGroupStatus(req *types.GroupStatusUpdate) error {
	in := &socialservice.GroupStatusUpdate{
		UserId:  req.UserId,
		GroupId: req.GroupId,
		Status: &socialservice.GroupStatus{
			IsMuted:  &req.Status.IsMuted,
			IsTopped: &req.Status.IsTopped,
			Remark:   &req.Status.Remark,
		},
	}
	_, err := l.svcCtx.SocialService.UpdateGroupStatus(l.ctx, in)

	return err
}
