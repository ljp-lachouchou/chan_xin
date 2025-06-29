package logic

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/pkg/lerr"

	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/dynamics"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPostInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostInfoLogic {
	return &GetPostInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 单个post信息
func (l *GetPostInfoLogic) GetPostInfo(in *dynamics.GetPostInfoReq) (*dynamics.Post, error) {
	findOne, err := l.svcCtx.PostsModel.FindOne(l.ctx, in.PostId)
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "GetPostInfo FindOne ", in.PostId)
	}
	return &dynamics.Post{
		PostId: findOne.ID.Hex(),
		UserId: findOne.UserId,
		Content: &dynamics.PostContent{
			Text:      findOne.Content,
			ImageUrls: findOne.ImageUrls,
			Emoji:     findOne.Emoji,
		},
		Meta: &dynamics.PostMeta{
			Location:       findOne.Location,
			Scope:          dynamics.VisibleScope(findOne.Visibility),
			VisibleUserIds: findOne.VisibleTo,
		},
		IsPinned: findOne.IsPinned,
	}, nil
}
