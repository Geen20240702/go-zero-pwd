// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	pwdv1 "go-zero-pwd/app/pwd/api/internal/handler/pwd/v1"
	"go-zero-pwd/app/pwd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/create",
				Handler: pwdv1.CreatePwdHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/del",
				Handler: pwdv1.DelPwdHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/list",
				Handler: pwdv1.GetPwdListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/update/:id",
				Handler: pwdv1.UpdatePwdHandler(serverCtx),
			},
		},
		rest.WithPrefix("/pwd/v1"),
	)
}