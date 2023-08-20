/*
 * Copyright (c) 2023.
 * Project: Douyin_Demo
 * File: main.go
 * Last date: 2023/7/28 下午6:59
 * Developer: KingYen
 *
 * Created by KingYen on 2023/7/28 18:59:3.
 */

package main

import (
	"Douyin_Demo/controller"
	"github.com/gin-gonic/gin"
)


func collectRoutes(route *gin.Engine) *gin.Engine {

	// user
	// User registration.
	route.POST("/register", controller.Register)
	// User login.
	route.POST("/login", controller.Login)

	// message
	// Get chat logs.
	route.GET("/message/chat", controller.GetMessage)
	return route
}

func main() {
	//	获取初始化数据库
	// common.InitDB()
	//	创建路由
	route := gin.Default()
	route.ForwardedByClientIP = true
	proxyErr := route.SetTrustedProxies([]string{"127.0.0.1"})
	if proxyErr != nil {
		panic(proxyErr)
	}

	douyin := route.Group("/douyin")

	//	启动路由
	publishService := douyin.Group("/publish")
	publishService.POST("/action", controller.PublishActionController)
	publishService.GET("/list", controller.PublishListController)

	feedService := douyin.Group("/feed")
	feedService.GET("/", controller.FeedAction)

	userService := douyin.Group("/user")
	userService.GET("", controller.GetUserProfileController)
	userService.POST("/register", controller.Register)
	userService.POST("/login", controller.Login)
  
  messageService := douyin.Group("/meesage")
  messageService.GET("/chat", controller.GetMessage)

	//	启动服务
	err := route.Run(":5500")
	if err != nil {
		panic("service start failed" + err.Error())

	}
}
