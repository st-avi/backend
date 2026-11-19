package article

import (
	v1 "backend/api/article/v1"
	"backend/internal/dao"
	"backend/internal/model/entity"
	"context"
	"net/http"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (c *ControllerV1) GetAdminArticle(ctx context.Context, req *v1.GetAdminArticleReq) (res *v1.GetAdminArticleRes, err error) {
	r := ghttp.RequestFromCtx(ctx)

	cols := dao.Articles.Columns()
	var a *entity.Articles
	err = dao.Articles.Ctx(ctx).
		FieldsEx(cols.SearchVector).
		Where(cols.Id, req.Id).
		Scan(&a)
	if err != nil {
		r.Response.Status = http.StatusInternalServerError
		return nil, gerror.NewCode(gcode.CodeInternalError)
	}
	if a == nil {
		r.Response.Status = http.StatusNotFound
		return nil, gerror.NewCode(gcode.CodeNotFound, "找不到文章")
	}

	tagCols := dao.ArticleTags.Columns()
	tagIds, err := dao.ArticleTags.Ctx(ctx).
		Where(tagCols.ArticleId, req.Id).
		OrderAsc(tagCols.TagId).
		Array(tagCols.TagId)
	if err != nil {
		r.Response.Status = http.StatusInternalServerError
		return nil, gerror.NewCode(gcode.CodeInternalError)
	}
	ids := make([]int, 0, len(tagIds))
	for _, id := range tagIds {
		ids = append(ids, id.Int())
	}

	return &v1.GetAdminArticleRes{
		Id:          a.Id,
		Title:       a.Title,
		Slug:        a.Slug,
		Summary:     a.Summary,
		Content:     a.Content,
		CoverImage:  a.CoverImage,
		CategoryId:  a.CategoryId,
		TagIds:      ids,
		Status:      a.Status,
		PublishedAt: a.PublishedAt,
		CreatedAt:   a.CreatedAt,
		UpdatedAt:   a.UpdatedAt,
	}, nil
}
