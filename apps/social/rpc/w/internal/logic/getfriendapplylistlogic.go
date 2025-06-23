package logic

import (
	"context"

	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/social"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/w/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFriendApplyListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFriendApplyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendApplyListLogic {
	return &GetFriendApplyListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFriendApplyListLogic) GetFriendApplyList(in *social.FriendApplyListReq) (*social.FriendApplyListResp, error) {
	// todo: add your logic here and delete this line

	return &social.FriendApplyListResp{}, nil
}
