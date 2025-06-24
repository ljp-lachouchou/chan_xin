package logic

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/ljp-lachouchou/chan_xin/apps/im/immodels"
	"github.com/ljp-lachouchou/chan_xin/pkg/lerr"

	"github.com/ljp-lachouchou/chan_xin/apps/im/rpc/im"
	"github.com/ljp-lachouchou/chan_xin/apps/im/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetConversationsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetConversationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConversationsLogic {
	return &GetConversationsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取会话
func (l *GetConversationsLogic) GetConversations(in *im.GetConversationsReq) (*im.GetConversationsResp, error) {
	conversations, err := l.svcCtx.ConversationsModel.FindByUserId(l.ctx, in.UserId)
	if err != nil {
		if err == immodels.ErrNotFound {
			return &im.GetConversationsResp{}, nil
		}
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "im-rpc GetConversations FindByUserId ", in.UserId)
	}
	var res im.GetConversationsResp
	err = copier.Copy(&res, &conversations)
	if err != nil {
		return nil, err
	}
	ids := make([]string, 0, len(conversations.ConversationList))
	for _, conversation := range conversations.ConversationList {
		ids = append(ids, conversation.ConversationId)
	}
	newConversations, err := l.svcCtx.ConversationModel.ListByConversationIds(l.ctx, ids)
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "im-rpc GetConversations ListByConversationIds ", ids)
	}
	for _, conversation := range newConversations {
		if _, ok := res.ConversationList[conversation.ConversationId]; !ok {
			continue
		}
		oldTotal := conversations.ConversationList[conversation.ConversationId].Total
		if oldTotal < conversation.Total {
			res.ConversationList[conversation.ConversationId].Total = int32(conversation.Total)
			res.ConversationList[conversation.ConversationId].ToRead = int32(conversation.Total - oldTotal)
			res.ConversationList[conversation.ConversationId].IsShow = true
			if conversation.LastMsg == nil {
				continue
			}
			fmt.Println("lastMsg", conversation.LastMsg)
			res.ConversationList[conversation.ConversationId].Msg = &im.ChatLog{
				Id:             conversation.LastMsg.ID.Hex(),
				ConversationId: conversation.LastMsg.ConversationId,
				SendId:         conversation.LastMsg.SendId,
				RecvId:         conversation.LastMsg.RecvId,
				MsgType:        int32(conversation.LastMsg.MsgType),
				MsgContent:     conversation.LastMsg.MsgContent,
				ChatType:       int32(conversation.LastMsg.ChatType),
				SendTime:       conversation.LastMsg.SendTime,
			}
		}
	}
	return &res, nil
}
