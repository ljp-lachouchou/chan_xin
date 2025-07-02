package logic

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/pkg/lerr"

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
	member, err := l.svcCtx.GroupMemberModel.FindOneByGroupIdUserId(l.ctx, in.GroupId, in.UserId)
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "social-rpc SetGroupMemberSetting GroupMemberModel.FindOneByGroupIdUserId", in.GroupId, in.UserId)
	}
	if in.Setting == nil {
		return &social.GroupMemberSettingUpdateResp{}, nil
	}

	if in.Setting.GroupNickname != nil {
		member.GroupNickname = *in.Setting.GroupNickname
	}
	if in.Setting.ShowMemberNickname != nil {
		if *in.Setting.ShowMemberNickname {
			member.ShowNickname = 1
		} else {
			member.ShowNickname = 0
		}
	}
	err = l.svcCtx.GroupMemberModel.Update(l.ctx, member)
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "social-rpc SetGroupMemberSetting GroupMemberModel.Update", member)
	}
	return &social.GroupMemberSettingUpdateResp{}, nil
}
