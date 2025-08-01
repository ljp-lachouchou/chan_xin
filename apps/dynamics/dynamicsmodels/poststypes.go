package dynamicsmodels

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Posts struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserId       string             `bson:"userId"`
	Content      string             `bson:"content"`
	ImageUrls    []string           `bson:"imageUrls"`
	Emoji        string             `bson:"emoji"` // 表情符号（Unicode或自定义标识）
	Visibility   int                `bson:"visibility"`
	VisibleTo    []string           `bson:"visibleTo"`
	LikeCount    int                `bson:"likeCount"`    // 点赞数[7](@ref)
	CommentCount int                `bson:"commentCount"` // 评论数
	Location     string             `bson:"location"`     // 发布位置（如"北京·故宫"）[1](@ref)
	CreateTime   int64              `bson:"createTime"`
	//Tags         []string           `bson:"tags"`         // 话题标签（如#旅行）
	IsPinned bool      `bson:"isPinned"`
	UpdateAt time.Time `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
