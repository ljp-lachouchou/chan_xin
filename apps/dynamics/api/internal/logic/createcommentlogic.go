package logic

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/dynamicsclient"

	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/api/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建评论
func NewCreateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCommentLogic {
	return &CreateCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCommentLogic) CreateComment(req *types.CreateCommentReq) (*types.Empty, error) {

	_, err := l.svcCtx.Dynamics.CreateComment(l.ctx, &dynamicsclient.CreateCommentReq{
		PostId:  req.PostId,
		UserId:  req.UserId,
		Content: req.Content,
	})
	l.svcCtx.Dynamics.CreateNotification(l.ctx, &dynamicsclient.CreateNotificationReq{
		UserId:        "",
		Type:          1,
		TriggerUserId: req.UserId,
		PostId:        req.PostId,
	})
	if err != nil {
		return nil, err
	}
	return &types.Empty{}, nil
}
