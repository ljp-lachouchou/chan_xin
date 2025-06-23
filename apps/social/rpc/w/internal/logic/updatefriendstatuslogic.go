package logic

import (
	"context"

	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/social"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/w/internal/svc"

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
	// todo: add your logic here and delete this line

	return &social.FriendStatusUpdateResp{}, nil
}
