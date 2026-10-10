package article

import (
	"backend/internal/consts"
	"backend/internal/dao"
	"context"
	"net/http"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"

	"backend/api/article/v1"
)

func (c *ControllerV1) GetCategories(ctx context.Context, req *v1.GetCategoriesReq) (res *v1.GetCategoriesRes, err error) {
	r := ghttp.RequestFromCtx(ctx)

	var resCategories []v1.GetCategoriesResCategory

	cacheCategories, _ := g.Redis().Get(ctx, consts.CacheArticleCategories)
	if err := cacheCategories.Scan(&resCategories); err == nil && resCategories != nil {
		return &v1.GetCategoriesRes{
			Categories: resCategories,
		}, nil
	}

	categories, err := dao.Categories.Ctx(ctx).Fields("name, slug").Order("id ASC").All()
	if err != nil {
		r.Response.Status = http.StatusInternalServerError
		return nil, gerror.NewCode(gcode.CodeInternalError, "取得文章系列失敗")
	}

	for _, category := range categories {
		resCategories = append(resCategories, v1.GetCategoriesResCategory{
			Name: category["name"].String(),
			Slug: category["slug"].String(),
		})
	}

	_, _ = g.Redis().Set(ctx, consts.CacheArticleCategories, resCategories)

	return &v1.GetCategoriesRes{
		Categories: resCategories,
	}, nil
}
