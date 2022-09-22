package logic

import (
	"cloud-mall/app/api/internal/service"
	"cloud-mall/app/api/internal/svc"
	"cloud-mall/app/api/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginReply, err error) {
	userService := service.UserService{}
	return userService.Login(req.Username, req.Password)
}
