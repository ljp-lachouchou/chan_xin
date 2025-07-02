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

type CreateCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCommentLogic {
	return &CreateCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建评论
func (l *CreateCommentLogic) CreateComment(in *dynamics.CreateCommentReq) (*dynamics.CreateCommentResp, error) {
	commentId := wuid.GenUid(l.svcCtx.Config.Mysql.DataSource)
	_, err := l.svcCtx.CommentsModel.Insert(l.ctx, &dynamicsmodels.Comments{
		CommentId: commentId,
		PostId:    in.PostId,
		UserId:    in.UserId,
		Content:   in.Content,
		IsDeleted: false,
	})
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "dynamics-rpc CreateComment")
	}
	return &dynamics.CreateCommentResp{
		CommentId: commentId,
	}, nil
}
