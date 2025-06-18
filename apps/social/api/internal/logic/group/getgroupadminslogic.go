package group

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/socialservice"

	"github.com/ljp-lachouchou/chan_xin/apps/social/api/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/social/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupAdminsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取管理员
func NewGetGroupAdminsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupAdminsLogic {
	return &GetGroupAdminsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGroupAdminsLogic) GetGroupAdmins(req *types.GetGroupMembersReq) (resp *types.GetGroupMembersResp, err error) {
	in := &socialservice.GetGroupMembersReq{
		GroupId: req.GroupId,
	}
	admins, err := l.svcCtx.SocialService.GetGroupAdmins(l.ctx, in)
	if err != nil {
		return nil, err
	}
	var respList []*types.BaseUserInfo
	for _, v := range admins.List {
		respOne := &types.BaseUserInfo{
			UserId:    v.UserId,
			Nickname:  v.Nickname,
			AvatarUrl: v.AvatarUrl,
			Gender:    v.Gender,
		}
		respList = append(respList, respOne)
	}

	return &types.GetGroupMembersResp{
		List: respList,
	}, nil
}
