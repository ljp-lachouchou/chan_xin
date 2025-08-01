package logic

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/dynamics"

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

func (l *ListVisiblePostsLogic) ListVisiblePosts(req *types.ListVisiblePostsRequest) (*types.PostListResponse, error) {
	var token = req.PageToken
	if token == "NONE" {
		token = ""
	}
	posts, err := l.svcCtx.ListVisiblePosts(l.ctx, &dynamics.ListVisiblePostsRequest{
		ViewerId: req.ViewerId,
		Pagination: &dynamics.Pagination{
			PageSize:  int32(req.PageSize),
			PageToken: token,
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
			IsPinned:   v.IsPinned,
			CreateTime: v.CreateTime,
		})
	}
	return &types.PostListResponse{
		Posts:         postList,
		NextPageToken: posts.NextPageToken,
	}, nil
}
