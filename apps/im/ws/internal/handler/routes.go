package handler

import (
	"github.com/ljp-lachouchou/chan_xin/apps/im/ws/internal/handler/conversation"
	"github.com/ljp-lachouchou/chan_xin/apps/im/ws/internal/handler/user"
	"github.com/ljp-lachouchou/chan_xin/apps/im/ws/internal/svc"
	websocket2 "github.com/ljp-lachouchou/chan_xin/apps/im/ws/websocket"
)

func RegisterHandler(s *websocket2.Server, svc *svc.ServiceContext) {

	s.AddRoutes([]websocket2.Route{
		{
			Method:  "conversation.push",
			Handler: conversation.Push(svc),
		},
		{
			Method:  "conversation.chat",
			Handler: conversation.Chat(svc),
		},
		{
			Method:  "user.online",
			Handler: user.Online(svc),
		},
	})
}
