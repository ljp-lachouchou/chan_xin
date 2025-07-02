package main

import (
	"flag"
	"fmt"
	"github.com/ljp-lachouchou/chan_xin/apps/task/mq/internal/config"
	"github.com/ljp-lachouchou/chan_xin/apps/task/mq/internal/handler"
	"github.com/ljp-lachouchou/chan_xin/apps/task/mq/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/deploy/configserver"
	"github.com/zeromicro/go-zero/core/service"
)

var configFile = flag.String("f", "etc/dev/task.yaml", "the config file")

func main() {
	flag.Parse()
	var c config.Config
	err := configserver.NewConfigServer(*configFile, configserver.NewSail(&configserver.Config{
		ETCDEndpoints:  "114.215.194.88:3379",
		ProjectKey:     "98c6f2c2287f4c73cea3d40ae7ec3ff2",
		Namespace:      "task",
		Configs:        "task-mq.yaml",
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
	serviceGroup := service.NewServiceGroup()
	listener := handler.NewListener(ctx)
	for _, s := range listener.Services() {
		serviceGroup.Add(s)
	}
	fmt.Println("starting mqueue at .....")
	serviceGroup.Start()
}
