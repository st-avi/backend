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

func (c *ControllerV1) GetCategory(ctx context.Context, req *v1.GetCategoryReq) (res *v1.GetCategoryRes, err error) {
	r := ghttp.RequestFromCtx(ctx)

	var resList []v1.GetCategoryData

	cacheKey := consts.CacheArticleCategory + req.Slug
	cacheList, _ := g.Redis().Get(ctx, cacheKey)
	if err := cacheList.Scan(&resList); err == nil && resList != nil {
		return &v1.GetCategoryRes{
			List: resList,
		}, nil
	}

	articles, err := dao.Articles.Ctx(ctx).As("a").
		Fields("a.title, a.slug, a.summary").
		LeftJoin("categories c", "c.id = a.category_id").
		Where("a.status", consts.ArticlePublished).
		Where("c.slug", req.Slug).
		Order("a.id ASC").
		All()
	if err != nil {
		r.Response.Status = http.StatusInternalServerError
		return nil, gerror.NewCode(gcode.CodeInternalError, "取得系列文章失敗")
	}
	if articles.IsEmpty() {
		r.Response.Status = http.StatusNotFound
		return nil, gerror.NewCode(gcode.CodeNotFound, "找不到該系列文章")
	}

	for _, article := range articles {
		resList = append(resList, v1.GetCategoryData{
			Title:   article["title"].String(),
			Slug:    article["slug"].String(),
			Summary: article["summary"].String(),
		})
	}

	_, _ = g.Redis().Set(ctx, cacheKey, resList)

	return &v1.GetCategoryRes{
		List: resList,
	}, nil
}
