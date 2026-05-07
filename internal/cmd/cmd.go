package cmd

import (
	"backend/api"
	"backend/internal/controller/auth"
	"backend/internal/controller/user"
	"backend/internal/jobs"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/glog"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			_, err = gcron.Add(ctx, "@every 30m", func(ctx context.Context) {
				_ = jobs.CleanCacheTable()
			}, "Clean cache table")
			if err != nil {
				glog.Errorf(ctx, "Clean cache table error: %v", err)
			}

			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Middleware(MiddlewareCORS)
				group.Bind(
					auth.NewV1().Login,
					auth.NewV1().Refresh,
					user.NewV1().CreateUser,
				)
				group.Middleware(MiddlewareAuth)
				group.Bind(
					auth.NewV1().AuthMe,
				)
			})
			oai := s.GetOpenApi()
			oai.Config.CommonResponse = api.CommonRes{}
			oai.Config.CommonResponseDataField = `Data`
			s.Run()
			return nil
		},
	}
)
