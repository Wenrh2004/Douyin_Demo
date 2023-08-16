package main

import (
	"Douyin_Demo/constants"
	publish "Douyin_Demo/kitex_gen/douyin/publish"
	"Douyin_Demo/model"
	"Douyin_Demo/repo"
	"Douyin_Demo/service/storage"
	"bytes"
	"context"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

// PublishServiceImpl implements the last service interface defined in the IDL.
type PublishServiceImpl struct{}

// DouyinPublishAction implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) DouyinPublishAction(ctx context.Context, req *publish.DouyinPublishActionRequest) (resp *publish.DouyinPublishActionResponse, err error) {
	// check if is a valid video file
	if http.DetectContentType(req.Data) != "video/mp4" {
		// TODO log error
		msg := constants.INVALID_CONTENT_TYPE
		return &publish.DouyinPublishActionResponse{
			StatusCode: constants.PARAMS_ERROR_CODE,
			StatusMsg:  &msg,
		}, nil
	}

	// video file upload to s3
	fileReader := bytes.NewReader(req.Data)
	fileId := uuid.New().String()

	// TODO: get user id from token
	userId := req.Token
	fileName := fmt.Sprintf("%s-%s.mp4", userId, fileId)
	_, err = storage.UploadFile(fileReader, fileName)
	if err != nil {
		// TODO log error
		msg := constants.UPLOAD_FAILED
		return &publish.DouyinPublishActionResponse{
			StatusCode: constants.STATUS_INTERNAL_ERR,
			StatusMsg:  &msg,
		}, nil
	}

	// get file link
	fileLink := storage.GetObjectLink(fileName)

	// get file cover link
	coverLink, err := storage.GetThumbnailLink(fileName)
	if err != nil {
		// TODO log error
		msg := constants.GET_THUMBNAIL_LINK_FAILED
		fmt.Println("get thumbnail link error == > ", err.Error())

		return &publish.DouyinPublishActionResponse{
			StatusCode: constants.STATUS_INTERNAL_ERR,
			StatusMsg:  &msg,
		}, nil
	}

	// set to a model
	// TODO: get user id from token
	newPublishModel := model.Publish{
		UserId:   123456,
		Title:    req.Title,
		PlayUrl:  fileLink,
		CoverUrl: coverLink,
	}
	fmt.Println("new publish model == > ", newPublishModel)

	err = repo.Q.WithContext(ctx).Publish.Create(&newPublishModel)
	if err != nil {
		// TODO log error
		fmt.Println("create publish error == > ", err.Error())
		msg := constants.DB_SAVE_FAILED

		return &publish.DouyinPublishActionResponse{
			StatusCode: constants.STATUS_INTERNAL_ERR,
			StatusMsg:  &msg,
		}, nil
	}

	return &publish.DouyinPublishActionResponse{
		StatusCode: constants.STATUS_SUCCESS,
		StatusMsg:  nil,
	}, nil
}

// PublishList implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishList(ctx context.Context, req *publish.PublishListRequest) (resp *publish.PublishListResponse, err error) {
	// TODO: get requestingUserId from token
	var token string
	if req.Token != "" {
		token = req.Token
	} else {
		token = "mock_token"
	}
	// get user id from req
	userId := req.UserId
	publishQ := repo.Q.Publish

	// get publish list from db
	publishList, err := publishQ.WithContext(ctx).Where(publishQ.UserId.Eq(userId)).Order(publishQ.CreatedAt.Desc()).Find()

	if err != nil {
		fmt.Println("get publish list error == > ", err.Error())
		statusMsg := constants.DB_QUERY_FAILED
		return &publish.PublishListResponse{
			StatusCode: constants.STATUS_UNABLE_QUERY,
			StatusMsg:  &statusMsg,
		}, nil
	}

	// get author info by user service
	videoList, err := getVideos(publishList, token)

	if err != nil {
		fmt.Println("get videos error == > ", err.Error())
		statusMsg := err.Error()
		return &publish.PublishListResponse{
			StatusCode: constants.STATUS_INTERNAL_ERR,
			StatusMsg:  &statusMsg,
		}, nil
	}

	return &publish.PublishListResponse{
		StatusCode: constants.STATUS_SUCCESS,
		StatusMsg:  nil,
		VideoList:  videoList,
	}, nil

}
