package article

import (
	v1 "backend/api/article/v1"
	"backend/internal/consts"
	"backend/internal/dao"
	"context"
	"net/http"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (c *ControllerV1) GetArticles(ctx context.Context, req *v1.GetArticlesReq) (res *v1.GetArticlesRes, err error) {
	r := ghttp.RequestFromCtx(ctx)

	var resList []v1.GetArticlesData

	cacheList, _ := g.Redis().Get(ctx, consts.CacheArticleListALL)
	if err := cacheList.Scan(&resList); err == nil && resList != nil {
		return &v1.GetArticlesRes{
			List: resList,
		}, nil
	}

	articles, err := dao.Articles.Ctx(ctx).As("a").
		Fields("a.title, a.slug, c.name, c.slug as category_slug").
		LeftJoin("categories c", "c.id = a.category_id").
		Where("a.status", consts.ArticlePublished).
		Order("a.id DESC").
		All()
	if err != nil {
		r.Response.Status = http.StatusInternalServerError
		return nil, gerror.NewCode(gcode.CodeInternalError, "取得文章列表失敗")
	}

	for _, article := range articles {
		resList = append(resList, v1.GetArticlesData{
			Category:     article["name"].String(),
			CategorySlug: article["category_slug"].String(),
			Title:        article["title"].String(),
			Slug:         article["slug"].String(),
		})
	}

	_, _ = g.Redis().Set(ctx, consts.CacheArticleListALL, resList)

	return &v1.GetArticlesRes{
		List: resList,
	}, nil
}
