package logic

import (
	"context"

	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/api/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListVisiblePostsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 浏览动态流
func NewListVisiblePostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListVisiblePostsLogic {
	return &ListVisiblePostsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListVisiblePostsLogic) ListVisiblePosts(req *types.ListVisiblePostsRequest) (resp *types.PostListResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
