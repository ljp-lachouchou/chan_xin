package logic

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/dynamics"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCoverLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCoverLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCoverLogic {
	return &GetCoverLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取个人动态封面
func (l *GetCoverLogic) GetCover(in *dynamics.GetCoverRequest) (*dynamics.GetCoverResp, error) {
	socialCircle, err := l.svcCtx.SocialCircleModel.FindByUserId(l.ctx, in.UserId)

	if err != nil {
		return nil, err
	}
	return &dynamics.GetCoverResp{
		Cover: socialCircle.CoverUrl.String,
	}, nil
}
