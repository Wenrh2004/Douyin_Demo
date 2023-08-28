package main

import (
	"Douyin_Demo/constants"
	user "Douyin_Demo/kitex_gen/douyin/user"
	"Douyin_Demo/model"
	"Douyin_Demo/repo"
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
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

// UserRegister implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserRegister(ctx context.Context, req *user.UserRegisterRequest) (resp *user.UserRegisterResponse, err error) {

	// get params
	registerName := req.Username
	registerPassword := req.Password

	userQ := repo.Q.User

	// check if username exist
	exist, err := userQ.WithContext(ctx).Where(userQ.Username.Eq(registerName)).First()
	if exist != nil {
		msg := constants.EXIST_USERNAME
		return &user.UserRegisterResponse{
			StatusCode: constants.STATUS_FAILED,
			StatusMsg:  &msg,
		}, nil
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		msg := constants.DB_QUERY_FAILED
		return &user.UserRegisterResponse{
			StatusCode: constants.STATUS_UNABLE_QUERY,
			StatusMsg:  &msg,
		}, nil
	}

	// hashed password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerPassword), bcrypt.DefaultCost)
	if err != nil {
		msg := constants.INTERNAL_SERVER_ERROR

		return &user.UserRegisterResponse{
			StatusCode: constants.STATUS_INTERNAL_ERR,
			StatusMsg:  &msg,
		}, nil
	}

	registerUser := &model.User{
		Username: registerName,
		Password: string(hashedPassword),
	}

	err = repo.Q.Transaction(func(tx *repo.Query) error {
		// create user
		err = tx.User.Create(registerUser)
		if err != nil {
			return err
		}

		// create user profile
		profile, err := GetNewProfile(int64(registerUser.ID))
		if err != nil {
			return err
		}

		// save profile
		err = tx.UserProfile.Create(profile)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		msg := constants.DB_SAVE_FAILED
		return &user.UserRegisterResponse{
			StatusCode: constants.STATUS_UNABLE_SAVE,
			StatusMsg:  &msg,
		}, nil
	}

	return &user.UserRegisterResponse{
		UserId:     int64(registerUser.ID),
		StatusCode: constants.STATUS_SUCCESS,
		StatusMsg:  nil,
		// TODO get token
		Token: "1234",
	}, nil
}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, req *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	// TODO: Your code here...
	return
}
