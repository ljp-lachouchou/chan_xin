package logic

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/dynamics"

	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/api/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListCommentByPostIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 评论列表
func NewListCommentByPostIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCommentByPostIdLogic {
	return &ListCommentByPostIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListCommentByPostIdLogic) ListCommentByPostId(req *types.GetPostInfoReq) (*types.ListCommentResp, error) {
	ls, err := l.svcCtx.Dynamics.ListCommentByPostId(l.ctx, &dynamics.GetPostInfoReq{
		PostId: req.PostId,
	})
	if err != nil {
		return nil, err
	}
	var list []types.ListCommentRespStruct
	for _, v := range ls.List {
		list = append(list, types.ListCommentRespStruct{
			CommentId: v.CommentId,
			UserId:       v.UserId,
			TargetUserId: v.TargetUserId,
			Content:      v.Content,
		})
	}
	return &types.ListCommentResp{
		List: list,
	}, nil
}
