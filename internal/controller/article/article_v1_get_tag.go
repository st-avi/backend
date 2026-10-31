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

func (c *ControllerV1) GetTag(ctx context.Context, req *v1.GetTagReq) (res *v1.GetTagRes, err error) {
	r := ghttp.RequestFromCtx(ctx)

	var resList []v1.GetTagData

	cacheKey := consts.CacheArticleTag + req.Slug
	cacheList, _ := g.Redis().Get(ctx, cacheKey)
	if err := cacheList.Scan(&resList); err == nil && resList != nil {
		return &v1.GetTagRes{
			List: resList,
		}, nil
	}

	articles, err := dao.Articles.Ctx(ctx).As("a").
		Fields("a.title, a.slug, a.summary").
		InnerJoin("article_tags at", "at.article_id = a.id").
		InnerJoin("tags t", "t.id = at.tag_id").
		Where("a.status", consts.ArticlePublished).
		Where("t.slug", req.Slug).
		Order("a.published_at DESC, a.id DESC").
		All()
	if err != nil {
		r.Response.Status = http.StatusInternalServerError
		return nil, gerror.NewCode(gcode.CodeInternalError, "取得標籤文章失敗")
	}
	if articles.IsEmpty() {
		r.Response.Status = http.StatusNotFound
		return nil, gerror.NewCode(gcode.CodeNotFound, "找不到該標籤文章")
	}

	for _, article := range articles {
		resList = append(resList, v1.GetTagData{
			Title:   article["title"].String(),
			Slug:    article["slug"].String(),
			Summary: article["summary"].String(),
		})
	}

	_, _ = g.Redis().Set(ctx, cacheKey, resList)

	return &v1.GetTagRes{
		List: resList,
	}, nil
}
