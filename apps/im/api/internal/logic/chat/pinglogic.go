package chat

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/im/rpc/im"

	"github.com/ljp-lachouchou/chan_xin/apps/im/api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PingLogic) Ping() error {
	_, err := l.svcCtx.Ping(l.ctx, &im.PingRep{})
	if err != nil {
		return err
	}
	return nil
}
