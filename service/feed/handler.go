package main

import (
	"Douyin_Demo/constants"
	feed "Douyin_Demo/kitex_gen/douyin/feed"
	"Douyin_Demo/kitex_gen/douyin/user"
	"Douyin_Demo/repo"
	"context"
	"time"
)

// FeedServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// GetVideoFeed implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) GetVideoFeed(ctx context.Context, req *feed.FeedRequest) (resp *feed.FeedResponse, err error) {
	// get latest time from req, if not exist, set to now
	var latestTime int64
	if req.LatestTime != nil {
		latestTime = *req.LatestTime
	} else {
		latestTime = time.Now().UnixMilli()
	}

	publish := repo.Q.Publish

	// get feed list from repo with created_at < latestTime
	feedList, err := publish.Where(publish.CreatedAt.Lte(time.UnixMilli(latestTime))).Order(publish.CreatedAt.Desc()).Limit(20).Find()
	if err != nil {
		return nil, err
	}

	// nextTime is the last time of the feed list
	var nextTime int64
	if len(feedList) > 0 {
		nextTime = feedList[len(feedList)-1].CreatedAt.UnixMilli()
	} else {
		nextTime = latestTime
	}

	// create video list from feed list
	var videoList []*feed.Video
	for _, item := range feedList {
		// TODO: get user info from repo
		fakeUser := &user.User{
			Id:   int64(item.UserId),
			Name: "fake user",
		}

		videoList = append(videoList, &feed.Video{
			Id:       int64(item.ID),
			PlayUrl:  item.PlayUrl,
			CoverUrl: item.CoverUrl,
			Title:    item.Title,
			// TODO: get user info from repo
			Author: fakeUser,
			// TODO: implement like count
			FavoriteCount: 0,
			// TODO: implement comment count
			CommentCount: 0,
			// TODO: implement favorite
			IsFavorite: true,
		})
	}

	// create repsonse
	return &feed.FeedResponse{
		VideoList:  videoList,
		NextTime:   &nextTime,
		StatusCode: constants.STATUS_SUCCESS,
	}, nil
}

// GetVideo implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) GetVideo(ctx context.Context, req *feed.GetVideoRequest) (resp *feed.GetVideoResponse, err error) {
	// get param from req
	videoId := req.VideoId
	var queryToekn string

	if req.Token != nil {
		queryToekn = *req.Token
	} else {
		queryToekn = ""
	}

	// get video from db
	publish := repo.Q.Publish
	publishModel, err := publish.WithContext(ctx).Where(publish.VideoId.Eq(videoId)).First()
	if err != nil {
		msg := constants.DB_QUERY_FAILED
		return &feed.GetVideoResponse{
			StatusCode: constants.STATUS_UNABLE_QUERY,
			StatusMsg:  &msg,
		}, nil
	}

	userResp, err := getUserById(publishModel.UserId, queryToekn)

	if err != nil {
		msg := constants.INTERNAL_SERVER_ERROR
		return &feed.GetVideoResponse{
			StatusCode: constants.STATUS_INTERNAL_ERR,
			StatusMsg:  &msg,
		}, nil
	}

	// create response
	return &feed.GetVideoResponse{
		StatusCode: constants.STATUS_SUCCESS,
		StatusMsg:  nil,
		Video: &feed.Video{
			Id:       publishModel.VideoId,
			PlayUrl:  publishModel.PlayUrl,
			CoverUrl: publishModel.CoverUrl,
			Title:    publishModel.Title,
			Author:   userResp,
		},
	}, nil

}
