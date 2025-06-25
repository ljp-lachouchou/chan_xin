package logic

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/dynamicsmodels"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/dynamics"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/pkg/lerr"
	"github.com/ljp-lachouchou/chan_xin/pkg/wuid"
	"github.com/zeromicro/go-zero/core/logx"
)

type ToggleLikeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewToggleLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ToggleLikeLogic {
	return &ToggleLikeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 点赞/取消点赞
func (l *ToggleLikeLogic) ToggleLike(in *dynamics.LikeAction) (*dynamics.Empty, error) {
	findOne, err := l.svcCtx.PostLikesModel.FindOneByPostIdUserId(l.ctx, in.PostId, in.LikerId)
	if err != nil {
		if err == dynamicsmodels.ErrNotFound {
			_, err := l.svcCtx.PostLikesModel.Insert(l.ctx, &dynamicsmodels.PostLikes{
				Id:        wuid.GenUid(l.svcCtx.Config.Mysql.DataSource),
				PostId:    in.PostId,
				UserId:    in.LikerId,
				IsDeleted: in.IsCancel,
			})
			if err != nil {
				return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "dynamics-rpc PinPost PostsModel.Insert")
			}
			return &dynamics.Empty{}, nil
		}
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "dynamics-rpc PinPost PostsModel.FindOne ", in.PostId)
	}
	findOne.IsDeleted = in.IsCancel
	err = l.svcCtx.PostLikesModel.Update(l.ctx, findOne)
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "dynamics-rpc PinPost PostsModel.Update ", findOne)
	}
	return &dynamics.Empty{}, nil
}
