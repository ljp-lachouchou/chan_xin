package logic

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/pkg/lerr"

	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/dynamics"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCommentReplayLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCommentReplayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCommentReplayLogic {
	return &UpdateCommentReplayLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新评论回复
func (l *UpdateCommentReplayLogic) UpdateCommentReplay(in *dynamics.UpdateCommentReplayReq) (*dynamics.Empty, error) {
	findOne, err := l.svcCtx.CommentRepliesModel.FindOne(l.ctx, in.CommentReplayId)
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "dynamics-rpc UpdateCommentReplay", in.CommentReplayId)
	}
	findOne.IsDeleted = in.IsDeleted
	err = l.svcCtx.CommentRepliesModel.Update(l.ctx, findOne)
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "dynamics-rpc UpdateCommentReplay", findOne)
	}
	return &dynamics.Empty{}, nil
}
