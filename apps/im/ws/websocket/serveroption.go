package websocket

import "time"

type ServerOption func(opt *serverOption)

// 服务端配置
type serverOption struct {
	patten string
	Authentication
	ackTimeOut  time.Duration
	maxConnIdle time.Duration
}

func newServerOption(opt ...ServerOption) serverOption {
	o := serverOption{
		patten:         "/ws",
		Authentication: NewAuthentication(),
		ackTimeOut:     defaultAckTimeOut,
		maxConnIdle:    defaultMaxConnIdle,
	}
	for _, v := range opt {
		v(&o)
	}
	return o
}
func WithAuthentication(a Authentication) ServerOption {
	return func(opt *serverOption) {
		opt.Authentication = a
	}
}
func WithMaxConnIdle(d time.Duration) ServerOption {
	return func(opt *serverOption) {
		if d <= 0 {
			return
		}
		opt.maxConnIdle = d
	}
}
