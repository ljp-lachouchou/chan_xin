package friend

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/socialservice"

	"github.com/ljp-lachouchou/chan_xin/apps/social/api/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/social/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFriendInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取好友信息
func NewGetFriendInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendInfoLogic {
	return &GetFriendInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFriendInfoLogic) GetFriendInfo(req *types.FriendInfoRequest) (resp *types.UserInfo, err error) {
	in := &socialservice.FriendInfoRequest{
		UserId:   req.UserId,
		FriendId: req.FriendId,
	}
	info, err := l.svcCtx.SocialService.GetFriendInfo(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return &types.UserInfo{
		UserId:    info.UserId,
		Nickname:  info.Nickname,
		AvatarUrl: info.AvatarUrl,
		Gender:    info.Gender,
		Status: types.FriendStatusInfo{
			IsMuted:   info.Status.IsMuted,
			IsTopped:  info.Status.IsTopped,
			IsBlocked: info.Status.IsBlocked,
			Remark:    info.Status.Remark,
		},
	}, nil
}
