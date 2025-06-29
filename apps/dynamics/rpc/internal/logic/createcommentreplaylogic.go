package logic

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/dynamicsmodels"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/dynamics"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/pkg/lerr"
	"github.com/ljp-lachouchou/chan_xin/pkg/wuid"

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
func (l *CreateCommentReplayLogic) CreateCommentReplay(in *dynamics.CreateCommentReplayReq) (*dynamics.Empty, error) {
	_, err := l.svcCtx.CommentRepliesModel.Insert(l.ctx, &dynamicsmodels.CommentReplies{
		CommentReplieId: wuid.GenUid(l.svcCtx.Config.Mysql.DataSource),
		CommentId:       in.CommentId,
		UserId:          in.UserId,
		TargetUserId:    in.TargetUserId,
		Content:         in.Content,
		IsDeleted:       false,
	})
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "dynamics-rpc CreateCommentReplay")
	}
	return &dynamics.Empty{}, nil
}
