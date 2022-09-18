package user

import (
	"context"

	"gitee.com/cloud-mall/app/api/internal/svc"
	"gitee.com/cloud-mall/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRepoListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRepoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRepoListLogic {
	return &GetRepoListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRepoListLogic) GetRepoList(req *types.RepoList) (resp *types.RepoListReply, err error) {
	// todo: add your logic here and delete this line

	return
}
