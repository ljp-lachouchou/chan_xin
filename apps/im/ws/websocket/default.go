package websocket

import (
	"syscall"
	"time"
)

const (
	defaultMaxConnIdle = time.Second * syscall.INFINITE
)
