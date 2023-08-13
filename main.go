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

// collectRoutes ==> The function then sets up routes for handling user registration and login requests.
// There are two POST routers in this function，which are register and login.
// The register and login endpoint which calls the Register and Login functions from the controller.
func collectRoutes(route *gin.Engine) *gin.Engine {

	// user
	// User Login
	route.POST("/douyin/user/register", controller.Register)
	// User Login
	route.POST("/douyin/user/login", controller.Login)

	// message
	// Get message
	route.GET("/douyin/message/chat", controller.GetMessage)

	return route
}

func main() {

	// Get the initialization database
	common.InitDB()
	common.InitMongoDB()

	// Create routes
	route := gin.Default()
	route.ForwardedByClientIP = true
	proxyErr := route.SetTrustedProxies([]string{"127.0.0.1"})
	if proxyErr != nil {
		panic(proxyErr)
	}

	// Start routing
	collectRoutes(route)

	publishService := route.Group("/publish")
	publishService.POST("/action", controller.PublishAction)


	feedService := route.Group("/feed")
	feedService.GET("/", controller.FeedAction)

	// Run service on 5500
	err := route.Run(":5500")
	if err != nil {
		panic("service start failed" + err.Error())

	}
}
