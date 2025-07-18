package main

import (
	"flag"
	"fmt"
	"github.com/ljp-lachouchou/chan_xin/deploy/configserver"
	"github.com/ljp-lachouchou/chan_xin/pkg/lresoult"
	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/api/internal/config"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/api/internal/handler"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/dev/dynamics.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	err := configserver.NewConfigServer(*configFile, configserver.NewSail(&configserver.Config{
		ETCDEndpoints:  "114.215.194.88:3379",
		ProjectKey:     "98c6f2c2287f4c73cea3d40ae7ec3ff2",
		Namespace:      "dynamics",
		Configs:        "dynamics-api.yaml",
		ConfigFilePath: "",
		LogLevel:       "DEBUG",
	})).MustLoad(&c, func(bytes []byte) error {
		fmt.Println("更新配置", string(bytes))
		return nil
	})
	if err != nil {
		panic(err)
	}
	rest.WithCors()
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	httpx.SetOkHandler(lresoult.OkHandler)
	httpx.SetErrorHandler(lresoult.ErrorHandler(c.Name))
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
