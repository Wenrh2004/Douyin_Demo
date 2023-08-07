package sso

import (
	"Douyin_Demo/constants"
	"Douyin_Demo/middleware"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

const stringNil = ""

// SSOPermissionHandler SSOHandler Sso identity verification
func SSOPermissionHandler(needsVerifyTokenContentType string, needsVerifyContent ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get current token
		BusinessToken := ctx.GetHeader("x-authentication-business")
		if BusinessToken == stringNil {
			ctx.JSON(http.StatusNetworkAuthenticationRequired, gin.H{
				"code":        420,
				"message":     "token is null",
				"description": constants.PARAMS_ERROR,
			})
		}
		// verifyBasic info token including fmt expr time marshal
		BusTokenClaims := middleware.VerifyBasicTokenClaim(BusinessToken)

		for _, claim := range needsVerifyContent {
			// get token claim
			BusTokenRes := middleware.GetTokenSingleClaim(BusTokenClaims, needsVerifyTokenContentType)
			// permission check
			PermissionVerify(BusTokenClaims, claim, BusTokenRes)
		}
	}
}

func SecurityHandler(ctx *gin.Context) {
	// get current ip
	realIp := GetCurrentIP(ctx)
	// check ip exists in black list
	getIllegalIp(realIp)
	CsrfRefererCheck(ctx)
}

// PermissionVerify user permission checker
func PermissionVerify(token *jwt.Token, content string, res string) map[string]interface{} {
	if token == nil {
		return map[string]interface{}{
			"code":        422,
			"message":     "token is null",
			"description": constants.PARAMS_ERROR,
		}
	}
	// get single token claim
	claim := middleware.GetTokenSingleClaim(token, content)
	if claim == res && res != stringNil && claim != stringNil {
		return map[string]interface{}{
			"code":        422,
			"message":     "verify success",
			"description": constants.SUCCESS,
		}
	}
	return map[string]interface{}{
		"code":        422,
		"message":     "token is invalid",
		"description": constants.PRIVILEGE_NOT_ALLOWED,
	}
}

func GetCurrentIP(ctx *gin.Context) string {
	ip := ctx.Request.Header.Get("X-Real-IP")
	if ip == "" {
		ip = ctx.Request.Header.Get("X-Forwarded-For")
	}
	if ip == "" {
		ip = ctx.Request.RemoteAddr
	}
	return ip
}

func CsrfRefererCheck(ctx *gin.Context) {
	referer := ctx.Request.Referer()

	if referer == "" {
		panic("CsrfRefererCheck  nil error == >")
	}

	// 检查referer是否为合法的URL格式
	_, err := url.ParseRequestURI(referer)
	if err != nil {
		panic("CsrfRefererCheck fmt error == >")
	}
}

func StoreIllegalIp(ipList ...string) {
	var arr = [10]string{}
	for i, res := range ipList {
		arr[i] = res
	}
}

func getIllegalIp(ipList string) {
}
