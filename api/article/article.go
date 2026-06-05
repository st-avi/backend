// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package article

import (
	"context"

	"backend/api/article/v1"
)

type IArticleV1 interface {
	GetArticle(ctx context.Context, req *v1.GetArticleReq) (res *v1.GetArticleRes, err error)
	GetArticles10(ctx context.Context, req *v1.GetArticles10Req) (res *v1.GetArticles10Res, err error)
	GetArticles(ctx context.Context, req *v1.GetArticlesReq) (res *v1.GetArticlesRes, err error)
	GetTags(ctx context.Context, req *v1.GetTagsReq) (res *v1.GetTagsRes, err error)
}
