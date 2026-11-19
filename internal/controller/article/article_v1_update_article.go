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

func (c *ControllerV1) UpdateArticle(ctx context.Context, req *v1.UpdateArticleReq) (res *v1.UpdateArticleRes, err error) {
	r := ghttp.RequestFromCtx(ctx)

	// 取得原文章與其系列 slug（清快取用）
	old, err := dao.Articles.Ctx(ctx).As("a").
		Fields("a.slug, a.status, a.published_at, c.slug AS category_slug").
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

	// 檢查 slug 是否與其他文章重複
	exist, err := dao.Articles.Ctx(ctx).Where("slug", req.Slug).WhereNot("id", req.Id).Exist()
	if err != nil {
		r.Response.Status = http.StatusInternalServerError
		return nil, gerror.NewCode(gcode.CodeInternalError)
	}
	if exist {
		r.Response.Status = http.StatusConflict
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "文章 Slug 已存在")
	}

	// 檢查系列與標籤是否存在
	categorySlug, err := checkCategory(ctx, r, req.CategoryId)
	if err != nil {
		return nil, err
	}
	tagIds, tagSlugs, err := checkTags(ctx, r, req.TagIds)
	if err != nil {
		return nil, err
	}

	// 取得原標籤 slug（清快取用），需在更新前查詢
	oldTagSlugs, err := dao.Tags.Ctx(ctx).As("t").
		InnerJoin("article_tags at", "at.tag_id = t.id").
		Where("at.article_id", req.Id).
		Array("t.slug")
	if err != nil {
		r.Response.Status = http.StatusInternalServerError
		return nil, gerror.NewCode(gcode.CodeInternalError)
	}

	// 更新文章
	firstPublish := consts.ArticleStatus(req.Status) == consts.ArticlePublished && old["published_at"].IsNil()
	if err := article.UpdateArticle(ctx, *req, tagIds, firstPublish); err != nil {
		g.Log().Error(ctx, "編輯文章失敗: ", err)
		r.Response.Status = http.StatusInternalServerError
		return nil, gerror.NewCode(gcode.CodeInternalError, "編輯文章失敗")
	}

	// 編輯前或編輯後為發佈狀態時，清除前台相關快取（新舊 slug、系列、標籤都要清）
	wasPublished := consts.ArticleStatus(old["status"].String()) == consts.ArticlePublished
	if wasPublished || consts.ArticleStatus(req.Status) == consts.ArticlePublished {
		cacheKeys := []string{
			consts.CacheArticleList10,
			consts.CacheArticleListALL,
			consts.CacheArticleCategories,
			consts.CacheArticleTags,
			consts.CacheArticleInfo + old["slug"].String(),
			consts.CacheArticleInfo + req.Slug,
		}
		for _, slug := range []string{old["category_slug"].String(), categorySlug} {
			if slug != "" {
				cacheKeys = append(cacheKeys, consts.CacheArticleCategory+slug)
			}
		}
		for _, slug := range oldTagSlugs {
			cacheKeys = append(cacheKeys, consts.CacheArticleTag+slug.String())
		}
		for _, slug := range tagSlugs {
			cacheKeys = append(cacheKeys, consts.CacheArticleTag+slug)
		}
		if _, err := g.Redis().Del(ctx, cacheKeys...); err != nil {
			g.Log().Error(ctx, "清除文章快取失敗: ", err)
		}
	}

	return &v1.UpdateArticleRes{
		Id:   req.Id,
		Slug: req.Slug,
	}, nil
}
