package main

import (
	action "Douyin_Demo/kitex_gen/douyin/publish/action"
	"Douyin_Demo/model"
	"Douyin_Demo/repo"
	"Douyin_Demo/service/storage"
	"bytes"
	"context"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

// DouyinPublishActionServiceImpl implements the last service interface defined in the IDL.
type DouyinPublishActionServiceImpl struct{}

// DouyinPublishAction implements the DouyinPublishActionServiceImpl interface.
func (s *DouyinPublishActionServiceImpl) DouyinPublishAction(ctx context.Context, req *action.DouyinPublishActionRequest) (resp *action.DouyinPublishActionResponse, err error) {
	// check if is a valid video file
	if http.DetectContentType(req.Data) != "video/mp4" {
		return nil, fmt.Errorf("invalid video file")
	}

	// video file upload to s3
	fileReader := bytes.NewReader(req.Data)
	fileId := uuid.New().String()

	// TODO: get user id from token
	userId := req.Token
	fileName := fmt.Sprintf("%s-%s.mp4", userId, fileId)
	_, err = storage.UploadFile(fileReader, fileName)
	if err != nil {
		fmt.Println("upload file error == > ", err.Error())
		return nil, err
	}
	// get file link
	fileLink := storage.GetObjectLink(fileName)

	// get file cover link
	coverLink, err := storage.GetThumbnailLink(fileName)
	if err != nil {
		fmt.Println("get thumbnail link error == > ", err.Error())
		return nil, err
	}

	// set to a model
	// TODO: get user id from token
	newPublishModel := model.Publish{
		UserId:   123456,
		Title:    req.Title,
		PlayUrl:  fileLink,
		CoverUrl: coverLink,
	}

	// publish.withContext(ctx).Create()
	err = repo.Q.Publish.WithContext(ctx).Create(&newPublishModel)
	if err != nil {
		fmt.Println("create publish error == > ", err.Error())
		return nil, err
	}

	return &action.DouyinPublishActionResponse{}, nil

}
