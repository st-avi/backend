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

func (c *ControllerV1) UpdateCategory(ctx context.Context, req *v1.UpdateCategoryReq) (res *v1.UpdateCategoryRes, err error) {
	r := ghttp.RequestFromCtx(ctx)

	// 取得原系列 slug（清快取用）
	oldSlug, err := dao.Categories.Ctx(ctx).Where("id", req.Id).Value("slug")
	if err != nil {
		r.Response.Status = http.StatusInternalServerError
		return nil, gerror.NewCode(gcode.CodeInternalError)
	}
	if oldSlug.IsNil() {
		r.Response.Status = http.StatusNotFound
		return nil, gerror.NewCode(gcode.CodeNotFound, "找不到系列")
	}

	// 檢查名稱與 slug 是否與其他系列重複
	if err := checkNameSlugExist(r, dao.Categories.Ctx(ctx), req.Id, req.Name, req.Slug, "系列"); err != nil {
		return nil, err
	}

	// 更新系列
	if err := article.UpdateCategory(ctx, *req); err != nil {
		g.Log().Error(ctx, "編輯系列失敗: ", err)
		r.Response.Status = http.StatusInternalServerError
		return nil, gerror.NewCode(gcode.CodeInternalError, "編輯系列失敗")
	}

	// 清除引用此系列的前台快取
	clearCategoryCache(ctx, req.Id, oldSlug.String(), req.Slug)

	return &v1.UpdateCategoryRes{
		Id:   req.Id,
		Name: req.Name,
		Slug: req.Slug,
	}, nil
}
