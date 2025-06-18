package logic

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/social"
	"github.com/ljp-lachouchou/chan_xin/apps/social/socialmodels"
	"github.com/ljp-lachouchou/chan_xin/pkg/ldefault"
	"github.com/ljp-lachouchou/chan_xin/pkg/lerr"
	"github.com/ljp-lachouchou/chan_xin/pkg/wuid"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type CreateGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateGroupLogic {
	return &CreateGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// === 群组管理接口 ===
func (l *CreateGroupLogic) CreateGroup(in *social.GroupCreationRequest) (*social.GroupInfo, error) {
	groupId := wuid.GenUid(l.svcCtx.Config.Mysql.DataSource)
	info := &socialmodels.GroupInfo{
		GroupId:    groupId,
		Name:       in.GroupName,
		OwnerId:    in.CreatorId,
		MaxMembers: ldefault.DEFAULT_MAX_GROUP_MEMBERS,
	}
	_, err := l.svcCtx.GroupInfoModel.Insert(l.ctx, info)
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "social-rpc CreateGroup", info)
	}
	err = l.svcCtx.GroupMemberModel.Transx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		var members []*socialmodels.GroupMember
		member := &socialmodels.GroupMember{
			GroupId:       groupId,
			UserId:        in.CreatorId,
			GroupNickname: "",
			ShowNickname:  1,
			IsAdmin:       1,
			IsMuted:       0,
			IsTopped:      0,
		}
		members = append(members, member)
		for _, v := range in.MemberIds {
			member := &socialmodels.GroupMember{
				GroupId:       groupId,
				UserId:        v,
				GroupNickname: "",
				ShowNickname:  1,
				IsAdmin:       0,
				IsMuted:       0,
				IsTopped:      0,
			}
			members = append(members, member)
		}
		_, err2 := l.svcCtx.GroupMemberModel.InsertMembers(l.ctx, session, members...)
		return err2
	})
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "social-rpc CreateGroup Transx")
	}
	return &social.GroupInfo{
		GroupId:    in.CreatorId,
		Name:       in.GroupName,
		OwnerId:    info.OwnerId,
		MemberIds:  in.MemberIds,
		MaxMembers: uint32(info.MaxMembers),
	}, nil
}
