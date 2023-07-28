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
	"Douyin_Demo/common"
	"Douyin_Demo/controller"
	"github.com/gin-gonic/gin"
)

func collectRoutes(route *gin.Engine) *gin.Engine {

	//	用户注册
	route.POST("/register", controller.Register)
	//	用户登陆
	route.POST("/login", controller.Login)

	return route
}

func main() {
	//	获取初始化数据库
	db := common.InitDB()

	defer db.Close()

	//	创建路由
	route := gin.Default()
	route.ForwardedByClientIP = true
	route.SetTrustedProxies([]string{"127.0.0.1"})

	//	启动路由
	collectRoutes(route)

	//	启动服务
	route.Run(":5500")
}
