package service

import "cloud-mall/app/api/internal/types"

type UserService struct {
}

func (us *UserService) Login(username, password string) (resp *types.LoginReply, err error) {
	return &types.LoginReply{
		Id:           1,
		Name:         username,
		Gender:       "male",
		AccessToken:  password,
		AccessExpire: 0,
		RefreshAfter: 0,
	}, nil
}
