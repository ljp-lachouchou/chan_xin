package logic

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/dynamics"

	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/api/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLikedPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户是否点在动态
func NewUserLikedPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLikedPostLogic {
	return &UserLikedPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLikedPostLogic) UserLikedPost(req *types.UserLikedPostReq) (*types.UserLiekdPostResp,error) {
	resp, err := l.svcCtx.Dynamics.UserLikedPost(l.ctx, &dynamics.UserLikedPostReq{
		UserId: req.UserId,
		PostId: req.PostId,
	})
	if err != nil {
		return nil, err
	}

	return &types.UserLiekdPostResp{
		IsLiked: resp.IsLiked,
	},nil
}
