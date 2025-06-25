package logic

import (
	"context"

	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/dynamics"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListUserPostsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListUserPostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListUserPostsLogic {
	return &ListUserPostsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户所有动态（按置顶状态+时间倒序）
func (l *ListUserPostsLogic) ListUserPosts(in *dynamics.ListUserPostsRequest) (*dynamics.PostListResponse, error) {
	// todo: add your logic here and delete this line

	return &dynamics.PostListResponse{}, nil
}
