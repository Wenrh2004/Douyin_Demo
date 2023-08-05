package middleware

import (
	"Douyin_Demo/Constants"
	"Douyin_Demo/DTO"
	"Douyin_Demo/manager"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTClaims struct {
	UserID     string `json:"user_id"`
	GrantScope string
	jwt.StandardClaims
}

// AUDIENCE define a constant used by jwt-token.generate token audience
const AUDIENCE = "USER_NORMAL"

// GenerateJWTClaim generate json web token claim used by check login status
func GenerateJWTClaim(user DTO.UserRequestDTO) string {
	var tokenExpr = time.Now().Add(time.Duration(manager.GetYamlConfigByInt("token.expire")) * time.Second)
	var claims = JWTClaims{
		UserID:     user.UserID,
		GrantScope: Constants.User_NORMAL,
		StandardClaims: jwt.StandardClaims{
			Audience:  AUDIENCE,
			ExpiresAt: tokenExpr.Unix(),
			Id:        user.Password,
			Issuer:    Constants.Issuer,
		},
	}
	var token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(manager.GetYamlConfigByString("token.mySecret")))
	if err != nil {
		panic("current Error:" + err.Error())
	}
	return token
}
func Secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte(manager.GetYamlConfigByString("token.mySecret")), nil // 这是我的secret
	}
}

// ParseTokenClaim ParseTokenClaim
func ParseTokenClaim(token string) *jwt.Token {
	tokenVal, err := jwt.ParseWithClaims(token, &JWTClaims{}, Secret())
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				panic(errors.New("that's not even a token"))
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				panic(errors.New("token is expired"))
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				panic(errors.New("token not active yet"))
			} else {
				panic(errors.New("couldn't handle this token"))
			}
		}
	}
	return tokenVal
}

// VerifyToken get token single prof value
func VerifyToken(token *jwt.Token, content string) interface{} {
	var tokenProfValue interface{}
	if claims, err := token.Claims.(jwt.MapClaims); err {
		tokenProfValue = claims[content].(string) // 获取 user_id 字段的值
	} else {
		panic(err)
	}
	return tokenProfValue
}
