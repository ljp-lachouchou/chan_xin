package logic

import (
	"context"

	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/dynamics"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListVisiblePostsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListVisiblePostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListVisiblePostsLogic {
	return &ListVisiblePostsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 浏览可见动态流（根据权限过滤+分页）
func (l *ListVisiblePostsLogic) ListVisiblePosts(in *dynamics.ListVisiblePostsRequest) (*dynamics.PostListResponse, error) {

	return &dynamics.PostListResponse{}, nil
}
