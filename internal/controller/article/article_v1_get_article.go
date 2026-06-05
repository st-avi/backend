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

func (c *ControllerV1) GetArticle(ctx context.Context, req *v1.GetArticleReq) (res *v1.GetArticleRes, err error) {
	r := ghttp.RequestFromCtx(ctx)

	var resData v1.GetArticleRes

	cacheInfo, _ := g.Redis().Get(ctx, consts.CacheArticleInfo+req.Slug)
	if err := cacheInfo.Scan(&resData); err == nil && resData.Title != "" {
		return &resData, nil
	}

	article, err := dao.Articles.Ctx(ctx).As("a").
		Fields("a.id, a.title, a.summary, a.content, a.cover_image, c.name").
		LeftJoin("categories c", "c.id = a.category_id").
		Where("a.status", consts.ArticlePublished).
		Where("a.slug", req.Slug).
		One()
	if err != nil || article == nil {
		r.Response.Status = http.StatusNotFound
		return nil, gerror.NewCode(gcode.CodeNotFound, "找不到文章")
	}

	resTags := make([]v1.GetArticelResTag, 0)
	err = dao.ArticleTags.Ctx(ctx).As("at").
		Fields("t.name, t.slug").
		LeftJoin("tags t", "t.id = at.tag_id").
		Where("at.article_id", article["id"].Int()).
		Order("t.id ASC").
		Scan(&resTags)
	if err != nil {
		r.Response.Status = http.StatusNotFound
		return nil, gerror.NewCode(gcode.CodeNotFound, "找不到文章")
	}

	resData = v1.GetArticleRes{
		Title:      article["title"].String(),
		Summary:    article["summary"].String(),
		Content:    article["content"].String(),
		CoverImage: filepath.Join(utility.S3Url, article["cover_image"].String()),
		Category:   article["name"].String(),
		Tags:       resTags,
	}

	_, _ = g.Redis().Set(ctx, consts.CacheArticleInfo+req.Slug, resData)

	return &resData, nil
}
