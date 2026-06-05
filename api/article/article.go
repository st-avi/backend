// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package article

import (
	"context"

	"backend/api/article/v1"
)

type IArticleV1 interface {
	GetArticles10(ctx context.Context, req *v1.GetArticles10Req) (res *v1.GetArticles10Res, err error)
	GetArticles(ctx context.Context, req *v1.GetArticlesReq) (res *v1.GetArticlesRes, err error)
}
