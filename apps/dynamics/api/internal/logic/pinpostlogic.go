package logic

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/dynamics"

	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/api/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PinPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 置顶/取消置顶
func NewPinPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PinPostLogic {
	return &PinPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PinPostLogic) PinPost(req *types.PinPostRequest) (*types.Empty, error) {
	_, err := l.svcCtx.Dynamics.PinPost(l.ctx, &dynamics.PinPostRequest{
		UserId: req.UserId,
		PostId: req.PostId,
		Pin:    req.Pin,
	})
	if err != nil {
		return nil, err
	}
	return &types.Empty{}, nil
}
