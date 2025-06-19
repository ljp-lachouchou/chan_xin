package handler

import (
	"github.com/ljp-lachouchou/chan_xin/apps/task/mq/internal/handler/msgChatTransfer"
	"github.com/ljp-lachouchou/chan_xin/apps/task/mq/internal/svc"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
)

type Listener struct {
	svc *svc.ServiceContext
}

func NewListener(svc *svc.ServiceContext) *Listener {
	return &Listener{
		svc: svc,
	}
}
func (l *Listener) Services() []service.Service {
	return []service.Service{
		kq.MustNewQueue(l.svc.Config.MsgChatTransfer, msgChatTransfer.NewMsgChatTransfer(l.svc)),
	}
}
