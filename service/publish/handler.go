package main

import (
	action "Douyin_Demo/kitex_gen/douyin/publish/action"
	"context"
)

// DouyinPublishActionServiceImpl implements the last service interface defined in the IDL.
type DouyinPublishActionServiceImpl struct{}

// DouyinPublishAction implements the DouyinPublishActionServiceImpl interface.
func (s *DouyinPublishActionServiceImpl) DouyinPublishAction(ctx context.Context, req *action.DouyinPublishActionRequest) (resp *action.DouyinPublishActionResponse, err error) {
	// TODO: Your code here...
	return
}
