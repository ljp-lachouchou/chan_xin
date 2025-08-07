package logic

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/dynamicsmodels"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/dynamics"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/pkg/lerr"
	"sort"

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
	list, err := l.svcCtx.CommentsModel.FindByPostId(l.ctx, in.PostId)
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "ListCommentByPostId CommentsModel.FindByPostId ", in.PostId)
	}
	var resp []*dynamics.ListCommentRespStruct
	var commentIds []string
	listOne := make([]*dynamicsmodels.CommentReplies, len(list))
	for i, v := range list {
		listOne[i] = &dynamicsmodels.CommentReplies{
			CommentId:    v.CommentId,
			UserId:       v.UserId,
			TargetUserId: "",
			Content:      v.Content,
			CreatedAt:    v.CreatedAt,
			IsDeleted:    v.IsDeleted,
		}
		commentIds = append(commentIds, v.CommentId)
	}
	replies, err := l.svcCtx.CommentRepliesModel.FindByCommentIds(l.ctx, commentIds...)
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "ListCommentByPostId CommentRepliesModel.FindByCommentIds ", commentIds)
	}
	for _, v := range replies {
		listOne = append(listOne, v)
	}
	sort.Slice(listOne, func(i, j int) bool {
		return listOne[i].CreatedAt.Unix() < listOne[j].CreatedAt.Unix()
	})
	for _, v := range listOne {
		resp = append(resp, &dynamics.ListCommentRespStruct{
			CommentId:    v.CommentId,
			UserId:       v.UserId,
			TargetUserId: v.TargetUserId,
			Content:      v.Content,
		})
	}
	return &dynamics.ListCommentResp{
		List: resp,
	}, nil
}
