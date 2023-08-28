package main

import (
	"Douyin_Demo/constants"
	"Douyin_Demo/kitex_gen/douyin/user"
	"Douyin_Demo/repo"
	"context"
	"testing"
)

func TestUserServiceImpl_UserRegister(t *testing.T) {
	repo.SetDefault(repo.DB)

	type args struct {
		ctx context.Context
		req *user.UserRegisterRequest
	}

	successArgs := args{
		ctx: context.Background(),
		req: &user.UserRegisterRequest{
			Username: "test@test.com",
			Password: "a12345678",
		},
	}

	successResp := &user.UserRegisterResponse{
		StatusCode: constants.STATUS_SUCCESS,
	}

	invaildArgs := []args{
		{
			ctx: context.Background(),
			req: &user.UserRegisterRequest{
				Username: "novalid",
				Password: "a12345678",
			},
		},
		{
			ctx: context.Background(),
			req: &user.UserRegisterRequest{
				Username: "valid@valid.com",
				Password: "123",
			},
		},
		{
			ctx: context.Background(),
			req: &user.UserRegisterRequest{
				Username: "test@test.com",
				Password: "a123456789",
			},
		},
	}

	tests := []struct {
		name    string
		args    args
		want    *user.UserRegisterResponse
		wantErr bool
	}{
		{
			name: "success",
			args: successArgs,
			want: successResp,
		},
		{
			name:    "invalid email",
			args:    invaildArgs[0],
			wantErr: true,
		},
		{
			name:    "invalid password",
			args:    invaildArgs[1],
			wantErr: true,
		},
		{
			name:    "email already exist",
			args:    invaildArgs[2],
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserServiceImpl{}
			got, err := s.UserRegister(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserServiceImpl.UserRegister() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.StatusCode != tt.want.StatusCode {
				t.Errorf("UserServiceImpl.UserRegister() = %v, want %v", got, tt.want)
			}
		})
	}
}
