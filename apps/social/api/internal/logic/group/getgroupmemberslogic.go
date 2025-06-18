package group

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/socialservice"

	"github.com/ljp-lachouchou/chan_xin/apps/social/api/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/social/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupMembersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取群里成员
func NewGetGroupMembersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupMembersLogic {
	return &GetGroupMembersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGroupMembersLogic) GetGroupMembers(req *types.GetGroupMembersReq) (resp *types.GetGroupMembersResp, err error) {
	in := &socialservice.GetGroupMembersReq{
		GroupId: req.GroupId,
	}
	members, err := l.svcCtx.SocialService.GetGroupMembers(l.ctx, in)
	if err != nil {
		return nil, err
	}
	var respList []*types.BaseUserInfo
	for _, v := range members.List {
		resoOne := &types.BaseUserInfo{
			UserId:    v.UserId,
			Nickname:  v.Nickname,
			AvatarUrl: v.AvatarUrl,
			Gender:    v.Gender,
		}
		respList = append(respList, resoOne)
	}

	return &types.GetGroupMembersResp{
		List: respList,
	}, nil
}
