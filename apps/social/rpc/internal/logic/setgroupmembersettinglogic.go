package logic

import (
	"context"

	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/social"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetGroupMemberSettingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetGroupMemberSettingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetGroupMemberSettingLogic {
	return &SetGroupMemberSettingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetGroupMemberSettingLogic) SetGroupMemberSetting(in *social.GroupMemberSettingUpdate) (*social.GroupMemberSettingUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &social.GroupMemberSettingUpdateResp{}, nil
}
