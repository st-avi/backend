package cmd

import (
	"backend/api"
	"backend/internal/controller/article"
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
					article.NewV1().GetArticle,
					article.NewV1().GetArticles10,
					article.NewV1().GetArticles,
					article.NewV1().GetCategories,
					article.NewV1().GetCategory,
					article.NewV1().GetTag,
					article.NewV1().GetTags,
					article.NewV1().SearchArticles,
				)
				group.Group("/", func(authGroup *ghttp.RouterGroup) {
					authGroup.Middleware(MiddlewareAuth)
					authGroup.Bind(
						auth.NewV1().AuthMe,
						auth.NewV1().Logout,
					)
				})
				group.Group("/admin", func(adminGroup *ghttp.RouterGroup) {
					adminGroup.Middleware(MiddlewareAuth)
					adminGroup.Middleware(MiddlewareAdmin)
					adminGroup.Bind(
						user.NewV1().GetUsers,
						article.NewV1().CreateArticle,
						article.NewV1().UpdateArticle,
						article.NewV1().GetAdminArticles,
						article.NewV1().GetAdminArticle,
						article.NewV1().DeleteArticle,
						article.NewV1().GetAdminCategories,
						article.NewV1().GetAdminTags,
						article.NewV1().CreateCategory,
						article.NewV1().CreateTag,
						article.NewV1().UpdateCategory,
						article.NewV1().UpdateTag,
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
