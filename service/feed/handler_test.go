package main

import (
	"Douyin_Demo/constants"
	"Douyin_Demo/kitex_gen/douyin/feed"
	"Douyin_Demo/repo"
	"context"
	"reflect"
	"testing"
)

func TestFeedServiceImpl_GetVideoFeed(t *testing.T) {
	repo.SetDefault(repo.DB)

	type args struct {
		ctx context.Context
		req *feed.FeedRequest
	}

	successArgs := args{
		ctx: context.Background(),
		req: &feed.FeedRequest{
			LatestTime: nil,
		},
	}

	successResp := &feed.FeedResponse{
		StatusCode: constants.STATUS_SUCCESS,
		// TODO: implement next time
		VideoList: []*feed.Video{},
	}

	tests := []struct {
		name    string
		args    args
		want    *feed.FeedResponse
		wantErr bool
	}{
		{
			name: "normal",
			args: successArgs,
			want: successResp,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &FeedServiceImpl{}
			got, err := s.GetVideoFeed(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVideoFeed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.StatusCode, tt.want.StatusCode) {
				t.Errorf("GetVideoFeed() got = %v, want %v", got, tt.want)
			}
		})
	}
}
