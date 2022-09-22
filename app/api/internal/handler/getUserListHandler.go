package handler

import (
	"cloud-mall/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"cloud-mall/app/api/internal/logic"
	"cloud-mall/app/api/internal/svc"
	"cloud-mall/app/api/internal/types"
)

func getUserListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CustomerList
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetUserListLogic(r.Context(), ctx)
		resp, err := l.GetUserList(&req)
		response.Response(w, resp, err)

	}
}
