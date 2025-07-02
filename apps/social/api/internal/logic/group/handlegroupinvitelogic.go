package group

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/im/rpc/im"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/socialservice"
	"github.com/ljp-lachouchou/chan_xin/deploy/constant"

	"github.com/ljp-lachouchou/chan_xin/apps/social/api/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/social/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HandleGroupInviteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 被邀请者处理群申请
func NewHandleGroupInviteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HandleGroupInviteLogic {
	return &HandleGroupInviteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HandleGroupInviteLogic) HandleGroupInvite(req *types.GroupInviteAction) error {
	in := &socialservice.GroupInviteAction{
		InviteId:   req.InviteId,
		IsAccepted: req.IsAccepted,
	}
	resp, err := l.svcCtx.SocialService.HandleGroupInvite(l.ctx, in)
	if err != nil {
		return err
	}
	_, err = l.svcCtx.Im.SetUpUserConversation(l.ctx, &im.SetUpUserConversationReq{
		SendId:   resp.ApplicationId,
		RecvId:   resp.TargetId,
		ChatType: int32(constant.GroupChat),
	})
	return err
}
