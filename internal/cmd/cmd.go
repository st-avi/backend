package cmd

import (
	"backend/api"
	"backend/internal/controller/auth"
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
				group.Bind(
					auth.NewV1(),
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
