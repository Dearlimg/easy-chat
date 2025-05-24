package main

import (
	"easy-chat/apps/im/ws/websocket"
	"flag"
	"fmt"

	"easy-chat/apps/im/ws/internal/config"
	"easy-chat/apps/im/ws/internal/svc"
	
)

var configFile = flag.String("f", "etc/dev/im.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	if err := c.SetUp(); err != nil {
		panic(err)
	}
	svc.NewServiceContext(c)
	srv := websocket.NewServer(c.ListenOn)
	defer srv.Stop()

	fmt.Println("start websocket server at ", c.ListenOn, " ..... ")
	srv.Start()

}
