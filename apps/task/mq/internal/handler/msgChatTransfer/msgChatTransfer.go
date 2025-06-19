package msgChatTransfer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ljp-lachouchou/chan_xin/apps/im/immodels"
	"github.com/ljp-lachouchou/chan_xin/apps/im/ws/websocket"
	"github.com/ljp-lachouchou/chan_xin/apps/task/mq/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/task/mq/mq"
	"github.com/ljp-lachouchou/chan_xin/deploy/constant"
	"github.com/zeromicro/go-zero/core/logx"
)

type MsgChatTransfer struct {
	svc *svc.ServiceContext
	logx.Logger
}

func NewMsgChatTransfer(svc *svc.ServiceContext) *MsgChatTransfer {
	return &MsgChatTransfer{
		svc:    svc,
		Logger: logx.WithContext(context.Background()),
	}
}
func (m *MsgChatTransfer) Consume(ctx context.Context, key, value string) error {
	fmt.Println("consume key:", key, "value:", value)
	var (
		data mq.MsgChatTransfer
	)
	if err := json.Unmarshal([]byte(value), &data); err != nil {
		return err
	}
	err := m.addChatLog(ctx, &data)
	if err != nil {
		return err
	}
	fmt.Println("ws client is", m.svc.Client)
	return m.svc.Client.Send(websocket.Message{
		FrameType: websocket.FrameData,
		FromId:    constant.SYSTEM_ROOT_ID,
		Method:    "conversation.push",
		Data:      data,
	})
}
func (m *MsgChatTransfer) addChatLog(ctx context.Context, data *mq.MsgChatTransfer) error {
	chatLog := &immodels.ChatLog{
		ConversationId: data.ConversationId,
		SendId:         data.SendId,
		RecvId:         data.RecvId,
		MsgType:        data.MsgType,
		MsgContent:     data.MsgContent,
		SendTime:       data.SendTime,
	}
	//加入到消息记录里面
	if err := m.svc.ChatLogModel.Insert(ctx, chatLog); err != nil {
		return err
	}
	return m.svc.ConversationModel.UpdateMsg(ctx, chatLog)
}
