package chat

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/im/rpc/im"

	"github.com/ljp-lachouchou/chan_xin/apps/im/api/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/im/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetConversationsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetConversationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConversationsLogic {
	return &GetConversationsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetConversationsLogic) GetConversations(req *types.GetConversationsReq) (*types.GetConversationsResp, error) {
	conversations, err := l.svcCtx.GetConversations(l.ctx, &im.GetConversationsReq{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, err
	}
	mapList := make(map[string]types.Conversation, len(conversations.ConversationList))
	for k, v := range conversations.ConversationList {
		e := types.Conversation{
			ConversationId: v.ConversationId,
			ChatType:       v.ChatType,
			TargetId:       v.TargetId,
			IsShow:         v.IsShow,
			Seq:            v.Seq,
			Total:          v.Total,
			ToRead:         v.ToRead,
			Read:           v.Read,
		}
		if v.Msg != nil {
			e.Msg = types.ChatLog{
				Id:             v.Msg.Id,
				ConversationId: v.Msg.ConversationId,
				SendId:         v.Msg.SendId,
				RecvId:         v.Msg.RecvId,
				MsgType:        v.Msg.MsgType,
				MsgContent:     v.Msg.MsgContent,
				ChatType:       v.Msg.ChatType,
				SendTime:       v.Msg.SendTime,
			}
		}
		mapList[k] = e
	}
	return &types.GetConversationsResp{
		ConversationList: mapList,
	}, nil
}
