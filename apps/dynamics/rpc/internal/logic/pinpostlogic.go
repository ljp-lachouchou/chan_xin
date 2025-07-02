package logic

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/dynamicsmodels"
	"github.com/ljp-lachouchou/chan_xin/pkg/lerr"
	"github.com/pkg/errors"

	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/dynamics"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	PostsNoFind    = errors.New("没有这条动态")
	UserIdNotMatch = errors.New("用户ID不匹配")
)

type PinPostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPinPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PinPostLogic {
	return &PinPostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 置顶/取消置顶动态
func (l *PinPostLogic) PinPost(in *dynamics.PinPostRequest) (*dynamics.Empty, error) {
	findOne, err := l.svcCtx.PostsModel.FindOne(l.ctx, in.PostId)
	if err != nil {
		if err == dynamicsmodels.MongoErrNotFound {
			return nil, errors.WithStack(PostsNoFind)
		}
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "dynamics-rpc PinPost PostsModel.FindOne", in.PostId)
	}
	if findOne.UserId != in.UserId {
		return nil, errors.WithStack(UserIdNotMatch)
	}
	findOne.IsPinned = in.Pin
	_, err = l.svcCtx.PostsModel.Update(l.ctx, findOne)
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "dynamics-rpc PinPost PostsModel.Update", findOne)
	}
	return &dynamics.Empty{}, nil
}
