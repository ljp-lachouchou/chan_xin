package immodels

import "github.com/zeromicro/go-zero/core/stores/mon"

var _ ConversationsModel = (*customConversationsModel)(nil)

type (
	// ConversationsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customConversationsModel.
	ConversationsModel interface {
		conversationsModel
	}

	customConversationsModel struct {
		*defaultConversationsModel
	}
)

// NewConversationsModel returns a model for the mongo.
func NewConversationsModel(url, db, collection string) ConversationsModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customConversationsModel{
		defaultConversationsModel: newDefaultConversationsModel(conn),
	}
}
func MustNewConversationsModel(url, db string) ConversationsModel {
	conn := mon.MustNewModel(url, db, "conversations")
	return &customConversationsModel{
		defaultConversationsModel: newDefaultConversationsModel(conn),
	}
}
