package article

import (
	v1 "backend/api/article/v1"
	"backend/internal/consts"
	"backend/internal/dao"
	"context"
	"net/http"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (c *ControllerV1) GetAdminArticles(ctx context.Context, req *v1.GetAdminArticlesReq) (res *v1.GetAdminArticlesRes, err error) {
	r := ghttp.RequestFromCtx(ctx)

	m := dao.Articles.Ctx(ctx).As("a").
		Fields("a.id, a.title, a.slug, a.status, a.published_at, a.updated_at, c.name AS category, c.slug AS category_slug").
		LeftJoin("categories c", "c.id = a.category_id")
	if req.Status != "" {
		m = m.Where("a.status", req.Status)
	} else {
		// 預設不含 ArticleTrashed
		m = m.WhereIn("a.status", []consts.ArticleStatus{consts.ArticleDraft, consts.ArticlePublished})
	}

	articles, total, err := m.Order("a.id DESC").Page(req.Page, req.Size).AllAndCount(false)
	if err != nil {
		r.Response.Status = http.StatusInternalServerError
		return nil, gerror.NewCode(gcode.CodeInternalError, "取得文章列表失敗")
	}

	list := make([]v1.GetAdminArticlesData, 0, len(articles))
	for _, a := range articles {
		list = append(list, v1.GetAdminArticlesData{
			Id:           a["id"].Int(),
			Title:        a["title"].String(),
			Slug:         a["slug"].String(),
			Status:       a["status"].String(),
			Category:     a["category"].String(),
			CategorySlug: a["category_slug"].String(),
			PublishedAt:  a["published_at"].GTime(),
			UpdatedAt:    a["updated_at"].GTime(),
		})
	}

	return &v1.GetAdminArticlesRes{
		List:  list,
		Total: total,
	}, nil
}
