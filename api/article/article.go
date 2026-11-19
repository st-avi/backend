// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package article

import (
	"context"

	"backend/api/article/v1"
)

type IArticleV1 interface {
	GetAdminArticle(ctx context.Context, req *v1.GetAdminArticleReq) (res *v1.GetAdminArticleRes, err error)
	GetAdminArticles(ctx context.Context, req *v1.GetAdminArticlesReq) (res *v1.GetAdminArticlesRes, err error)
	CreateArticle(ctx context.Context, req *v1.CreateArticleReq) (res *v1.CreateArticleRes, err error)
	DeleteArticle(ctx context.Context, req *v1.DeleteArticleReq) (res *v1.DeleteArticleRes, err error)
	GetArticle(ctx context.Context, req *v1.GetArticleReq) (res *v1.GetArticleRes, err error)
	UpdateArticle(ctx context.Context, req *v1.UpdateArticleReq) (res *v1.UpdateArticleRes, err error)
	GetArticles10(ctx context.Context, req *v1.GetArticles10Req) (res *v1.GetArticles10Res, err error)
	GetArticles(ctx context.Context, req *v1.GetArticlesReq) (res *v1.GetArticlesRes, err error)
	GetCategories(ctx context.Context, req *v1.GetCategoriesReq) (res *v1.GetCategoriesRes, err error)
	GetCategory(ctx context.Context, req *v1.GetCategoryReq) (res *v1.GetCategoryRes, err error)
	GetTag(ctx context.Context, req *v1.GetTagReq) (res *v1.GetTagRes, err error)
	GetTags(ctx context.Context, req *v1.GetTagsReq) (res *v1.GetTagsRes, err error)
}
