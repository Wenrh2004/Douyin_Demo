package controller

import (
	"Douyin_Demo/config"
	"Douyin_Demo/constants"
	"Douyin_Demo/kitex_gen/douyin/feed"
	"Douyin_Demo/kitex_gen/douyin/feed/feedservice"
	"github.com/cloudwego/kitex/client"
	"github.com/gin-gonic/gin"
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

func FeedAction(ctx *gin.Context) {
	//	获取参数
	var req feed.FeedRequest
	err := ctx.Bind(&req)
	if err != nil {
		ctx.JSON(200, gin.H{
			"status_code": constants.STATUS_FAILED,
			"status_msg":  err.Error(),
		})
		return
	}

	//	调用服务
	resp, err := feedServiceClient.GetVideoFeed(ctx, &req)
	if err != nil {
		ctx.JSON(200, gin.H{
			"status_code": constants.STATUS_FAILED,
			"status_msg":  err.Error(),
		})
		return
	}

	//	返回结果
	ctx.JSON(200, resp)

	return
}