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
	"Douyin_Demo/Constants"
	"Douyin_Demo/common"
	"Douyin_Demo/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// 用户注册
func Register(ctx *gin.Context) {
	db := common.GetDB()

	//	获取参数
	var requestUser model.User
	err := ctx.Bind(&requestUser)
	if err != nil {
		panic("Parameter bind failed" + err.Error())
	}
	userName := requestUser.Username
	password := requestUser.Password

	//	数据校验
	//	用户名校验
	if len(userName) == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":        422,
			"message":     "用户名不能为空",
			"description": Constants.PARAMS_ERROR,
		})
		return
	}
	//	密码校验
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":        422,
			"message":     "密码至少6位",
			"description": Constants.MISMATCH,
		})
		return
	}
	//	用户是否已注册
	var user model.User
	db.Where("userName = ?", userName).First(&user)
	if user.ID != 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":        422,
			"message":     "用户已被注册",
			"description": Constants.USER_PROFILE_ALREAD_UESD,
		})
		return
	}

	//	新增用户
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":        500,
			"message":     "新增用户信息失败",
			"description": Constants.DB_SAVE_FAILED,
		})
		return
	}
	newUser := model.User{
		Username: userName,
		Password: string(hasedPassword),
	}
	db.Create(&newUser)

	ctx.JSON(http.StatusOK, gin.H{
		"code":        200,
		"message":     "用户信息录入成功",
		"description": Constants.SUCCESS,
	})
}

// 用户登陆
func Login(ctx *gin.Context) {
	db := common.GetDB()

	var requestUser model.User
	err := ctx.Bind(&requestUser)
	if err != nil {
		panic("Parameter bind failed" + err.Error())
	}
	//	获取参数
	userName := requestUser.Username
	password := requestUser.Password

	//	数据校验
	//	用户名校验
	if len(userName) == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":        422,
			"message":     "用户名不能为空",
			"description": Constants.PARAMS_ERROR,
		})
		return
	}
	//	密码校验
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":        422,
			"message":     "密码至少6位",
			"description": Constants.PARAMS_ERROR,
		})
		return
	}
	//	用户是否注册
	var user model.User
	db.Where("userName = ?", userName).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":        422,
			"message":     "用户不存在",
			"description": Constants.MISMATCH,
		})
		return
	}
	//	密码校验
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":        422,
			"message":     "密码错误",
			"description": Constants.PARAMS_ERROR,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":        200,
		"message":     "登陆成功",
		"description": Constants.LOGIN_SUCCESS,
	})
}
