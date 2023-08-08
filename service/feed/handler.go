package main

import (
	feed "Douyin_Demo/kitex_gen/douyin/feed"
	"context"
)

// FeedServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// GetVideoFeed implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) GetVideoFeed(ctx context.Context, req *feed.FeedRequest) (resp *feed.FeedResponse, err error) {
	// TODO: Your code here...
	return
}
