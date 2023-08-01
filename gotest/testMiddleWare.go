package gotest

import (
	"Douyin_Demo/DTO"
	"Douyin_Demo/middleware"
	"fmt"
)

func Test() {
	claim := middleware.GenerateJWTClaim(DTO.UserRequestDTO{Username: "username1", Password: "password2"})
	fmt.Println(claim)
}
