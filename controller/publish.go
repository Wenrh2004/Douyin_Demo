package controller

import (
	"Douyin_Demo/Constants"
	"Douyin_Demo/kitex_gen/douyin/publish/action"
	"Douyin_Demo/kitex_gen/douyin/publish/action/douyinpublishactionservice"
	"github.com/gin-gonic/gin"
	"net/http"
)

var publishServiceClient douyinpublishactionservice.Client

func PublishAction(ctx *gin.Context) {
	// get parameter
	token := "123456"
	_, err := publishServiceClient.DouyinPublishAction(ctx, &action.DouyinPublishActionRequest{
		Token: token,
	})

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status_code": Constants.STATUS_FAILED,
			"status_msg":  err.Error(),
		})
		return
	}

	//	返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": Constants.STATUS_SUCCESS,
	})

	return
}
