package user

import (
	"net/http"

	"gitee.com/cloud-mall/app/api/internal/logic/user"
	"gitee.com/cloud-mall/app/api/internal/svc"
	"gitee.com/cloud-mall/app/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetGoodsListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GoodsList
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewGetGoodsListLogic(r.Context(), svcCtx)
		resp, err := l.GetGoodsList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
