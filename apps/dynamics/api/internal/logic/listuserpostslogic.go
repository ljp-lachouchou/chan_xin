package logic

import (
	"context"

	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/api/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListUserPostsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户动态列表
func NewListUserPostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListUserPostsLogic {
	return &ListUserPostsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListUserPostsLogic) ListUserPosts(req *types.ListUserPostsRequest) (resp *types.PostListResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
