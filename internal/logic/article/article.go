package article

import (
	v1 "backend/api/article/v1"
	"backend/internal/consts"
	"backend/internal/dao"
	"backend/internal/model/do"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
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

func UpdateArticle(ctx context.Context, req v1.UpdateArticleReq, tagIds []int, firstPublish bool) error {
	cols := dao.Articles.Columns()
	data := g.Map{
		cols.Title:      req.Title,
		cols.Slug:       req.Slug,
		cols.Summary:    req.Summary,
		cols.Content:    req.Content,
		cols.CoverImage: req.CoverImage,
		cols.Status:     req.Status,
		cols.CategoryId: nil,
	}
	if req.CategoryId > 0 {
		data[cols.CategoryId] = req.CategoryId
	}
	if firstPublish {
		data[cols.PublishedAt] = gtime.Now()
	}

	return dao.Articles.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		if _, err := dao.Articles.Ctx(ctx).Data(data).Where(cols.Id, req.Id).Update(); err != nil {
			return err
		}

		if _, err := dao.ArticleTags.Ctx(ctx).Where(dao.ArticleTags.Columns().ArticleId, req.Id).Delete(); err != nil {
			return err
		}
		if len(tagIds) == 0 {
			return nil
		}
		articleTags := make([]do.ArticleTags, 0, len(tagIds))
		for _, tagId := range tagIds {
			articleTags = append(articleTags, do.ArticleTags{
				ArticleId: req.Id,
				TagId:     tagId,
			})
		}
		_, err := dao.ArticleTags.Ctx(ctx).Data(articleTags).Insert()
		return err
	})
}

func TrashArticle(ctx context.Context, id int) error {
	cols := dao.Articles.Columns()
	_, err := dao.Articles.Ctx(ctx).
		Data(cols.Status, consts.ArticleTrashed).
		Where(cols.Id, id).
		Update()
	return err
}
