// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	role "github.com/zouchangfu/goctlx/zztest/user/internal/handler/role"
	"github.com/zouchangfu/goctlx/zztest/user/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: role.LoginHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
	)
}
