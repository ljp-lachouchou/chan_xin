package logic

import (
	"context"

	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/api/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCommentReplayLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新回复状态
func NewUpdateCommentReplayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCommentReplayLogic {
	return &UpdateCommentReplayLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCommentReplayLogic) UpdateCommentReplay(req *types.UpdateCommentReplayReq) (resp *types.Empty, err error) {
	// todo: add your logic here and delete this line

	return
}
