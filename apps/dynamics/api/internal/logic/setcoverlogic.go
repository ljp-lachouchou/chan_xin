package logic

import (
	"context"

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

func (l *SetCoverLogic) SetCover(req *types.SetCoverRequest) (resp *types.Empty, err error) {
	// todo: add your logic here and delete this line

	return
}
