package immodels

import (
	"github.com/ljp-lachouchou/chan_xin/deploy/constant"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Conversation struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ConversationId string             `bson:"conversationId"`
	ChatType       constant.ChatType  `bson:"chatType"`
	SendId         string             `bson:"sendId"`
	RecvId         string             `bson:"recvId"`
	IsShow         bool               `bson:"isShow"`
	Total          int64              `bson:"total"`
	Seq            int64              `bson:"seq"` //在我的会话列表中的序号
	LastMsg        *ChatLog           `bson:"lastMsg"`
	UpdateAt       time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt       time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
