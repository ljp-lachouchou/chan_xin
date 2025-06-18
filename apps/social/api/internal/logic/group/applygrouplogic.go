package group

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/socialservice"

	"github.com/ljp-lachouchou/chan_xin/apps/social/api/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/social/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApplyGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 某人申请入群
func NewApplyGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApplyGroupLogic {
	return &ApplyGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApplyGroupLogic) ApplyGroup(req *types.GroupApplyReq) (resp *types.GroupApplyResp, err error) {
	in := &socialservice.GroupApplyReq{
		ApplicantId: req.ApplicantId,
		TargetId:    req.TargetId,
		GreetMsg:    req.GreetMsg,
	}
	groupApplyResp, err := l.svcCtx.SocialService.ApplyGroup(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return &types.GroupApplyResp{
		ApplyId:   groupApplyResp.ApplyId,
		ApplyTime: groupApplyResp.ApplyTime,
	}, nil
}
