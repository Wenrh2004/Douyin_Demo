package main

import (
	"Douyin_Demo/config"
	"Douyin_Demo/constants"
	"Douyin_Demo/kitex_gen/douyin/feed"
	"Douyin_Demo/kitex_gen/douyin/feed/feedservice"
	"Douyin_Demo/model"
	"context"
	"errors"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"log"
)

var feedServiceClient feedservice.Client

func init() {
	r, err := consul.NewConsulResolver(config.AppConfig.CONSUL_ADDRESS)
	if err != nil {
		log.Fatal(err)
	}

	feedServiceClient, err = feedservice.NewClient(config.FeedServiceName, client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
}

func getVideos(publishList []*model.Publish, token string) ([]*feed.Video, error) {
	var videos []*feed.Video

	for _, p := range publishList {
		//	调用服务
		resp, err := feedServiceClient.GetVideo(context.Background(), &feed.GetVideoRequest{
			VideoId: int64(p.ID),
			Token:   &token,
		})

		if err != nil {
			return nil, err
		}

		if resp.StatusCode != constants.STATUS_SUCCESS {
			// return resp.StatusMsg as error
			return nil, errors.New(*resp.StatusMsg)
		}

		videos = append(videos, resp.Video)
	}

	return videos, nil
}
