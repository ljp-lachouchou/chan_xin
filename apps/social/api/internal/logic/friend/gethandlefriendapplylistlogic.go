package friend

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/socialservice"

	"github.com/ljp-lachouchou/chan_xin/apps/social/api/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/social/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetHandleFriendApplyListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取谁想添加我为好友的列表
func NewGetHandleFriendApplyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHandleFriendApplyListLogic {
	return &GetHandleFriendApplyListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetHandleFriendApplyListLogic) GetHandleFriendApplyList(req *types.HandleFriendApplyReq) (*types.FriendApplyListResp, error) {
	in := &socialservice.HandleFriendApplyReq{
		TargetId: req.TargetId,
	}
	list, err := l.svcCtx.SocialService.GetHandleFriendApplyList(l.ctx, in)
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
