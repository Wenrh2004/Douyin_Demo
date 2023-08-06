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
	"Douyin_Demo/common"
	"Douyin_Demo/constants"
	"Douyin_Demo/model"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Register method for user registry
func Register(ctx *gin.Context) {
	db := common.GetDB()

	//	observe web request
	var requestUser model.User
	err := ctx.Bind(&requestUser)
	if err != nil {
		panic("Parameter bind failed" + err.Error())
	}
	userName := requestUser.Username
	password := requestUser.Password
	// todo 数据解密

	//	data check

	//	check username
	if len(userName) == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":        422,
			"message":     "用户名不能为空",
			"description": constants.PARAMS_ERROR,
		})
		var rgx = "^[a-zA-Z\\u4e00-\\u9fa5]{1,8}\\$" // 1-8 中文英文但是不包含下划线等符号
		matchedRes, _ := regexp.MatchString(rgx, userName)

		if !matchedRes {
			ctx.JSON(http.StatusUnavailableForLegalReasons, gin.H{
				"code":        422,
				"message":     "用户名不符合规范,",
				"description": constants.PARAMS_ERROR,
			})
		}
	}
	//	password check
	if len(password) < 6 || len(password) > 18 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":        422,
			"message":     "密码长度在6位-18位之间,且必须使用字母数字和特殊符号",
			"description": constants.MISMATCH,
		})
		var rgx = "^[a-zA-Z0-9~!@#\\$%^&*()_+}{\":?><,.';\\]\\[\\\\\\/\\-]{6,18}\\$" // 6 - 18 英语字母数字特殊符号组成

		matchedRes, _ := regexp.MatchString(rgx, userName)
		if !matchedRes {
			ctx.JSON(http.StatusUnavailableForLegalReasons, gin.H{
				"code":        422,
				"message":     "密码不符合规范,",
				"description": constants.PARAMS_ERROR,
			})
		}
	}
	// promise only one in db
	var user model.User
	db.Where("userName = ?", userName).First(&user)
	if user.ID != 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":        422,
			"message":     "用户已被注册",
			"description": constants.USER_PROFILE_ALREAD_UESD,
		})
	}

	//	save user into database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":        500,
			"message":     "用户信息异常",
			"description": constants.DB_MISMATCH,
		})
	}
	newUser := model.User{
		Username: userName,
		Password: string(hashedPassword),
	}
	tx := db.Create(&newUser)

	if tx != nil {
		ctx.JSON(http.StatusExpectationFailed, gin.H{
			"code":        500,
			"message":     "新增用户信息失败",
			"description": constants.DB_SAVE_FAILED,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":        200,
		"message":     "用户信息录入成功",
		"description": constants.SUCCESS,
	})
}

// Login user login
func Login(ctx *gin.Context) {
	db := common.GetDB()

	var requestUser model.User
	err := ctx.Bind(&requestUser)
	if err != nil {
		panic("Parameter bind failed" + err.Error())
	}
	//	get params from web
	userName := requestUser.Username
	password := requestUser.Password

	//	data check

	//	username check
	if len(userName) == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":        422,
			"message":     "用户名不能为空",
			"description": constants.PARAMS_ERROR,
		})
	}
	//	user in db status
	var user model.User
	db.Where("userName = ?", userName).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":        422,
			"message":     "用户不存在",
			"description": constants.MISMATCH,
		})
		return
	}
	//	password check
	if len(password) < 6 || len(password) > 18 {
		var rgx = "^[a-zA-Z0-9~!@#\\$%^&*()_+}{\":?><,.';\\]\\[\\\\\\/\\-]{6,18}\\$" // 6 - 18 英语字母数字特殊符号组成

		matchedRes, _ := regexp.MatchString(rgx, userName)

		if !matchedRes {
			ctx.JSON(http.StatusUnavailableForLegalReasons, gin.H{
				"code":        422,
				"message":     "密码不符合规范,",
				"description": constants.PARAMS_ERROR,
			})
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
			ctx.JSON(http.StatusExpectationFailed, gin.H{
				"code":        422,
				"message":     "密码错误",
				"description": constants.PARAMS_ERROR,
			})
		}
	}
	// todo write a token and refresh token
	ctx.JSON(http.StatusOK, gin.H{
		"code":        200,
		"message":     "登陆成功",
		"description": constants.LOGIN_SUCCESS,
	})
}
