package rpcserver

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
	"time"
)

func LogInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	resp, err = handler(ctx, req)
	start := time.Now()
	logx.Infof("Service Method is %s,Duration is %v", info.FullMethod, time.Since(start))
	return resp, err
}
