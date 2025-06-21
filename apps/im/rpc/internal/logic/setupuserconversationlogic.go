package logic

import (
	"context"
	"fmt"
	"github.com/ljp-lachouchou/chan_xin/apps/im/immodels"
	"github.com/ljp-lachouchou/chan_xin/apps/im/rpc/im"
	"github.com/ljp-lachouchou/chan_xin/apps/im/rpc/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/deploy/constant"
	"github.com/ljp-lachouchou/chan_xin/pkg/lerr"
	"github.com/ljp-lachouchou/chan_xin/pkg/wuid"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SetUpUserConversationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetUpUserConversationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetUpUserConversationLogic {
	return &SetUpUserConversationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 建立会话: 群聊, 私聊
func (l *SetUpUserConversationLogic) SetUpUserConversation(in *im.SetUpUserConversationReq) (*im.SetUpUserConversationResp, error) {
	switch constant.ChatType(in.ChatType) {
	case constant.SingleChat:
		conversationId := wuid.CombineId(in.SendId, in.RecvId)
		conversation, err := l.svcCtx.ConversationModel.FindOneByConversationId(l.ctx, conversationId)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				data := &immodels.Conversation{
					ConversationId: conversationId,
					ChatType:       constant.SingleChat,
					IsShow:         true,
					Total:          0,
					Seq:            1,
				}
				err := l.svcCtx.ConversationModel.Insert(l.ctx, data)
				if err != nil {
					return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "im-rpc SetUpUserConversation insert", conversationId)
				}
			} else {
				return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "im-rpc SetUpUserConversation FindOneByConversationId", conversationId)
			}
		} else if conversation != nil {
			return nil, nil
		}
		err = l.setUpUserConversation(conversationId, in.SendId, in.RecvId, constant.SingleChat, true)
		if err != nil {
			return nil, err
		}
		err = l.setUpUserConversation(conversationId, in.RecvId, in.SendId, constant.SingleChat, false)
		if err != nil {
			return nil, err
		}
	case constant.GroupChat:
		conversationId := wuid.CombineId(in.SendId, in.RecvId)
		conversation, err := l.svcCtx.ConversationModel.FindOneByConversationId(l.ctx, conversationId)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				data := &immodels.Conversation{
					ConversationId: conversationId,
					ChatType:       constant.GroupChat,
					IsShow:         true,
					Total:          0,
					Seq:            1,
				}
				err := l.svcCtx.ConversationModel.Insert(l.ctx, data)
				if err != nil {
					return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "im-rpc SetUpUserConversation insert", conversationId)
				}
			} else {
				return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "im-rpc SetUpUserConversation FindOneByConversationId", conversationId)
			}

		} else if conversation != nil {
			return nil, nil
		}
		err = l.setUpUserConversation(conversationId, in.SendId, in.RecvId, constant.GroupChat, true)
		if err != nil {
			return nil, err
		}
	}
	return &im.SetUpUserConversationResp{}, nil
}
func (l *SetUpUserConversationLogic) setUpUserConversation(conversationId, userId,
	recvId string, chatType constant.ChatType, isShow bool) error {
	conversations, err := l.svcCtx.ConversationsModel.FindByUserId(l.ctx, userId)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			conversations = &immodels.Conversations{
				ID:               primitive.NewObjectID(),
				UserId:           userId,
				ConversationList: make(map[string]*immodels.Conversation),
			}
		} else {
			return errors.Wrapf(lerr.NEWDBError(), fmt.Sprintf("ConversationsModel.FindByUserId err %v req %v", err, userId))
		}
	}
	if _, ok := conversations.ConversationList[conversationId]; ok {
		return nil
	}
	//添加会话记录
	conversations.ConversationList[conversationId] = &immodels.Conversation{
		ConversationId: conversationId,
		ChatType:       chatType,
		IsShow:         isShow,
		Seq:            1,
	}
	//更新
	_, err = l.svcCtx.ConversationsModel.Update(l.ctx, conversations)
	if err != nil {
		return errors.Wrapf(lerr.NEWDBError(), fmt.Sprintf("ConversationsModel.Update err %v req %v", err, userId))
	}
	return nil
}
