package logic

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/dynamicsmodels"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/dynamics"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/pkg/lerr"
	"github.com/zeromicro/go-zero/core/logx"
)

type ListCommentByPostIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListCommentByPostIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCommentByPostIdLogic {
	return &ListCommentByPostIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 评论列表
func (l *ListCommentByPostIdLogic) ListCommentByPostId(in *dynamics.GetPostInfoReq) (*dynamics.ListCommentResp, error) {
	list, err := l.svcCtx.CommentRepliesModel.FindByPostId(l.ctx, in.PostId)
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "ListCommentByPostId CommentsModel.FindByPostId ", in.PostId)
	}
	var resp []*dynamics.ListCommentRespStruct
	listOne := make([]*dynamicsmodels.CommentReplies, len(list))
	for i, v := range list {
		listOne[i] = &dynamicsmodels.CommentReplies{
			CommentReplieId: v.CommentReplieId,
			CommentId:       v.CommentId,
			UserId:          v.UserId,
			TargetUserId:    v.TargetUserId,
			Content:         v.Content,
			CreatedAt:       v.CreatedAt,
			IsDeleted:       v.IsDeleted,
		}
	}
	for _, v := range listOne {
		resp = append(resp, &dynamics.ListCommentRespStruct{
			CommentId:    v.CommentReplieId,
			UserId:       v.UserId,
			TargetUserId: v.TargetUserId,
			Content:      v.Content,
		})
	}
	return &dynamics.ListCommentResp{
		List: resp,
	}, nil
}
