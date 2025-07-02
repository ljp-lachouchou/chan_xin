package group

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/socialservice"

	"github.com/ljp-lachouchou/chan_xin/apps/social/api/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/social/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveAdminLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 移除管理员
func NewRemoveAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveAdminLogic {
	return &RemoveAdminLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveAdminLogic) RemoveAdmin(req *types.RemoveAdminReq) error {
	in := &socialservice.RemoveAdminReq{
		OperatorId: req.OperatorId,
		GroupId:    req.GroupId,
		TargetId:   req.TargetId,
	}
	_, err := l.svcCtx.SocialService.RemoveAdmin(l.ctx, in)
	return err
}
