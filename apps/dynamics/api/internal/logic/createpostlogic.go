package logic

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/dynamics"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/dynamicsclient"

	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/api/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建动态
func NewCreatePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePostLogic {
	return &CreatePostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreatePostLogic) CreatePost(req *types.CreatePostRequest) (*types.Post, error) {
	post, err := l.svcCtx.CreatePost(l.ctx, &dynamicsclient.CreatePostRequest{
		UserId: req.UserId,
		Content: &dynamics.PostContent{
			Text:      req.Content.Text,
			ImageUrls: req.Content.ImageUrls,
			Emoji:     req.Content.Emoji,
		},
		Meta: &dynamics.PostMeta{
			Location:       req.Meta.Location,
			Scope:          dynamics.VisibleScope(req.Meta.Scope),
			VisibleUserIds: req.Meta.VisibleUserIds,
		},
	})
	if err != nil {
		return nil, err
	}
	return &types.Post{
		PostId: post.PostId,
		UserId: post.UserId,
		Content: types.PostContent{
			Text:      post.Content.Text,
			ImageUrls: post.Content.ImageUrls,
			Emoji:     post.Content.Emoji,
		},
		Meta: types.PostMeta{
			Location:       post.Meta.Location,
			Scope:          int(post.Meta.Scope),
			VisibleUserIds: post.Meta.VisibleUserIds,
		},
		IsPinned: post.IsPinned,
	}, nil
}
