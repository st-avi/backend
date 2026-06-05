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

	var resCategories []v1.GetArticlesResCategory

	cacheList, _ := g.Redis().Get(ctx, consts.CacheArticleListAside)
	if err := cacheList.Scan(&resCategories); err == nil && resCategories != nil {
		return &v1.GetArticlesRes{
			Categories: resCategories,
		}, nil
	}

	articles, err := dao.Articles.Ctx(ctx).As("a").
		Fields("a.title, a.slug, c.name").
		LeftJoin("categories c", "c.id = a.category_id").
		Where("a.status", consts.ArticlePublished).
		Order("a.category_id ASC, a.id ASC").
		All()
	if err != nil {
		r.Response.Status = http.StatusInternalServerError
		return nil, gerror.NewCode(gcode.CodeInternalError, "取得文章列表失敗")
	}

	resArticles := map[string][]v1.GetArticlesResArticle{}
	for _, article := range articles {
		name := article["name"].String()
		resArticles[name] = append(resArticles[name], v1.GetArticlesResArticle{
			Title: article["title"].String(),
			Slug:  article["slug"].String(),
		})
	}
	for name, articles := range resArticles {
		resCategories = append(resCategories, v1.GetArticlesResCategory{
			Name:     name,
			Articles: articles,
		})
	}

	_, _ = g.Redis().Set(ctx, consts.CacheArticleListAside, resCategories)

	return &v1.GetArticlesRes{
		Categories: resCategories,
	}, nil
}
