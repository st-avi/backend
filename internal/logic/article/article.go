package article

import (
	v1 "backend/api/article/v1"
	"backend/internal/consts"
	"backend/internal/dao"
	"backend/internal/model/do"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
)

func CreateArticle(ctx context.Context, req v1.CreateArticleReq, tagIds []int) (id int, err error) {
	data := do.Articles{
		Title:      req.Title,
		Slug:       req.Slug,
		Summary:    req.Summary,
		Content:    req.Content,
		CoverImage: req.CoverImage,
		Status:     req.Status,
	}
	if req.CategoryId > 0 {
		data.CategoryId = req.CategoryId
	}
	if consts.ArticleStatus(req.Status) == consts.ArticlePublished {
		data.PublishedAt = gtime.Now()
	}

	err = dao.Articles.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		lastId, err := dao.Articles.Ctx(ctx).Data(data).InsertAndGetId()
		if err != nil {
			return err
		}
		id = int(lastId)

		if len(tagIds) == 0 {
			return nil
		}
		articleTags := make([]do.ArticleTags, 0, len(tagIds))
		for _, tagId := range tagIds {
			articleTags = append(articleTags, do.ArticleTags{
				ArticleId: id,
				TagId:     tagId,
			})
		}
		_, err = dao.ArticleTags.Ctx(ctx).Data(articleTags).Insert()
		return err
	})
	return id, err
}
