package dynamicsmodels

import (
	"errors"

	"github.com/zeromicro/go-zero/core/stores/mon"
)

var (
	MongoErrNotFound   = mon.ErrNotFound
	ErrInvalidObjectId = errors.New("invalid objectId")
)
