package friend

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/socialservice"

	"github.com/ljp-lachouchou/chan_xin/apps/social/api/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/social/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFriendApplyListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取你请求添加的好友列表
func NewGetFriendApplyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendApplyListLogic {
	return &GetFriendApplyListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFriendApplyListLogic) GetFriendApplyList(req *types.FriendApplyListReq) (resp *types.FriendApplyListResp, err error) {
	in := &socialservice.FriendApplyListReq{
		UserId: req.UserId,
	}
	list, err := l.svcCtx.SocialService.GetFriendApplyList(l.ctx, in)
	if err != nil {
		return nil, err
	}
	var respList []*types.FriendApplyResp
	for _, v := range list.List {
		one := &types.FriendApplyResp{
			UserId:    v.UserId,
			Nickname:  v.Nickname,
			AvatarUrl: v.AvatarUrl,
			Gender:    v.Gender,
			GreetMsg:  v.GreetMsg,
			Status:    v.Status,
		}
		respList = append(respList, one)
	}
	return &types.FriendApplyListResp{
		List: respList,
	}, nil
}
