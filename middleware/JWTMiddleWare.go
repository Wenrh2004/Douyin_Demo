package middleware

import (
	"Douyin_Demo/DTO"
	"Douyin_Demo/manager"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTClaims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

// AUDIENCE define a constant used by jwt-token.generate token audience
const AUDIENCE = "USER_NORMAL"

// GenerateJWTClaim generate json web token claim used by check login status
func GenerateJWTClaim(user DTO.UserRequestDTO) string {
	var tokenExpr = time.Now().Add(time.Duration(manager.GetTokenExpireConfig()) * time.Second)
	var claims = JWTClaims{
		UserID: user.UserID,
		StandardClaims: jwt.StandardClaims{
			Id:        user.Password,
			ExpiresAt: tokenExpr.Unix(),
			Issuer:    user.Username,
			Audience:  AUDIENCE,
		},
	}
	var token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(manager.GetTokenSecretConfig()))
	if err != nil {
		panic("current Error:" + err.Error())
	}
	return token
}
