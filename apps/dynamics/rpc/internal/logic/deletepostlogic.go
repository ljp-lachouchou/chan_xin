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
	PostNotCanDelete = errors.New("出现错误，此动态无法被删除，请稍后再试")
)

type DeletePostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePostLogic {
	return &DeletePostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除动态（仅创建者可操作）
func (l *DeletePostLogic) DeletePost(in *dynamics.DeletePostRequest) (*dynamics.Empty, error) {
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
	if findOne.UserId != in.UserId {
		return nil, errors.WithStack(UserIdNotMatch)
	}
	success, err := l.svcCtx.PostsModel.Delete(l.ctx, findOne.ID.Hex())
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "dynamics-rpc DeletePost PostsModel.Delete()", in.PostId)
	}
	if success == 0 {
		return nil, PostNotCanDelete
	}
	return &dynamics.Empty{}, nil
}
