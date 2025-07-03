package logic

import (
	"context"

	"github.com/ljp-lachouchou/chan_xin/apps/im/rpc/im"
	"github.com/ljp-lachouchou/chan_xin/apps/im/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *im.PingRep) (*im.PingResp, error) {

	return &im.PingResp{}, nil
}
