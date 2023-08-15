package main

import (
	"Douyin_Demo/config"
	user "Douyin_Demo/kitex_gen/douyin/user/userservice"
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

	addr, err := net.ResolveTCPAddr("tcp", config.UserServicePort)
	if err != nil {
		log.Fatal(err)
	}

	svr := user.NewServer(new(UserServiceImpl),
		server.WithRegistry(r),
		server.WithServiceAddr(addr),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: config.UserServiceName,
		}))

	err = svr.Run()

	if err != nil {
		log.Fatal(err.Error())
	}
}
