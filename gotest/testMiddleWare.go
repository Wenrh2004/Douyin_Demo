package main

import (
	"Douyin_Demo/dto"
	"Douyin_Demo/middleware"
	"fmt"
)

func Test() {
	claim := middleware.GenerateJWTClaim(dto.UserRequestDTO{Username: "username1", Password: "password2"})
	fmt.Println(claim)
}

func main() {
	claim := middleware.GenerateJWTClaim(dto.UserRequestDTO{Username: "username1", Password: "password2"})
	tokenClaim := middleware.ParseTokenClaim(claim)
	fmt.Println(tokenClaim)
}
