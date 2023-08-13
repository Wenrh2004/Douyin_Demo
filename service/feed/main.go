package main

import (
	"Douyin_Demo/config"
	feed "Douyin_Demo/kitex_gen/douyin/feed/feedservice"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"

	consul "github.com/kitex-contrib/registry-consul"
)

func main() {
	r, err := consul.NewConsulRegister(config.AppConfig.CONSUL_ADDRESS)
	if err != nil {
		log.Fatal(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", config.FeedServicePort)
	if err != nil {
		log.Fatal(err)
	}

	svr := feed.NewServer(new(FeedServiceImpl),
		server.WithRegistry(r),
		server.WithServiceAddr(addr),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: config.FeedServiceName,
		}))

	err = svr.Run()

	if err != nil {
		log.Fatal(err)
	}
}
