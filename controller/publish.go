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

func PublishActionController(ctx *gin.Context) {
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

	resp, _ := publishServiceClient.DouyinPublishAction(ctx, &req)

	//	返回结果
	ctx.JSON(http.StatusOK, resp)
	return
}

func PublishListController(ctx *gin.Context) {
	// TODO update data binding
	// get parameter
	var listParam PublishListParam
	err := ctx.ShouldBind(&listParam)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status_code": constants.PARAMS_ERROR_CODE,
			"status_msg":  err.Error(),
		})
		return
	}

	// params validation
	// TODO: validate token
	if listParam.UserId == 0 {
		ctx.JSON(200, gin.H{
			"status_code": constants.PARAMS_ERROR_CODE,
			"status_msg":  "user id is required",
		})
		return
	}

	resp, _ := publishServiceClient.PublishList(ctx, &publish.PublishListRequest{
		UserId: listParam.UserId,
		Token:  listParam.Token,
	})

	//	返回结果
	ctx.JSON(http.StatusOK, resp)
	return

}
