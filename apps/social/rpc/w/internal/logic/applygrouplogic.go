package logic

import (
	"context"

	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/social"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/w/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApplyGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewApplyGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApplyGroupLogic {
	return &ApplyGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ApplyGroupLogic) ApplyGroup(in *social.GroupApplyReq) (*social.GroupApplyResp, error) {
	// todo: add your logic here and delete this line

	return &social.GroupApplyResp{}, nil
}
