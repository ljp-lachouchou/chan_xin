package logic

import (
	"context"

	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/dynamics"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetCoverLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetCoverLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetCoverLogic {
	return &SetCoverLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 设置个人动态封面（用于个人主页）
func (l *SetCoverLogic) SetCover(in *dynamics.SetCoverRequest) (*dynamics.Empty, error) {
	// todo: add your logic here and delete this line

	return &dynamics.Empty{}, nil
}
