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

func (c *ControllerV1) CreateArticle(ctx context.Context, req *v1.CreateArticleReq) (res *v1.CreateArticleRes, err error) {
	r := ghttp.RequestFromCtx(ctx)

	// 檢查 Slug 是否重複
	exist, err := dao.Articles.Ctx(ctx).Where("slug", req.Slug).Exist()
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

	// 新增文章
	id, err := article.CreateArticle(ctx, *req, tagIds)
	if err != nil {
		g.Log().Error(ctx, "新增文章失敗: ", err)
		r.Response.Status = http.StatusInternalServerError
		return nil, gerror.NewCode(gcode.CodeInternalError, "新增文章失敗")
	}

	// 直接發佈時清除前台相關快取
	if consts.ArticleStatus(req.Status) == consts.ArticlePublished {
		cacheKeys := []string{
			consts.CacheArticleList10,
			consts.CacheArticleListALL,
			consts.CacheArticleCategories,
			consts.CacheArticleTags,
		}
		if categorySlug != "" {
			cacheKeys = append(cacheKeys, consts.CacheArticleCategory+categorySlug)
		}
		for _, slug := range tagSlugs {
			cacheKeys = append(cacheKeys, consts.CacheArticleTag+slug)
		}
		if _, err := g.Redis().Del(ctx, cacheKeys...); err != nil {
			g.Log().Error(ctx, "清除文章快取失敗: ", err)
		}
	}

	return &v1.CreateArticleRes{
		Id:   id,
		Slug: req.Slug,
	}, nil
}
