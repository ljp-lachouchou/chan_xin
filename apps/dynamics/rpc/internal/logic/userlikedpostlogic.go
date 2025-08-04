package logic

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/dynamicsmodels"
	"github.com/ljp-lachouchou/chan_xin/pkg/lerr"

	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/dynamics"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLikedPostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserLikedPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLikedPostLogic {
	return &UserLikedPostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserLikedPostLogic) UserLikedPost(in *dynamics.UserLikedPostReq) (*dynamics.UserLikedPostResp, error) {
	isDeleted, err := l.svcCtx.PostLikesModel.UserInPost(l.ctx, in.UserId, in.PostId)
	if err != nil {
		if err == dynamicsmodels.ErrNotFound {
			return &dynamics.UserLikedPostResp{
				IsLiked: false,
			}, nil
		}
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "PostLikesModel.UserInPost is err ", in.UserId, in.PostId)
	}
	return &dynamics.UserLikedPostResp{
		IsLiked: !(*isDeleted),
	}, nil
}
