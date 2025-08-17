package logic

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/dynamics"

	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/api/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 单个post信息
func NewGetPostInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostInfoLogic {
	return &GetPostInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPostInfoLogic) GetPostInfo(req *types.GetPostInfoReq) (*types.Post, error) {

	info, err := l.svcCtx.Dynamics.GetPostInfo(l.ctx, &dynamics.GetPostInfoReq{
		PostId: req.PostId,
	})
	if err != nil {
		return nil, err
	}
	return &types.Post{
		PostId: info.PostId,
		UserId: info.UserId,
		Content: types.PostContent{
			Text:      info.Content.Text,
			ImageUrls: info.Content.ImageUrls,
			Emoji:     info.Content.Emoji,
		},
		Meta: types.PostMeta{
			Location:       info.Meta.Location,
			Scope:          int(info.Meta.Scope),
			VisibleUserIds: info.Meta.VisibleUserIds,
		},
		IsPinned: info.IsPinned,
		CreateTime: info.CreateTime,
	}, nil
}
