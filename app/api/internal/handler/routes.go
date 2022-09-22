// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"cloud-mall/app/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/v1/user/login",
				Handler: loginHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/user/list",
				Handler: getUserListHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/repo/list",
				Handler: getRepoListHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/goods/list",
				Handler: getGoodsListHandler(serverCtx),
			},
		},
	)
}