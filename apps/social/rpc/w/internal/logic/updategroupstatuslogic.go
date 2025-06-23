package logic

import (
	"context"

	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/social"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/w/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateGroupStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateGroupStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGroupStatusLogic {
	return &UpdateGroupStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateGroupStatusLogic) UpdateGroupStatus(in *social.GroupStatusUpdate) (*social.GroupStatusUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &social.GroupStatusUpdateResp{}, nil
}
