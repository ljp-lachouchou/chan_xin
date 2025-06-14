package user

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/user/api/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/user/rpc/userservice"
	"github.com/ljp-lachouchou/chan_xin/pkg/lerr"
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
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.UserService.Ping(l.ctx, &userservice.PingReq{})
	if err != nil {
		return lerr.NewWrapError(lerr.NewCOMMONError(), err, "PingRpc failed")
	}
	return nil
}
