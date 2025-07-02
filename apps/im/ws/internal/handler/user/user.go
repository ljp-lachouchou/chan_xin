package user

import (
	"github.com/ljp-lachouchou/chan_xin/apps/im/ws/internal/svc"
	websocket2 "github.com/ljp-lachouchou/chan_xin/apps/im/ws/websocket"
)

func Online(svc *svc.ServiceContext) websocket2.HandlerFunc {
	return func(server *websocket2.Server, conn *websocket2.Connection, msg *websocket2.Message) {
		server.SendPingMessage()
	}
}
