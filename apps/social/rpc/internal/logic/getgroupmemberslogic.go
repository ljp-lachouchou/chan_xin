package logic

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/pkg/lerr"

	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/social"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupMembersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGroupMembersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupMembersLogic {
	return &GetGroupMembersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGroupMembersLogic) GetGroupMembers(in *social.GetGroupMembersReq) (*social.GetGroupMembersResp, error) {
	members, err := l.svcCtx.GroupMemberModel.FindGroupMembers(l.ctx, in.GroupId)
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "social-rpc GetGroupMembers", in.GroupId)
	}
	var list []*social.BaseUserInfo
	for _, member := range members {
		findOne, err := l.svcCtx.UsersModel.FindOne(l.ctx, member.UserId)
		if err != nil {
			return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "social-rpc GetGroupMembers UsersModel.FindOne", member.UserId)
		}
		list = append(list, &social.BaseUserInfo{
			UserId:    findOne.Id,
			Nickname:  findOne.Nickname,
			AvatarUrl: findOne.Avatar,
			Gender:    uint32(findOne.Sex.Int64),
		})
	}
	return &social.GetGroupMembersResp{
		List: list,
	}, nil
}
