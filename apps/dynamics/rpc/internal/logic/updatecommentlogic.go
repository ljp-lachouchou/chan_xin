package logic

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/pkg/lerr"

	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/dynamics"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCommentLogic {
	return &UpdateCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新评论
func (l *UpdateCommentLogic) UpdateComment(in *dynamics.UpdateCommentReq) (*dynamics.Empty, error) {
	findOne, err := l.svcCtx.CommentsModel.FindOne(l.ctx, in.CommentId)
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "dynamics-rpc UpdateComment", in.CommentId)
	}
	findOne.IsDeleted = in.IsDeleted
	err = l.svcCtx.CommentsModel.Update(l.ctx, findOne)
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "dynamics-rpc UpdateComment", findOne)
	}
	return &dynamics.Empty{}, nil
}
