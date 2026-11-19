package article

import (
	v1 "backend/api/article/v1"
	"backend/internal/consts"
	"backend/internal/dao"
	"backend/internal/logic/article"
	"context"
	"net/http"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (c *ControllerV1) DeleteArticle(ctx context.Context, req *v1.DeleteArticleReq) (res *v1.DeleteArticleRes, err error) {
	r := ghttp.RequestFromCtx(ctx)

	// 取得原文章與其系列 slug（清快取用）
	old, err := dao.Articles.Ctx(ctx).As("a").
		Fields("a.slug, a.status, c.slug AS category_slug").
		LeftJoin("categories c", "c.id = a.category_id").
		Where("a.id", req.Id).
		One()
	if err != nil {
		r.Response.Status = http.StatusInternalServerError
		return nil, gerror.NewCode(gcode.CodeInternalError)
	}
	if old.IsEmpty() {
		r.Response.Status = http.StatusNotFound
		return nil, gerror.NewCode(gcode.CodeNotFound, "找不到文章")
	}

	// 已在垃圾桶則不做任何變更
	status := consts.ArticleStatus(old["status"].String())
	if status == consts.ArticleTrashed {
		return &v1.DeleteArticleRes{Id: req.Id}, nil
	}

	// 移至垃圾桶
	if err := article.TrashArticle(ctx, req.Id); err != nil {
		g.Log().Error(ctx, "刪除文章失敗: ", err)
		r.Response.Status = http.StatusInternalServerError
		return nil, gerror.NewCode(gcode.CodeInternalError, "刪除文章失敗")
	}

	// 原本為發佈狀態時，清除前台相關快取（草稿前台本來就看不到）
	if status == consts.ArticlePublished {
		cacheKeys := []string{
			consts.CacheArticleList10,
			consts.CacheArticleListALL,
			consts.CacheArticleCategories,
			consts.CacheArticleTags,
			consts.CacheArticleInfo + old["slug"].String(),
		}
		if slug := old["category_slug"].String(); slug != "" {
			cacheKeys = append(cacheKeys, consts.CacheArticleCategory+slug)
		}
		tagSlugs, err := dao.Tags.Ctx(ctx).As("t").
			InnerJoin("article_tags at", "at.tag_id = t.id").
			Where("at.article_id", req.Id).
			Array("t.slug")
		if err != nil {
			g.Log().Error(ctx, "取得文章標籤失敗: ", err)
		}
		for _, slug := range tagSlugs {
			cacheKeys = append(cacheKeys, consts.CacheArticleTag+slug.String())
		}
		if _, err := g.Redis().Del(ctx, cacheKeys...); err != nil {
			g.Log().Error(ctx, "清除文章快取失敗: ", err)
		}
	}

	return &v1.DeleteArticleRes{Id: req.Id}, nil
}
