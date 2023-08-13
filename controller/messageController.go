package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMessage(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":        200,
		"message":     "聊天记录获取成功",
		"description": "SUCCESS",
	})
}
