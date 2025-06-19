package main

import (
	"flag"
	"fmt"
	"github.com/ljp-lachouchou/chan_xin/apps/im/ws/internal/config"
	"github.com/ljp-lachouchou/chan_xin/apps/im/ws/internal/handler"
	"github.com/ljp-lachouchou/chan_xin/apps/im/ws/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/im/ws/websocket"
	"github.com/ljp-lachouchou/chan_xin/deploy/configserver"
)

var configFile = flag.String("f", "etc/dev/im.yaml", "the config file")

func main() {
	flag.Parse()
	var c config.Config
	err := configserver.NewConfigServer(*configFile, configserver.NewSail(&configserver.Config{
		ETCDEndpoints:  "192.168.142.101:3379",
		ProjectKey:     "98c6f2c2287f4c73cea3d40ae7ec3ff2",
		Namespace:      "im",
		Configs:        "im-ws.yaml",
		ConfigFilePath: "",
		LogLevel:       "DEBUG",
	})).MustLoad(&c, func(bytes []byte) error {
		fmt.Println("更新配置", string(bytes))
		return nil
	})
	if err != nil {
		panic(err)
	}

	if err := c.SetUp(); err != nil {
		panic(err)
	}
	ctx := svc.NewServiceContext(c)
	srv := websocket.NewServer(c.ListenOn, websocket.WithAuthentication(handler.NewAuth(ctx)))

	//websocket.WithServerMaxConnectionIdle(10*time.Second)

	defer srv.Stop()

	handler.RegisterHandler(srv, ctx)
	fmt.Println("start websocket server at", c.ListenOn, "....")
	srv.Start()
}
