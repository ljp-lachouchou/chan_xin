package dynamicsmodels

import "github.com/zeromicro/go-zero/core/stores/mon"

var _ PostsModel = (*customPostsModel)(nil)

type (
	// PostsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPostsModel.
	PostsModel interface {
		postsModel
	}

	customPostsModel struct {
		*defaultPostsModel
	}
)

// NewPostsModel returns a model for the mongo.
func NewPostsModel(url, db, collection string) PostsModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customPostsModel{
		defaultPostsModel: newDefaultPostsModel(conn),
	}
}
func MustNewPostsModel(url, db string) PostsModel {
	conn := mon.MustNewModel(url, db, "posts")
	return &customPostsModel{
		defaultPostsModel: newDefaultPostsModel(conn),
	}
}
