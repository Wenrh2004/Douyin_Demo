package main

import (
	"Douyin_Demo/config"
	publish "Douyin_Demo/kitex_gen/douyin/publish/publishservice"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	consul "github.com/kitex-contrib/registry-consul"
	"log"
	"net"
)

func main() {
	r, err := consul.NewConsulRegister(config.AppConfig.CONSUL_ADDRESS)
	if err != nil {
		log.Fatal(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", config.PublishServicePort)
	if err != nil {
		log.Fatal(err)
	}

	svr := publish.NewServer(new(PublishServiceImpl),
		server.WithRegistry(r),
		server.WithServiceAddr(addr),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: config.PublishServiceName,
		}))

	err = svr.Run()

	if err != nil {
		log.Fatal(err.Error())
	}
}
