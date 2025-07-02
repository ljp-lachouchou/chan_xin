package logic

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/dynamics"

	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/api/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetCoverLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 设置个人封面
func NewSetCoverLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetCoverLogic {
	return &SetCoverLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetCoverLogic) SetCover(req *types.SetCoverRequest) (*types.Empty, error) {

	_, err := l.svcCtx.Dynamics.SetCover(l.ctx, &dynamics.SetCoverRequest{
		UserId:   req.UserId,
		CoverUrl: req.CoverUrl,
	})
	if err != nil {
		return nil, err
	}
	return &types.Empty{}, nil
}
