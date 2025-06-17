package main

import (
	"flag"
	"fmt"
	"github.com/ljp-lachouchou/chan_xin/deploy/configserver"
	"github.com/ljp-lachouchou/chan_xin/pkg/redislock"
	"github.com/ljp-lachouchou/chan_xin/test/redislocktest/config"
	"github.com/ljp-lachouchou/chan_xin/test/redislocktest/svc"
	"log"
	"time"
)

var configFile = flag.String("f", "etc/dev/social.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	//conf.MustLoad(*configFile, &c)
	err := configserver.NewConfigServer(*configFile, configserver.NewSail(&configserver.Config{
		ETCDEndpoints:  "192.168.142.101:3379",
		ProjectKey:     "98c6f2c2287f4c73cea3d40ae7ec3ff2",
		Namespace:      "social",
		Configs:        "social-rpc.yaml",
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
	lock := redislock.NewRedisLock(ctx.Redis, "test", "20", time.Second*20)
	acquired, err := lock.Acquire()
	if err != nil {
		log.Fatalf("获取锁失败: %v", err)
	}
	if acquired {
		defer lock.Release()
		fmt.Println("获取锁成功，执行临界区代码")
		// 执行需要互斥的操作
		time.Sleep(5 * time.Second)
		fmt.Println("操作完成，释放锁")
	} else {
		fmt.Println("获取锁失败，资源正在被其他进程使用")
	}
}
