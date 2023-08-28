package main

import (
	"Douyin_Demo/model"
	"Douyin_Demo/repo"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetNewProfile(userId int64) (profile *model.UserProfile, err error) {
	strId := strconv.FormatInt(userId, 7)
	avatar := GenerateAvatar(strId)
	signature, err := GenerateSignature()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	background, err := GenerateBackground()
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	return &model.UserProfile{
		UserId:          userId,
		Name:            "User " + strId,
		Avatar:          avatar,
		Signature:       signature,
		BackgroundImage: background,
	}, nil
}

// only for test
func GenerateUserProfile(userId int64) (profile *model.UserProfile, err error) {
	strId := strconv.FormatInt(userId, 7)
	avatar := GenerateAvatar(strId)
	signature, err := GenerateSignature()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	background, err := GenerateBackground()
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	userP, err := saveUserProfile(userId, avatar, signature, background)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	return userP, nil
}

// TODO remove db code outside of handler
func saveUserProfile(userId int64, avatar string, signature string, background string) (userPro *model.UserProfile, err error) {
	strId := strconv.FormatInt(userId, 7)
	name := "User " + strId
	userPro = &model.UserProfile{
		UserId:          userId,
		Name:            name,
		Avatar:          avatar,
		Signature:       signature,
		BackgroundImage: background,
	}

	err = repo.Q.UserProfile.Create(userPro)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	return userPro, nil
}

func GenerateAvatar(userId string) string {
	return "https://api.multiavatar.com/" + userId + ".png"
}

func GenerateSignature() (signature string, err error) {
	url := "https://v1.hitokoto.cn/?encode=text"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return "wrong", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: Unexpected status code:", resp.StatusCode)
		return "wrong", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return "wrong", err
	}

	return string(body), nil
}

func GenerateBackground() (imgUrl string, err error) {
	url := "https://source.unsplash.com/random/1125x633"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: Unexpected status code:", resp.StatusCode)
		return "", err
	}

	hotlinkURL := resp.Request.URL.String()
	return hotlinkURL, nil
}
