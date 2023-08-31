/*
 * Copyright (c) 2023.
 * Project: Douyin_Demo
 * File: userController.go
 * Last date: 2023/7/28 下午6:52
 * Developer: KingYen
 *
 * Created by KingYen on 2023/7/28 18:52:9.
 */

package controller

import (
	"Douyin_Demo/config"
	"Douyin_Demo/constants"
	"Douyin_Demo/kitex_gen/douyin/user"
	"Douyin_Demo/kitex_gen/douyin/user/userservice"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"log"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

var userServiceClient userservice.Client

func init() {
	r, err := consul.NewConsulResolver(config.AppConfig.CONSUL_ADDRESS)
	if err != nil {
		log.Fatal(err)
	}

	userServiceClient, err = userservice.NewClient(config.UserServiceName, client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
}

// Register method for user registry
func Register(ctx *gin.Context) {
	//	observe web request
	var registerParam UserRegisterParam
	err := ctx.Bind(&registerParam)
	if err != nil {
		ctx.JSON(200, gin.H{
			"status_code": constants.STATUS_FAILED,
			"status_msg":  err.Error(),
		})
		return
	}
	userName := registerParam.Username
	password := registerParam.Password

	//	data check
	//	null param check
	if len(userName) == 0 || len(password) == 0 {
		msg := constants.PARAMS_ERROR
		ctx.JSON(http.StatusOK, gin.H{
			"status_code": constants.PARAMS_ERROR_CODE,
			"status_msg":  &msg,
		})
		return
	}

	//	password check
	if len(password) < 6 || len(password) > 18 {
		msg := constants.INVALID_REGISTER_PWD
		ctx.JSON(http.StatusOK, gin.H{
			"status_code": constants.PARAMS_ERROR_CODE,
			"status_msg":  &msg,
		})
		return
	}

	if !ValidateEmail(userName) {
		msg := constants.INVALID_REGISTER_EMAIL
		ctx.JSON(http.StatusOK, gin.H{
			"status_code": constants.PARAMS_ERROR_CODE,
			"status_msg":  &msg,
		})
		return
	}

	resp, _ := userServiceClient.UserRegister(ctx, &user.UserRegisterRequest{
		Username: userName,
		Password: password,
	})

	ctx.JSON(200, resp)
	return
}

// Login user login
func Login(ctx *gin.Context) {
	var loginParam UserLoginParam
	err := ctx.Bind(&loginParam)
	if err != nil {
		ctx.JSON(200, gin.H{
			"status_code": constants.STATUS_FAILED,
			"status_msg":  err.Error(),
		})
		return
	}
	userName := loginParam.Username
	password := loginParam.Password

	//	data check
	//	null param check
	if len(userName) == 0 || len(password) == 0 {
		msg := constants.PARAMS_ERROR
		ctx.JSON(http.StatusOK, gin.H{
			"status_code": constants.PARAMS_ERROR_CODE,
			"status_msg":  &msg,
		})
		return
	}

	resp, err := userServiceClient.UserLogin(ctx, &user.UserLoginRequest{
		Username: userName,
		Password: password,
	})

	if err != nil {
		ctx.JSON(200, gin.H{
			"status_code": constants.STATUS_FAILED,
			"status_msg":  err.Error(),
		})
		return
	}

	ctx.JSON(200, resp)

}

// GetUserProfileController get user profile
func GetUserProfileController(ctx *gin.Context) {
	var param UserProfileParam
	err := ctx.Bind(&param)

	if err != nil {
		ctx.JSON(200, gin.H{
			"status_code": constants.STATUS_FAILED,
			"status_msg":  err.Error(),
		})
		return
	}

	// TODO Validate token
	if param.UserId == 0 {
		ctx.JSON(200, gin.H{
			"status_code": constants.PARAMS_ERROR_CODE,
			"status_msg":  "user id is required",
		})
		return
	}

	resp, _ := userServiceClient.GetUserInfo(ctx, &user.UserInfoRequest{
		UserId: param.UserId,
		Token:  param.Token,
	})

	ctx.JSON(200, resp)
	return
}

func ValidateEmail(username string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(username)
}
