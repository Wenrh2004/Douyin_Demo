package controller

import (
	"Douyin_Demo/config"
	"Douyin_Demo/constants"
	"Douyin_Demo/kitex_gen/douyin/publish"
	"Douyin_Demo/kitex_gen/douyin/publish/publishservice"
	"github.com/cloudwego/kitex/client"
	"github.com/gin-gonic/gin"
	consul "github.com/kitex-contrib/registry-consul"
	"log"
	"net/http"
)

var publishServiceClient publishservice.Client

func init() {
	r, err := consul.NewConsulResolver(config.AppConfig.CONSUL_ADDRESS)
	if err != nil {
		log.Fatal(err)
	}

	publishServiceClient, err = publishservice.NewClient(config.PublishServiceName, client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
}

func PublishAction(ctx *gin.Context) {
	// get parameter
	var req publish.DouyinPublishActionRequest
	err := ctx.Bind(&req)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status_code": constants.PARAMS_ERROR_CODE,
			"status_msg":  err.Error(),
		})
		return
	}

	resp, err := publishServiceClient.DouyinPublishAction(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusOK, resp)
		return
	}

	//	返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": constants.STATUS_SUCCESS,
		"status_msg":  "success",
	})

	return
}

func PublishListController(ctx *gin.Context) {
	// get parameter
	var req publish.PublishListRequest
	err := ctx.Bind(&req)

	if err != nil {
		// TODO: Log error

		ctx.JSON(http.StatusOK, gin.H{
			"status_code": constants.PARAMS_ERROR_CODE,
			"status_msg":  err.Error(),
		})
		return
	}

	resp, err := publishServiceClient.PublishList(ctx, &req)

	if err != nil {
		// TODO: Log error
		ctx.JSON(http.StatusOK, resp)
		return
	}

	//	返回结果
	ctx.JSON(http.StatusOK, resp)
	return

}
