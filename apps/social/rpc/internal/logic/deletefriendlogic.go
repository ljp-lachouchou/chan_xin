package logic

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/pkg/lerr"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/social"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteFriendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFriendLogic {
	return &DeleteFriendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteFriendLogic) DeleteFriend(in *social.RelationRequest) (*social.RelationResp, error) {
	err := l.svcCtx.FriendRelationModel.Transx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		if err := l.svcCtx.FriendApplyModel.DeleteByUidAndFId(l.ctx, session, in.FromUid, in.ToUid); err != nil {
			return lerr.NewWrapError(lerr.NEWDBError(), err, "social-rpc HandleFriendApply", in.FromUid, in.ToUid)
		}
		return l.svcCtx.FriendRelationModel.DeleteByUserIdFriendId(l.ctx, session, in.FromUid, in.ToUid)
	})
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "DeleteFriend FriendRelationModel.Transx", in.ToUid, in.FromUid)
	}
	return &social.RelationResp{}, nil
}
