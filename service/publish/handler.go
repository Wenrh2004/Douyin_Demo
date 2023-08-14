package main

import (
	"Douyin_Demo/constants"
	"Douyin_Demo/kitex_gen/douyin/feed"
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
	fmt.Println("new publish model == > ", newPublishModel)

	//repo.Q.New
	err = repo.Q.WithContext(ctx).Publish.Create(&newPublishModel)
	//err = q.Publish.WithContext(ctx).Create(&newPublishModel)
	if err != nil {
		fmt.Println("create publish error == > ", err.Error())
		return nil, err
	}

	return &publish.DouyinPublishActionResponse{}, nil
}

// PublishList implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishList(ctx context.Context, req *publish.PublishListRequest) (resp *publish.PublishListResponse, err error) {
	// TODO: get requestingUserId from token

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

	// set to video list
	var videoList []*feed.Video
	for _, item := range publishList {
		// TODO: get author info by user service

		videoList = append(videoList, &feed.Video{
			Id:       int64(item.ID),
			PlayUrl:  item.PlayUrl,
			CoverUrl: item.CoverUrl,
			Title:    item.Title,
		})
	}

	return &publish.PublishListResponse{
		StatusCode: constants.STATUS_SUCCESS,
		StatusMsg:  nil,
		VideoList:  videoList,
	}, nil

}
