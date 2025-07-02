package logic

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/dynamics"

	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/api/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLikeByPostIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 点赞列表
func NewListLikeByPostIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLikeByPostIdLogic {
	return &ListLikeByPostIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLikeByPostIdLogic) ListLikeByPostId(req *types.GetPostInfoReq) (*types.Ids, error) {
	id, err := l.svcCtx.Dynamics.ListLikeByPostId(l.ctx, &dynamics.GetPostInfoReq{
		PostId: req.PostId,
	})
	if err != nil {
		return nil, err
	}
	return &types.Ids{
		Ids: id.Ids,
	}, nil
}
