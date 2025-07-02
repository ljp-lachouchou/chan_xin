package logic

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/social/socialmodels"
	"github.com/ljp-lachouchou/chan_xin/deploy/constant"
	"github.com/ljp-lachouchou/chan_xin/pkg/lerr"

	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/social"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGroupInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupInfoLogic {
	return &GetGroupInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGroupInfoLogic) GetGroupInfo(in *social.GroupInfoRequest) (*social.GroupInfo, error) {
	//首先检查是否为群内成员
	memberInGroup, err := l.svcCtx.GroupMemberModel.FindOneByGroupIdUserId(l.ctx, in.GroupId, in.UserId)
	if err != nil && err != socialmodels.ErrNotFound {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "social-rpc GetGroupInfo GroupMemberModel.FindOneByGroupIdUserId", in.GroupId, in.UserId)
	}
	if memberInGroup == nil {
		//不是群内成员
		group, err := l.svcCtx.GroupInfoModel.FindOne(l.ctx, in.GroupId)
		if err != nil {
			return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "social-rpc GetGroupInfo GroupInfoModel.FindOne", in.GroupId)
		}
		return &social.GroupInfo{
			GroupId:    group.GroupId,
			Name:       group.Name,
			OwnerId:    group.OwnerId,
			MaxMembers: uint32(group.MaxMembers),
		}, nil
	}
	group, err := l.svcCtx.GroupInfoModel.FindOne(l.ctx, in.GroupId)
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "social-rpc GetGroupInfo GroupInfoModel.FindOne", in.GroupId)
	}
	members, err := l.svcCtx.GroupMemberModel.FindGroupMembers(l.ctx, in.GroupId)
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "social-rpc GetGroupInfo GroupMemberModel.FindGroupMembers", in.GroupId)
	}
	var (
		memberIds []string
		adminIds  []string
	)
	for _, member := range members {
		memberIds = append(memberIds, member.UserId)
		if constant.IsAdminInGroup(member.IsAdmin) == constant.AdminInGroup {
			adminIds = append(adminIds, member.UserId)
		}
	}
	return &social.GroupInfo{
		GroupId:    group.GroupId,
		Name:       group.Name,
		OwnerId:    group.OwnerId,
		AdminIds:   adminIds,
		MemberIds:  memberIds,
		MaxMembers: uint32(group.MaxMembers),
	}, nil
}
