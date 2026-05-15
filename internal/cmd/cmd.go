package cmd

import (
	"backend/api"
	"backend/internal/controller/auth"
	"backend/internal/controller/user"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Middleware(MiddlewareCORS)
				group.Bind(
					auth.NewV1().Login,
					auth.NewV1().Refresh,
					user.NewV1().CreateUser,
				)
				group.Group("/", func(authGroup *ghttp.RouterGroup) {
					authGroup.Middleware(MiddlewareAuth)
					authGroup.Bind(
						auth.NewV1().AuthMe,
					)
				})
				group.Group("/admin", func(adminGroup *ghttp.RouterGroup) {
					adminGroup.Middleware(MiddlewareAuth)
					adminGroup.Middleware(MiddlewareAdmin)
					adminGroup.Bind(
						user.NewV1().GetUsers,
					)
				})
			})
			oai := s.GetOpenApi()
			oai.Config.CommonResponse = api.CommonRes{}
			oai.Config.CommonResponseDataField = `Data`
			s.Run()
			return nil
		},
	}
)
