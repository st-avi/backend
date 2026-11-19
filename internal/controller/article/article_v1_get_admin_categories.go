package article

import (
	v1 "backend/api/article/v1"
	"backend/internal/dao"
	"context"
	"net/http"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (c *ControllerV1) GetAdminCategories(ctx context.Context, req *v1.GetAdminCategoriesReq) (res *v1.GetAdminCategoriesRes, err error) {
	r := ghttp.RequestFromCtx(ctx)

	categories, err := dao.Categories.Ctx(ctx).As("c").
		Fields("c.id, c.name, c.slug, COUNT(a.id) AS article_count").
		LeftJoin("articles a", "a.category_id = c.id").
		Group("c.id").
		Order("c.id ASC").
		All()
	if err != nil {
		r.Response.Status = http.StatusInternalServerError
		return nil, gerror.NewCode(gcode.CodeInternalError, "取得系列列表失敗")
	}

	list := make([]v1.GetAdminCategoriesResCategory, 0, len(categories))
	for _, category := range categories {
		list = append(list, v1.GetAdminCategoriesResCategory{
			Id:           category["id"].Int(),
			Name:         category["name"].String(),
			Slug:         category["slug"].String(),
			ArticleCount: category["article_count"].Int(),
		})
	}

	return &v1.GetAdminCategoriesRes{
		Categories: list,
	}, nil
}
