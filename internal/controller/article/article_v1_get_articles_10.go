package article

import (
	v1 "backend/api/article/v1"
	"backend/internal/consts"
	"backend/internal/dao"
	"backend/utility"
	"context"
	"net/http"
	"path/filepath"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (c *ControllerV1) GetArticles10(ctx context.Context, req *v1.GetArticles10Req) (res *v1.GetArticles10Res, err error) {
	r := ghttp.RequestFromCtx(ctx)

	var resList []v1.GetArticles10Data

	cacheList, _ := g.Redis().Get(ctx, consts.CacheArticleList10)
	if err := cacheList.Scan(&resList); err == nil && resList != nil {
		return &v1.GetArticles10Res{
			List: resList,
		}, nil
	}

	articles, err := dao.Articles.Ctx(ctx).As("a").
		Fields("a.title, a.slug, a.summary, a.cover_image, c.name").
		LeftJoin("categories c", "c.id = a.category_id").
		Where("a.status", consts.ArticlePublished).
		Order("a.published_at DESC").
		Limit(10).
		All()
	if err != nil {
		r.Response.Status = http.StatusInternalServerError
		return nil, gerror.NewCode(gcode.CodeInternalError, "取得最新文章失敗")
	}

	for _, article := range articles {
		resList = append(resList, v1.GetArticles10Data{
			Category:   article["name"].String(),
			Title:      article["title"].String(),
			Slug:       article["slug"].String(),
			Summary:    article["summary"].String(),
			CoverImage: filepath.Join(utility.S3Url, article["cover_image"].String()),
		})
	}

	_, _ = g.Redis().Set(ctx, consts.CacheArticleList10, resList)

	return &v1.GetArticles10Res{
		List: resList,
	}, nil
}
