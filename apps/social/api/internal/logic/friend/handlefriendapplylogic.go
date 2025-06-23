package friend

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/im/rpc/im"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/socialservice"
	"github.com/ljp-lachouchou/chan_xin/deploy/constant"

	"github.com/ljp-lachouchou/chan_xin/apps/social/api/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/social/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HandleFriendApplyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 目标方处理申请
func NewHandleFriendApplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HandleFriendApplyLogic {
	return &HandleFriendApplyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HandleFriendApplyLogic) HandleFriendApply(req *types.FriendApplyAction) error {
	in := socialservice.FriendApplyAction{
		ApplyId:    req.ApplyId,
		IsApproved: req.IsApproved,
	}
	resp, err := l.svcCtx.SocialService.HandleFriendApply(l.ctx, &in)
	if err != nil {
		return err
	}
	if resp.IsApproved == false {
		return nil
	}
	_, err = l.svcCtx.Im.SetUpUserConversation(l.ctx, &im.SetUpUserConversationReq{
		SendId:   resp.ApplicantId,
		RecvId:   resp.TargetId,
		ChatType: int32(constant.SingleChat),
	})
	return err
}
