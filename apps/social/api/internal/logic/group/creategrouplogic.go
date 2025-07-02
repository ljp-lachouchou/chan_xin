package group

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/im/rpc/im"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/socialservice"
	"github.com/ljp-lachouchou/chan_xin/deploy/constant"

	"github.com/ljp-lachouchou/chan_xin/apps/social/api/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/social/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建群
func NewCreateGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateGroupLogic {
	return &CreateGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateGroupLogic) CreateGroup(req *types.GroupCreationRequest) (resp *types.GroupInfo, err error) {
	in := &socialservice.GroupCreationRequest{
		CreatorId: req.CreatorId,
		GroupName: req.GroupName,
		MemberIds: req.MemberIds,
	}

	group, err := l.svcCtx.SocialService.CreateGroup(l.ctx, in)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.Im.CreateGroupConversation(l.ctx, &im.CreateGroupConversationReq{
		GroupId:  group.GroupId,
		CreateId: group.OwnerId,
	})
	if err != nil {
		return nil, err
	}
	ids := req.MemberIds
	ids = append(ids, req.CreatorId)
	for _, id := range ids {
		l.svcCtx.Im.SetUpUserConversation(l.ctx, &im.SetUpUserConversationReq{
			SendId:   id,
			RecvId:   group.GroupId,
			ChatType: int32(constant.GroupChat),
		})
	}
	return &types.GroupInfo{
		GroupId:    group.GroupId,
		Name:       group.Name,
		OwnerId:    group.OwnerId,
		AdminIds:   group.AdminIds,
		MemberIds:  group.MemberIds,
		MaxMembers: group.MaxMembers,
	}, nil
}
