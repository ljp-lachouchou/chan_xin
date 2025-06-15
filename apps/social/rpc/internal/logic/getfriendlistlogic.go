package logic

import (
	"context"
	"fmt"
	"github.com/ljp-lachouchou/chan_xin/pkg/lerr"
	"github.com/ljp-lachouchou/chan_xin/pkg/ltool"

	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/social"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFriendListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendListLogic {
	return &GetFriendListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFriendListLogic) GetFriendList(in *social.FriendListReq) (*social.FriendListResp, error) {
	listByUserid, err := l.svcCtx.FriendRelationModel.ListByUserIdWithUsers(l.ctx, in.UserId)
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "social rpc GetFriendList", in.UserId)

	}
	var resp []*social.UserInfo
	for _, v := range listByUserid {

		resp = append(resp, &social.UserInfo{
			UserId:    v.Id,
			Nickname:  v.Nickname,
			AvatarUrl: v.Avatar,
			Gender:    v.Sex,
			Status: &social.FriendStatusInfo{
				IsMuted:   ltool.IntConvBool(int(v.IsMuted)),
				IsTopped:  ltool.IntConvBool(int(v.IsTopped)),
				IsBlocked: ltool.IntConvBool(int(v.IsBlocked)),
				Remark:    v.Remark,
			},
		})
	}
	fmt.Println("listByUserid:", listByUserid)
	fmt.Println("resp:", resp)
	return &social.FriendListResp{
		FriendList: resp,
	}, nil
}
