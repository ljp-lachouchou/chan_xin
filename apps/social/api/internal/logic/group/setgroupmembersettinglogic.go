package group

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/socialservice"

	"github.com/ljp-lachouchou/chan_xin/apps/social/api/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/social/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetGroupMemberSettingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 群员的个性化设置
func NewSetGroupMemberSettingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetGroupMemberSettingLogic {
	return &SetGroupMemberSettingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetGroupMemberSettingLogic) SetGroupMemberSetting(req *types.GroupMemberSettingUpdate) error {
	in := &socialservice.GroupMemberSettingUpdate{
		UserId:  req.UserId,
		GroupId: req.GroupId,
		Setting: &socialservice.GroupMemberSetting{
			GroupNickname:      &req.Setting.GroupNickname,
			ShowMemberNickname: &req.Setting.ShowMemberNickname,
		},
	}
	_, err := l.svcCtx.SocialService.SetGroupMemberSetting(l.ctx, in)

	return err
}
