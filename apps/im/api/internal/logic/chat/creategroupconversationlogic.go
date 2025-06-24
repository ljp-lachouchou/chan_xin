package chat

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/im/rpc/im"

	"github.com/ljp-lachouchou/chan_xin/apps/im/api/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/im/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateGroupConversationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateGroupConversationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateGroupConversationLogic {
	return &CreateGroupConversationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateGroupConversationLogic) CreateGroupConversation(req *types.CreateGroupConversationReq) error {
	_, err := l.svcCtx.Im.CreateGroupConversation(l.ctx, &im.CreateGroupConversationReq{
		GroupId:  req.GroupId,
		CreateId: req.CreateId,
	})
	return err
}
