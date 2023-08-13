package main

import (
	"Douyin_Demo/kitex_gen/douyin/publish"
	"Douyin_Demo/repo"
	"context"
	"os"
	"reflect"
	"testing"
)

func TestPublishServiceImpl_DouyinPublishAction(t *testing.T) {
	// get video file in curent directory
	testFile, err := os.ReadFile("./resource/test.mp4")
	if err != nil {
		t.Fatal(err)
	}
	repo.SetDefault(repo.DB)

	var mockNormalRequest = struct {
		ctx context.Context
		req *publish.DouyinPublishActionRequest
	}{
		ctx: context.Background(),
		req: &publish.DouyinPublishActionRequest{
			Title: "TestVideo",
			Data:  testFile,
			Token: "123456",
		}}

	var mockInvalidRequest = struct {
		ctx context.Context
		req *publish.DouyinPublishActionRequest
	}{
		ctx: context.Background(),
		req: &publish.DouyinPublishActionRequest{
			Title: "InvaildVideo",
			Data:  []byte{1, 2, 3, 4, 5},
			Token: "23455",
		},
	}

	// expected result
	var successResult = &publish.DouyinPublishActionResponse{
		StatusCode: 0,
	}

	type args struct {
		ctx context.Context
		req *publish.DouyinPublishActionRequest
	}

	tests := []struct {
		name    string
		args    args
		want    *publish.DouyinPublishActionResponse
		wantErr bool
	}{
		{
			name: "portaittestmp4",
			args: mockNormalRequest,
			want: successResult,
		},
		{
			name:    "invalidvideo",
			args:    mockInvalidRequest,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &PublishServiceImpl{}
			got, err := s.DouyinPublishAction(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("DouyinPublishActionServiceImpl.DouyinPublishAction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DouyinPublishActionServiceImpl.DouyinPublishAction() = %v, want %v", got, tt.want)
			}
		})
	}

}
