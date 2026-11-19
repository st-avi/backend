package article

import (
	v1 "backend/api/article/v1"
	"backend/internal/dao"
	"backend/internal/logic/article"
	"context"
	"net/http"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (c *ControllerV1) CreateCategory(ctx context.Context, req *v1.CreateCategoryReq) (res *v1.CreateCategoryRes, err error) {
	r := ghttp.RequestFromCtx(ctx)

	if err := checkNameSlugExist(r, dao.Categories.Ctx(ctx), req.Name, req.Slug, "系列"); err != nil {
		return nil, err
	}

	id, err := article.CreateCategory(ctx, *req)
	if err != nil {
		g.Log().Error(ctx, "新增系列失敗: ", err)
		r.Response.Status = http.StatusInternalServerError
		return nil, gerror.NewCode(gcode.CodeInternalError, "新增系列失敗")
	}

	return &v1.CreateCategoryRes{
		Id:   id,
		Name: req.Name,
		Slug: req.Slug,
	}, nil
}
