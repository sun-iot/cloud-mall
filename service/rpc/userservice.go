package rpc

import (
	"context"
	"gitee.com/cloud-mall/proto"
	"github.com/pkg/errors"
)

// UserService 用户服务后端
// @Params LoginName 登录用户名
// @Params LoginPassword 登录用户密码
type UserService struct {
	LoginName string
	LoginPassword string
}

func (us UserService) Login(ctx context.Context) (*proto.ReturnLoginUser , error){
	if us.LoginPassword == "" || us.LoginName == "" {
		return nil , errors.Wrap(nil , "Please confirm the validity of the input.")
	}
	return &proto.ReturnLoginUser{
		Token: "OK",
	} , nil
}
