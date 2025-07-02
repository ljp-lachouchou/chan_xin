package logic

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/dynamics"

	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/api/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCommentReplayLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除评论回复
func NewDeleteCommentReplayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommentReplayLogic {
	return &DeleteCommentReplayLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCommentReplayLogic) DeleteCommentReplay(req *types.DeleteCommentReplayReq) (*types.Empty, error) {
	_, err := l.svcCtx.Dynamics.DeleteCommentReplay(l.ctx, &dynamics.DeleteCommentReplayReq{
		CommentReplayId: req.CommentReplayId,
	})
	if err != nil {
		return nil, err
	}
	return &types.Empty{}, nil
}
