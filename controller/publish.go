package controller

import (
	"Douyin_Demo/config"
	"Douyin_Demo/constants"
	"Douyin_Demo/kitex_gen/douyin/publish/action"
	"Douyin_Demo/kitex_gen/douyin/publish/action/douyinpublishactionservice"
	"github.com/cloudwego/kitex/client"
	"github.com/gin-gonic/gin"
	consul "github.com/kitex-contrib/registry-consul"
	"log"
	"net/http"
)

var publishServiceClient douyinpublishactionservice.Client

func init() {
	r, err := consul.NewConsulResolver(config.AppConfig.CONSUL_ADDRESS)
	if err != nil {
		log.Fatal(err)
	}

	publishServiceClient, err = douyinpublishactionservice.NewClient(config.PublishServiceName, client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
}

func PublishAction(ctx *gin.Context) {
	// get parameter
	token := "123456"
	_, err := publishServiceClient.DouyinPublishAction(ctx, &action.DouyinPublishActionRequest{
		Token: token,
	})

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status_code": constants.STATUS_FAILED,
			"status_msg":  err.Error(),
		})
		return
	}

	//	返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": constants.STATUS_SUCCESS,
	})

	return
}
