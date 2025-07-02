package logic

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/dynamics"

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

func (l *ListUserPostsLogic) ListUserPosts(req *types.ListUserPostsRequest) (*types.PostListResponse, error) {

	posts, err := l.svcCtx.Dynamics.ListUserPosts(l.ctx, &dynamics.ListUserPostsRequest{
		UserId: req.UserId,
		IsPin:  &req.IsPin,
		Pagination: &dynamics.Pagination{
			PageSize:  int32(req.Pagination.PageSize),
			PageToken: req.Pagination.PageToken,
		},
	})
	if err != nil {
		return nil, err
	}
	var postList []types.Post
	for _, v := range posts.Posts {
		postList = append(postList, types.Post{
			PostId: v.PostId,
			UserId: v.UserId,
			Content: types.PostContent{
				Text:      v.Content.Text,
				ImageUrls: v.Content.ImageUrls,
				Emoji:     v.Content.Emoji,
			},
			Meta: types.PostMeta{
				Location:       v.Meta.Location,
				Scope:          int(v.Meta.Scope),
				VisibleUserIds: v.Meta.VisibleUserIds,
			},
			IsPinned: v.IsPinned,
		})
	}
	return &types.PostListResponse{
		Posts:         postList,
		NextPageToken: posts.NextPageToken,
	}, nil
}
