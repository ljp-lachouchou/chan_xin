package friend

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/socialservice"

	"github.com/ljp-lachouchou/chan_xin/apps/social/api/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/social/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFriendListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户好友列表
func NewGetFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendListLogic {
	return &GetFriendListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFriendListLogic) GetFriendList(req *types.FriendListReq) (resp *types.FriendListResp, err error) {
	in := &socialservice.FriendListReq{
		UserId: req.UserId,
	}
	list, err := l.svcCtx.SocialService.GetFriendList(l.ctx, in)
	if err != nil {
		return nil, err
	}
	var listResp []*types.UserInfo
	for _, v := range list.FriendList {
		resp := &types.UserInfo{
			UserId:    v.UserId,
			Nickname:  v.Nickname,
			AvatarUrl: v.AvatarUrl,
			Gender:    v.Gender,
			Status: types.FriendStatusInfo{
				IsMuted:   v.Status.IsMuted,
				IsTopped:  v.Status.IsTopped,
				IsBlocked: v.Status.IsBlocked,
				Remark:    v.Status.Remark,
			},
		}
		listResp = append(listResp, resp)
	}
	return &types.FriendListResp{
		FriendList: listResp,
	}, nil
}
