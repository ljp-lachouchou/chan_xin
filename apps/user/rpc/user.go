package main

import (
	"flag"
	"fmt"
	"github.com/ljp-lachouchou/chan_xin/apps/user/rpc/internal/config"
	"github.com/ljp-lachouchou/chan_xin/apps/user/rpc/internal/server"
	"github.com/ljp-lachouchou/chan_xin/apps/user/rpc/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/user/rpc/user"
	"github.com/ljp-lachouchou/chan_xin/deploy/configserver"
	"github.com/ljp-lachouchou/chan_xin/pkg/interceptor/rpcserver"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "E:\\go_proj\\fvx_proj\\chan_xin\\apps\\user\\rpc\\etc\\dev\\user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	//conf.MustLoad(*configFile, &c)
	err := configserver.NewConfigServer(*configFile, configserver.NewSail(&configserver.Config{
		ETCDEndpoints:  "114.215.194.88:3379",
		ProjectKey:     "98c6f2c2287f4c73cea3d40ae7ec3ff2",
		Namespace:      "user",
		Configs:        "user-rpc.yaml",
		ConfigFilePath: "",
		LogLevel:       "DEBUG",
	})).MustLoad(&c, func(bytes []byte) error {
		fmt.Println("更新配置", string(bytes))
		return nil
	})
	if err != nil {
		panic(err)
	}
	ctx := svc.NewServiceContext(c)
	if err := ctx.SetRootToken(); err != nil {
		panic(err)
	}
	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserServiceServer(grpcServer, server.NewUserServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	s.AddUnaryInterceptors(rpcserver.LogInterceptor)
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
