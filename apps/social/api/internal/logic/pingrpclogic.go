package logic

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/socialservice"

	"github.com/ljp-lachouchou/chan_xin/apps/social/api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type PingRpcLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 保持与etcd的连接
func NewPingRpcLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingRpcLogic {
	return &PingRpcLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PingRpcLogic) PingRpc() error {
	_, err := l.svcCtx.SocialService.Ping(l.ctx, &socialservice.PingReq{})
	return err
}
