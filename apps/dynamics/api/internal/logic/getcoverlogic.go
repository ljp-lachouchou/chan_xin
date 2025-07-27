package logic

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/dynamics"

	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/api/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCoverLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取封面
func NewGetCoverLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCoverLogic {
	return &GetCoverLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCoverLogic) GetCover(req *types.GetCoverRequest) (*types.GetCoverResp, error) {
	coverResp, err := l.svcCtx.GetCover(l.ctx, &dynamics.GetCoverRequest{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, err
	}

	return &types.GetCoverResp{
		Cover: coverResp.Cover,
	}, nil
}
