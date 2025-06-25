package logic

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/dynamicsmodels"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/dynamics"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/pkg/lerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePostLogic {
	return &CreatePostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建动态（需提供内容和隐私设置）
func (l *CreatePostLogic) CreatePost(in *dynamics.CreatePostRequest) (*dynamics.Post, error) {
	data := dynamicsmodels.Posts{
		UserId:       in.UserId,
		Content:      in.Content.Text,
		ImageUrls:    in.Content.ImageUrls,
		Emoji:        in.Content.Emoji,
		Visibility:   int(in.Meta.Scope),
		VisibleTo:    in.Meta.VisibleUserIds,
		LikeCount:    0,
		CommentCount: 0,
		Location:     in.Meta.Location,
		IsPinned:     false,
	}
	err := l.svcCtx.PostsModel.Insert(l.ctx, &data)
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "dynamics-rpc CreatePost PostsModel.Insert", data)
	}

	return &dynamics.Post{
		PostId: data.ID.Hex(),
		UserId: data.UserId,
		Content: &dynamics.PostContent{
			Text:      data.Content,
			ImageUrls: data.ImageUrls,
			Emoji:     data.Emoji,
		},
		Meta: &dynamics.PostMeta{
			Location:       data.Location,
			Scope:          dynamics.VisibleScope(data.Visibility),
			VisibleUserIds: data.VisibleTo,
		},
		IsPinned: false,
	}, nil
}
