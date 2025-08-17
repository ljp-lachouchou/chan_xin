package conversation

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/im/ws/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/im/ws/websocket"
	"github.com/ljp-lachouchou/chan_xin/apps/im/ws/ws"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/socialservice"
	"github.com/ljp-lachouchou/chan_xin/apps/task/mq/mq"
	"github.com/ljp-lachouchou/chan_xin/deploy/constant"
	"github.com/ljp-lachouchou/chan_xin/pkg/wuid"
	"github.com/mitchellh/mapstructure"
	"time"
)

func single(srv *websocket.Server, data *ws.Push, recvId string) error {
	rconn := srv.GetConn(recvId)
	if rconn == nil {
		// todo: 目标离线
		return nil
	}

	srv.Infof("push msg %v", data)

	return srv.Send(websocket.NewMessageWithId(data.Id, data.SendId, data.RecvId, &ws.Chat{
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
func group(svc *svc.ServiceContext, srv *websocket.Server, data *ws.Push) error {
	in := &socialservice.GetGroupMembersReq{
		GroupId: data.RecvId,
	}
	members, err := svc.SocialService.GetGroupMembers(context.Background(), in)
	if err != nil {
		return err
	}
	var ids []string
	for _, v := range members.List {
		if v.UserId == data.SendId {
			continue
		}
		ids = append(ids, v.UserId)
	}
	data.RecvIds = ids
	for _, id := range data.RecvIds {
		func(id string) {
			srv.Schedule(func() {
				single(srv, data, id)
			})
		}(id)
	}
	return nil
}
func Push(svc *svc.ServiceContext) websocket.HandlerFunc {
	return func(server *websocket.Server, conn *websocket.Connection, msg *websocket.Message) {
		var data ws.Push
		if err := mapstructure.Decode(msg.Data, &data); err != nil {
			server.Send(websocket.NewErrMessage(err))
			return
		}
		// 发送的目标
		switch data.ChatType {
		case constant.SingleChat:
			single(server, &data, data.RecvId)
		case constant.GroupChat:
			group(svc,server, &data)
		}

	}
}
func Chat(svc *svc.ServiceContext) websocket.HandlerFunc {
	return func(server *websocket.Server, conn *websocket.Connection, msg *websocket.Message) {
		var data ws.Chat
		if err := mapstructure.Decode(msg.Data, &data); err != nil {
			server.Send(websocket.NewErrMessage(err), conn)
			return
		}
		if data.ConversationId == "" {
			switch data.ChatType {
			case constant.SingleChat:
				data.ConversationId = wuid.CombineId(conn.Uid, data.RecvId)
			case constant.GroupChat:
				data.ConversationId = data.RecvId
			}
		}
		if data.SendTime == 0 {
			data.SendTime = time.Now().UnixMilli()
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
	}
}
