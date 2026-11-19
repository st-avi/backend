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

func (c *ControllerV1) UpdateTag(ctx context.Context, req *v1.UpdateTagReq) (res *v1.UpdateTagRes, err error) {
	r := ghttp.RequestFromCtx(ctx)

	// 取得原標籤 slug（清快取用）
	oldSlug, err := dao.Tags.Ctx(ctx).Where("id", req.Id).Value("slug")
	if err != nil {
		r.Response.Status = http.StatusInternalServerError
		return nil, gerror.NewCode(gcode.CodeInternalError)
	}
	if oldSlug.IsNil() {
		r.Response.Status = http.StatusNotFound
		return nil, gerror.NewCode(gcode.CodeNotFound, "找不到標籤")
	}

	// 檢查名稱與 slug 是否與其他標籤重複
	if err := checkNameSlugExist(r, dao.Tags.Ctx(ctx), req.Id, req.Name, req.Slug, "標籤"); err != nil {
		return nil, err
	}

	// 更新標籤
	if err := article.UpdateTag(ctx, *req); err != nil {
		g.Log().Error(ctx, "編輯標籤失敗: ", err)
		r.Response.Status = http.StatusInternalServerError
		return nil, gerror.NewCode(gcode.CodeInternalError, "編輯標籤失敗")
	}

	// 清除引用此標籤的前台快取
	clearTagCache(ctx, req.Id, oldSlug.String(), req.Slug)

	return &v1.UpdateTagRes{
		Id:   req.Id,
		Name: req.Name,
		Slug: req.Slug,
	}, nil
}
