package article

import (
	"backend/internal/consts"
	"backend/internal/dao"
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

func clearCategoryCache(ctx context.Context, categoryId int, oldSlug, newSlug string) {
	articleSlugs, err := dao.Articles.Ctx(ctx).
		Where("category_id", categoryId).
		Where("status", consts.ArticlePublished).
		Array("slug")
	if err != nil {
		g.Log().Error(ctx, "取得系列文章失敗: ", err)
		return
	}
	if len(articleSlugs) == 0 {
		return
	}

	cacheKeys := []string{
		consts.CacheArticleCategories,
		consts.CacheArticleCategory + oldSlug,
		consts.CacheArticleCategory + newSlug,
		consts.CacheArticleList10,
		consts.CacheArticleListALL,
	}
	for _, slug := range articleSlugs {
		cacheKeys = append(cacheKeys, consts.CacheArticleInfo+slug.String())
	}
	if _, err := g.Redis().Del(ctx, cacheKeys...); err != nil {
		g.Log().Error(ctx, "清除系列快取失敗: ", err)
	}
}

func clearTagCache(ctx context.Context, tagId int, oldSlug, newSlug string) {
	articleSlugs, err := dao.Articles.Ctx(ctx).As("a").
		InnerJoin("article_tags at", "at.article_id = a.id").
		Where("at.tag_id", tagId).
		Where("a.status", consts.ArticlePublished).
		Array("a.slug")
	if err != nil {
		g.Log().Error(ctx, "取得標籤文章失敗: ", err)
		return
	}
	if len(articleSlugs) == 0 {
		return
	}

	// 文章列表不顯示標籤，不需清除
	cacheKeys := []string{
		consts.CacheArticleTags,
		consts.CacheArticleTag + oldSlug,
		consts.CacheArticleTag + newSlug,
	}
	for _, slug := range articleSlugs {
		cacheKeys = append(cacheKeys, consts.CacheArticleInfo+slug.String())
	}
	if _, err := g.Redis().Del(ctx, cacheKeys...); err != nil {
		g.Log().Error(ctx, "清除標籤快取失敗: ", err)
	}
}
