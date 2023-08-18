package main

import (
	"Douyin_Demo/config"
	"Douyin_Demo/constants"
	"Douyin_Demo/kitex_gen/douyin/user"
	"Douyin_Demo/kitex_gen/douyin/user/userservice"
	"context"
	"errors"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"log"
)

var userServiceClient userservice.Client

func init() {
	r, err := consul.NewConsulResolver(config.AppConfig.CONSUL_ADDRESS)
	if err != nil {
		log.Fatal(err)
	}

	userServiceClient, err = userservice.NewClient(config.UserServiceName, client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
}

func getUserById(userId int64, token string) (*user.User, error) {
	//	调用服务
	resp, err := userServiceClient.GetUserInfo(context.Background(), &user.UserInfoRequest{
		UserId: userId,
		Token:  token,
	})

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != constants.STATUS_SUCCESS {
		// return resp.StatusMsg as error
		return nil, errors.New(*resp.StatusMsg)
	}

	return resp.User, nil
}
