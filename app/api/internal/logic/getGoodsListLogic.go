package logic

import (
	"context"

	"cloud-mall/app/api/internal/svc"
	"cloud-mall/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGoodsListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGoodsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoodsListLogic {
	return &GetGoodsListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGoodsListLogic) GetGoodsList(req *types.GoodsList) (resp *types.GoodsListReply, err error) {
	// todo: add your logic here and delete this line

	return
}
