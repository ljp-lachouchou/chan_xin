package logic

import (
	"context"
	"fmt"
	"github.com/ljp-lachouchou/chan_xin/pkg/lerr"

	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/social"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetHandleFriendApplyListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetHandleFriendApplyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHandleFriendApplyListLogic {
	return &GetHandleFriendApplyListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetHandleFriendApplyListLogic) GetHandleFriendApplyList(in *social.HandleFriendApplyReq) (*social.FriendApplyListResp, error) {
	listByUserid, err := l.svcCtx.FriendApplyModel.ListByTargetIdJoinUsers(l.ctx, in.TargetId)
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "social rpc GetFriendList", in.TargetId)

	}
	var resp []*social.FriendApplyResp
	for _, v := range listByUserid {

		resp = append(resp, &social.FriendApplyResp{
			UserId:    v.Id,
			Nickname:  v.Nickname,
			AvatarUrl: v.Avatar,
			Gender:    v.Sex,
			GreetMsg:  v.GreetMsg,
			Status:    int32(v.Status),
		})
	}
	fmt.Println("listByUserid:", listByUserid)
	fmt.Println("resp:", resp)
	return &social.FriendApplyListResp{
		List: resp,
	}, nil
}
