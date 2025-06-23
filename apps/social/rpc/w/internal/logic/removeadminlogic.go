package logic

import (
	"context"

	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/social"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/w/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveAdminLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveAdminLogic {
	return &RemoveAdminLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveAdminLogic) RemoveAdmin(in *social.RemoveAdminReq) (*social.RemoveAdminResp, error) {
	// todo: add your logic here and delete this line

	return &social.RemoveAdminResp{}, nil
}
