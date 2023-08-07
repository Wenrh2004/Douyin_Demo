package main

import (
	"Douyin_Demo/kitex_gen/douyin/publish/action"
	"context"
	"os"
	"testing"
)

func TestDouyinPublishActionServiceImpl_DouyinPublishAction(t *testing.T) {
	// get video file in curent directory
	testFile, err := os.ReadFile("./resource/test.mp4")
	if err != nil {
		t.Fatal(err)
	}

	var mockRequest = struct {
		ctx context.Context
		req *action.DouyinPublishActionRequest
	}{
		ctx: context.Background(),
		req: &action.DouyinPublishActionRequest{
			Title: "TestVideo",
			Data:  testFile,
			Token: "123456",
		}}

	// expected result
	var successResult = &action.DouyinPublishActionResponse{
		StatusCode: 0,
	}

	type args struct {
		ctx context.Context
		req *action.DouyinPublishActionRequest
	}

	tests := []struct {
		name    string
		args    args
		want    *action.DouyinPublishActionResponse
		wantErr bool
	}{
		{
			name: "success",
			args: mockRequest,
			want: successResult,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &DouyinPublishActionServiceImpl{}
			got, err := s.DouyinPublishAction(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("DouyinPublishActionServiceImpl.DouyinPublishAction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.StatusCode != tt.want.StatusCode {
				t.Errorf("DouyinPublishActionServiceImpl.DouyinPublishAction() = %v, want %v", got, tt.want)
			}
		})
	}

}
