package logic

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/dynamicsmodels"
	"github.com/ljp-lachouchou/chan_xin/pkg/lerr"
	"github.com/ljp-lachouchou/chan_xin/pkg/wuid"

	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/dynamics"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCommentReplayLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCommentReplayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCommentReplayLogic {
	return &CreateCommentReplayLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建评论回复
func (l *CreateCommentReplayLogic) CreateCommentReplay(in *dynamics.CreateCommentReplayReq) (*dynamics.CreateCommentReplayResp, error) {
	findOne, err2 := l.svcCtx.CommentsModel.FindOne(l.ctx, in.CommentId)
	if err2 != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err2, "dynamics-rpc CreateCommentReplay CommentsModel.FindOne ", in.CommentId)
	}
	commentReplyId := wuid.GenUid(l.svcCtx.Config.Mysql.DataSource)
	_, err := l.svcCtx.CommentRepliesModel.Insert(l.ctx, &dynamicsmodels.CommentReplies{
		CommentReplieId: commentReplyId,
		CommentId:       in.CommentId,
		UserId:          in.UserId,
		TargetUserId:    findOne.UserId,
		Content:         in.Content,
		IsDeleted:       false,
	})
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "dynamics-rpc CreateCommentReplay")
	}

	return &dynamics.CreateCommentReplayResp{
		CommentReplyId: commentReplyId,
		PostId:         findOne.PostId,
	}, nil
}
