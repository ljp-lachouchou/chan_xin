package logic

import (
	"context"

	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/dynamics"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PinPostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPinPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PinPostLogic {
	return &PinPostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 置顶/取消置顶动态
func (l *PinPostLogic) PinPost(in *dynamics.PinPostRequest) (*dynamics.Empty, error) {

	return &dynamics.Empty{}, nil
}
