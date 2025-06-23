package logic

import (
	"context"

	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/social"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/w/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupAdminsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGroupAdminsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupAdminsLogic {
	return &GetGroupAdminsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGroupAdminsLogic) GetGroupAdmins(in *social.GetGroupMembersReq) (*social.GetGroupMembersResp, error) {
	// todo: add your logic here and delete this line

	return &social.GetGroupMembersResp{}, nil
}
