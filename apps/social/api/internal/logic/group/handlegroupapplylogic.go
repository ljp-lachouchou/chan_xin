package group

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/socialservice"

	"github.com/ljp-lachouchou/chan_xin/apps/social/api/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/social/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HandleGroupApplyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 群处理申请
func NewHandleGroupApplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HandleGroupApplyLogic {
	return &HandleGroupApplyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HandleGroupApplyLogic) HandleGroupApply(req *types.GroupApplyAction) error {
	in := &socialservice.GroupApplyAction{
		ApplyId:    req.ApplyId,
		ManagerId:  req.ManagerId,
		IsApproved: req.IsApproved,
	}
	_, err := l.svcCtx.SocialService.HandleGroupApply(l.ctx, in)

	return err
}
