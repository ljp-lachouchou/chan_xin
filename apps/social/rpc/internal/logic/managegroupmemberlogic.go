package logic

import (
	"context"

	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/social"

	"github.com/zeromicro/go-zero/core/logx"
)

type ManageGroupMemberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewManageGroupMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ManageGroupMemberLogic {
	return &ManageGroupMemberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ManageGroupMemberLogic) ManageGroupMember(in *social.GroupMemberManage) (*social.GroupMemberManageResp, error) {
	// todo: add your logic here and delete this line

	return &social.GroupMemberManageResp{}, nil
}
