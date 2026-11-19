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

func (c *ControllerV1) CreateTag(ctx context.Context, req *v1.CreateTagReq) (res *v1.CreateTagRes, err error) {
	r := ghttp.RequestFromCtx(ctx)

	if err := checkNameSlugExist(r, dao.Tags.Ctx(ctx), 0, req.Name, req.Slug, "標籤"); err != nil {
		return nil, err
	}

	id, err := article.CreateTag(ctx, *req)
	if err != nil {
		g.Log().Error(ctx, "新增標籤失敗: ", err)
		r.Response.Status = http.StatusInternalServerError
		return nil, gerror.NewCode(gcode.CodeInternalError, "新增標籤失敗")
	}

	return &v1.CreateTagRes{
		Id:   id,
		Name: req.Name,
		Slug: req.Slug,
	}, nil
}
