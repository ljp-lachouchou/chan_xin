package logic

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/social/socialmodels"
	"github.com/ljp-lachouchou/chan_xin/pkg/lerr"
	"github.com/pkg/errors"

	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/social"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateFriendStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateFriendStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFriendStatusLogic {
	return &UpdateFriendStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateFriendStatusLogic) UpdateFriendStatus(in *social.FriendStatusUpdate) (*social.FriendStatusUpdateResp, error) {
	friend, err := l.svcCtx.FriendRelationModel.FindOneByUserIdFriendId(l.ctx, in.UserId, in.FriendId)
	if err != nil {
		if err == socialmodels.ErrNotFound {
			return nil, errors.WithStack(NoFriendErr)
		}
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "social-rpc GetFriendInfo FriendRelationModel.FindOneByUserIdFriendId", in.UserId, in.FriendId)
	}
	if in.Status == nil {
		return &social.FriendStatusUpdateResp{}, nil
	}
	if in.Status.IsBlocked != nil {
		if *in.Status.IsBlocked {
			friend.IsBlocked = 1
		} else {
			friend.IsBlocked = 0
		}
	}
	if in.Status.IsMuted != nil {
		if *in.Status.IsMuted {
			friend.IsMuted = 1
		} else {
			friend.IsMuted = 0
		}
	}
	if in.Status.IsTopped != nil {
		if *in.Status.IsTopped {
			friend.IsTopped = 1
		} else {
			friend.IsTopped = 0
		}
	}
	if in.Status.Remark != nil {
		friend.Remark = *in.Status.Remark
	}

	l.Info("friend ", friend)
	err = l.svcCtx.FriendRelationModel.Update(l.ctx, friend)
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "social-rpc UpdateFriendStatusLogic.UpdateFriendStatus", friend)
	}
	return &social.FriendStatusUpdateResp{}, nil
}
