package logic

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/api/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/api/internal/types"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/dynamicsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCommentReplayLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建评论回复
func NewCreateCommentReplayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCommentReplayLogic {
	return &CreateCommentReplayLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCommentReplayLogic) CreateCommentReplay(req *types.CreateCommentReplayReq) (*types.CreateCommentReplayResp, error) {
	resp, err := l.svcCtx.Dynamics.CreateCommentReplay(l.ctx, &dynamicsclient.CreateCommentReplayReq{
		CommentId: req.CommentId,
		UserId:    req.UserId,
		Content:   req.Content,
	})
	if err != nil {
		return nil, err
	}
	postInfo, err2 := l.svcCtx.Dynamics.GetPostInfo(l.ctx, &dynamicsclient.GetPostInfoReq{
		PostId: resp.PostId,
	})
	if err2 != nil {
		return nil, err2
	}
	_, err = l.svcCtx.Dynamics.CreateNotification(l.ctx, &dynamicsclient.CreateNotificationReq{
		UserId:        postInfo.UserId,
		Type:          1,
		TriggerUserId: req.UserId,
		PostId:        resp.PostId,
		CommentId:     resp.CommentReplyId,
	})
	if err != nil {
		return nil, err
	}
	return &types.CreateCommentReplayResp{
		CommentReplyId: resp.CommentReplyId,
		PostId:         resp.PostId,
	}, nil
}
