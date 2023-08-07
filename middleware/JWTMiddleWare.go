package middleware

import (
	"Douyin_Demo/constants"
	"Douyin_Demo/dto"
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
func GenerateJWTClaim(user dto.UserRequestDTO) string {
	var tokenExpr = time.Now().Add(time.Duration(manager.GetYamlConfigByInt("token.expire")) * time.Second)
	var claims = JWTClaims{
		UserID:     user.UserID,
		GrantScope: constants.User_NORMAL,
		StandardClaims: jwt.StandardClaims{
			Audience:  AUDIENCE,
			ExpiresAt: tokenExpr.Unix(),
			Id:        user.Password,
			Issuer:    constants.Issuer,
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

// VerifyBasicTokenClaim verify token fmt expr
func VerifyBasicTokenClaim(token string) *jwt.Token {
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

// GetTokenSingleClaim get token single prof value
func GetTokenSingleClaim(token *jwt.Token, content string) string {
	var tokenProfValue string
	if claims, err := token.Claims.(jwt.MapClaims); err {
		tokenProfValue = claims[content].(string) // 获取 user_id 字段的值
	} else {
		panic(err)
	}
	return tokenProfValue
}

// TokenClaimVerify token claims verify
func TokenClaimVerify(token *jwt.Token, content string, res string) map[string]interface{} {
	if token == nil {
		return map[string]interface{}{
			"code":        422,
			"message":     "token is null",
			"description": constants.PARAMS_ERROR,
		}
	}
	// get single token claim
	claim := GetTokenSingleClaim(token, content)
	if claim == res && res != "" && claim != "" {
		return map[string]interface{}{
			"code":        422,
			"message":     "verify success",
			"description": constants.SUCCESS,
		}
	}
	return map[string]interface{}{
		"code":        422,
		"message":     "token is invalid",
		"description": constants.PARAMS_ERROR,
	}
}
