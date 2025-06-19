package client

import (
	"context"
	"encoding/json"
	"github.com/ljp-lachouchou/chan_xin/apps/task/mq/mq"
	"github.com/zeromicro/go-queue/kq"
)

type Client interface {
	Push(msg *mq.MsgChatTransfer) error
}
type msgChatTransferClient struct {
	pusher *kq.Pusher
}

func NewMsgChatTransferClient(addr []string, topic string, opts ...kq.PushOption) Client {
	return &msgChatTransferClient{
		pusher: kq.NewPusher(addr, topic, opts...),
	}
}
func (c *msgChatTransferClient) Push(msg *mq.MsgChatTransfer) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return c.pusher.Push(context.Background(), string(data))
}
