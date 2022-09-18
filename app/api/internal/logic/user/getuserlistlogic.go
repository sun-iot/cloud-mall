package user

import (
	"context"

	"gitee.com/cloud-mall/app/api/internal/svc"
	"gitee.com/cloud-mall/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserListLogic) GetUserList(req *types.CustomerList) (resp *types.CustomerListReply, err error) {
	// todo: add your logic here and delete this line

	return
}
