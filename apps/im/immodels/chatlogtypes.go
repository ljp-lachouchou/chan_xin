package immodels

import (
	"github.com/ljp-lachouchou/chan_xin/apps/im/ws/ws"
	"github.com/ljp-lachouchou/chan_xin/deploy/constant"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChatLog struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ConversationId string             `bson:"conversationId"`
	SendId         string             `bson:"sendId"`
	RecvId         string             `bson:"recvId"`
	MsgType        constant.MType     `bson:"msgType"`
	MsgContent     ws.MessageContent  `bson:"msgContent"`
	SendTime       int64              `bson:"sendTime"`
	UpdateAt       time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt       time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
