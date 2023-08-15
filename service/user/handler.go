package main

import (
	"Douyin_Demo/constants"
	user "Douyin_Demo/kitex_gen/douyin/user"
	"Douyin_Demo/repo"
	"context"
	"errors"
	"gorm.io/gorm"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// GetUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserInfo(ctx context.Context, req *user.UserInfoRequest) (resp *user.UserInfoResponse, err error) {

	queryId := req.UserId

	profileQ := repo.Q.UserProfile

	profile, err := profileQ.WithContext(ctx).Where(profileQ.UserId.Eq(queryId)).First()

	// TODO update here when register is done
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
		profile, err = GenerateUserProfile(queryId)
		if err != nil {
			// return error
			msg := constants.INTERNAL_SERVER_ERROR
			return &user.UserInfoResponse{
				StatusCode: constants.STATUS_INTERNAL_ERR,
				StatusMsg:  &msg,
			}, nil
		}
	}

	if err != nil {
		// TODO log error
		// return error
		msg := constants.DB_QUERY_FAILED
		return &user.UserInfoResponse{
			StatusCode: constants.STATUS_UNABLE_QUERY,
			StatusMsg:  &msg,
		}, nil

	}

	resUser := user.User{
		Id:              profile.UserId,
		Name:            profile.Name,
		Avatar:          &profile.Avatar,
		BackgroundImage: &profile.BackgroundImage,
	}

	// TODO: get like count from like service
	zero := int64(0)
	resUser.FavoriteCount = &zero
	resUser.TotalFavorited = &zero

	// TODO: get follow count from relation service
	resUser.FollowCount = &zero
	resUser.FollowerCount = &zero
	resUser.IsFollow = false

	// TODO: get publish count from publish service
	resUser.WorkCount = &zero

	return &user.UserInfoResponse{
		StatusCode: constants.STATUS_SUCCESS,
		StatusMsg:  nil,
		User:       &resUser,
	}, nil
}
