package friend

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/socialservice"

	"github.com/ljp-lachouchou/chan_xin/apps/social/api/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/social/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApplyFriendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 申请好友
func NewApplyFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApplyFriendLogic {
	return &ApplyFriendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApplyFriendLogic) ApplyFriend(req *types.FriendApplyRequest) (resp *types.FriendApplyResponse, err error) {
	in := socialservice.FriendApplyRequest{
		ApplicantId: req.ApplicantId,
		TargetId:    req.TargetId,
		GreetMsg:    req.GreetMsg,
	}
	response, err := l.svcCtx.SocialService.ApplyFriend(l.ctx, &in)
	if err != nil {
		return nil, err
	}

	return &types.FriendApplyResponse{
		ApplyId:   response.ApplyId,
		ApplyTime: response.ApplyTime,
	}, nil
}
