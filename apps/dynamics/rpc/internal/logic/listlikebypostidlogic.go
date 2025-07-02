package logic

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/pkg/lerr"

	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/dynamics"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLikeByPostIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListLikeByPostIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLikeByPostIdLogic {
	return &ListLikeByPostIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 点赞列表
func (l *ListLikeByPostIdLogic) ListLikeByPostId(in *dynamics.GetPostInfoReq) (*dynamics.Ids, error) {
	list, err := l.svcCtx.PostLikesModel.FindByPostId(l.ctx, in.PostId)
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "dynamics-rpc ListLikeByPostId PostLikesModel.FindByPostId ", in.PostId)
	}

	ids := make([]string, len(list))
	for i, v := range list {
		if v.UserId == "" {
			continue
		}
		ids[i] = v.UserId
	}
	l.Info("ListLikeByPostId ids size:", len(ids), ids)
	return &dynamics.Ids{
		Ids: ids,
	}, nil
}
