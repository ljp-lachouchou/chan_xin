package logic

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/pkg/lerr"

	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/social"

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
	admins, err := l.svcCtx.GroupMemberModel.FindGroupAdmins(l.ctx, in.GroupId)
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "social-rpc GetGroupAdmins GroupMemberModel.FindGroupMembers", in.GroupId)
	}
	var list []*social.BaseUserInfo
	for _, admin := range admins {
		findOne, err := l.svcCtx.UsersModel.FindOne(l.ctx, admin.UserId)
		if err != nil {
			return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "social-rpc GetGroupAdmins UsersModel.FindOne", admin.UserId)
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
