package websocket

import (
	"time"
)

const (
	defaultMaxConnIdle = time.Hour * 9000
	defaultAckTimeOut  = time.Second * 10
)
