package conversation

import (
	"github.com/ljp-lachouchou/chan_xin/apps/im/ws/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/im/ws/websocket"
	"github.com/ljp-lachouchou/chan_xin/apps/im/ws/ws"
	"github.com/ljp-lachouchou/chan_xin/apps/task/mq/mq"
	"github.com/ljp-lachouchou/chan_xin/deploy/constant"
	"github.com/ljp-lachouchou/chan_xin/pkg/wuid"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"time"
)

func Push(svc *svc.ServiceContext) websocket.HandlerFunc {
	return func(server *websocket.Server, conn *websocket.Connection, msg *websocket.Message) {
		var data ws.Push
		if err := mapstructure.Decode(msg.Data, &data); err != nil {
			server.Send(websocket.NewErrMessage(err), conn)
			return
		}
		rconn := server.GetConn(data.RecvId)
		if rconn == nil {
			server.Send(websocket.NewErrMessage(errors.New("无接受者")), conn)
			return
		}
		server.Info("push msg :%v", data)
		server.Send(websocket.NewMessageWithId(msg.Id, data.SendId, msg.ToId, &ws.Chat{
			ConversationId: data.ConversationId,
			ChatType:       data.ChatType,
			SendId:         data.SendId,
			RecvId:         data.RecvId,
			Msg: ws.Msg{
				MsgType:    data.MsgType,
				MsgContent: data.MsgContent,
			},
			SendTime: data.SendTime,
		}), rconn)
	}
}
func Chat(svc *svc.ServiceContext) websocket.HandlerFunc {
	return func(server *websocket.Server, conn *websocket.Connection, msg *websocket.Message) {
		var data ws.Chat
		if err := mapstructure.Decode(msg.Data, &data); err != nil {
			server.Send(websocket.NewErrMessage(err), conn)
			return
		}
		switch data.ChatType {
		case constant.SingleChat:
			if data.ConversationId == "" {
				data.ConversationId = wuid.CombineId(data.SendId, data.RecvId)
			}
			if data.SendTime == 0 {
				data.SendTime = time.Now().Unix()
			}

			err := svc.Client.Push(&mq.MsgChatTransfer{
				MsgId:          msg.Id,
				ConversationId: data.ConversationId,
				SendId:         data.SendId,
				RecvId:         data.RecvId,
				ChatType:       data.ChatType,
				MsgType:        data.Msg.MsgType,
				MsgContent:     data.Msg.MsgContent,
				SendTime:       data.SendTime,
			})
			if err != nil {
				server.Send(websocket.NewErrMessage(err), conn)
				return
			}

		case constant.GroupChat:

		}
	}
}
