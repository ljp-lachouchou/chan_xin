package logic

import (
	"context"
	"fmt"
	"github.com/ljp-lachouchou/chan_xin/apps/social/socialmodels"
	"github.com/ljp-lachouchou/chan_xin/pkg/lerr"
	"github.com/ljp-lachouchou/chan_xin/pkg/ltool"
	"github.com/pkg/errors"

	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/social"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	NoFriendErr = errors.New("请先添加对方为好友")
)

type GetFriendInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFriendInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendInfoLogic {
	return &GetFriendInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFriendInfoLogic) GetFriendInfo(in *social.FriendInfoRequest) (*social.UserInfo, error) {
	friend, err := l.svcCtx.FriendRelationModel.FindOneByUserIdFriendId(l.ctx, in.UserId, in.FriendId)
	if err != nil {
		if err == socialmodels.ErrNotFound {
			return nil, errors.WithStack(NoFriendErr)
		}
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "social-rpc GetFriendInfo FriendRelationModel.FindOneByUserIdFriendId", in.UserId, in.FriendId)
	}
	friendInfo, err := l.svcCtx.UsersModel.FindOne(l.ctx, friend.FriendId)
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "social-rpc GetFriendInfo UserModel.FindOne", friend.FriendId)
	}
	fmt.Println("social-rpc getinfo fff", friend)
	status := &social.FriendStatusInfo{
		IsMuted:   ltool.IntConvBool(int(friend.IsMuted)),
		IsTopped:  ltool.IntConvBool(int(friend.IsTopped)),
		IsBlocked: ltool.IntConvBool(int(friend.IsBlocked)),
		Remark:    friend.Remark,
	}
	fmt.Println("social-rpc getinfo fffs", status)
	return &social.UserInfo{
		UserId:    friendInfo.Id,
		Nickname:  friendInfo.Nickname,
		AvatarUrl: friendInfo.Avatar,
		Gender:    uint32(friendInfo.Sex.Int64),
		Status:    status,
	}, nil
}
