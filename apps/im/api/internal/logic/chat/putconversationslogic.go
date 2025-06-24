package chat

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/im/rpc/im"

	"github.com/ljp-lachouchou/chan_xin/apps/im/api/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/im/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutConversationsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutConversationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutConversationsLogic {
	return &PutConversationsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutConversationsLogic) PutConversations(req *types.PutConversationsReq) (*types.PutConversationsResp, error) {
	list := make(map[string]*im.Conversation, len(req.ConversationList))
	for k, v := range req.ConversationList {
		list[k] = &im.Conversation{
			ConversationId: v.ConversationId,
			ChatType:       v.ChatType,
			TargetId:       v.TargetId,
			IsShow:         v.IsShow,
			Seq:            v.Seq,
			Total:          v.Total,
			ToRead:         v.ToRead,
			Read:           v.Read,
			Msg: &im.ChatLog{
				Id:             v.Msg.Id,
				ConversationId: v.Msg.ConversationId,
				SendId:         v.Msg.SendId,
				RecvId:         v.Msg.RecvId,
				MsgType:        v.Msg.MsgType,
				MsgContent:     v.Msg.MsgContent,
				ChatType:       v.Msg.ChatType,
				SendTime:       v.Msg.SendTime,
			},
		}
	}
	_, err := l.svcCtx.PutConversations(l.ctx, &im.PutConversationsReq{
		Id:               req.Id,
		UserId:           req.UserId,
		ConversationList: list,
	})
	if err != nil {
		return nil, err
	}
	res := &types.PutConversationsResp{}
	res.Success = true
	return res, nil
}
