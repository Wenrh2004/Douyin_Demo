package main

import (
	"Douyin_Demo/DTO"
	"Douyin_Demo/middleware"
	"fmt"
)

func Test() {
	claim := middleware.GenerateJWTClaim(DTO.UserRequestDTO{Username: "username1", Password: "password2"})
	fmt.Println(claim)
}

func main() {
	claim := middleware.GenerateJWTClaim(DTO.UserRequestDTO{Username: "username1", Password: "password2"})
	tokenClaim := middleware.ParseTokenClaim(claim)
	fmt.Println(tokenClaim)
}
