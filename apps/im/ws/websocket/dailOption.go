package websocket

import "net/http"

type DailOption func(opt *dailOption)
type dailOption struct {
	patten string
	header http.Header
}

func newDailOptions(opts ...DailOption) dailOption {
	o := dailOption{
		patten: "/ws",
		header: nil,
	}
	for _, opt := range opts {
		opt(&o)
	}
	return o
}
func WithClientHeader(header http.Header) DailOption {
	return func(opt *dailOption) {
		opt.header = header
	}
}
