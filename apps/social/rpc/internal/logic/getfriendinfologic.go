package logic

import (
	"context"

	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/social"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFriendInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFriendInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendInfoLogic {
	return &GetFriendInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFriendInfoLogic) GetFriendInfo(in *social.FriendInfoRequest) (*social.UserInfo, error) {
	// todo: add your logic here and delete this line

	return &social.UserInfo{}, nil
}
