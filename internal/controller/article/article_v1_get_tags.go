package article

import (
	v1 "backend/api/article/v1"
	"backend/internal/consts"
	"backend/internal/dao"
	"context"
	"net/http"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (c *ControllerV1) GetTags(ctx context.Context, req *v1.GetTagsReq) (res *v1.GetTagsRes, err error) {
	r := ghttp.RequestFromCtx(ctx)

	var resTags []v1.GetTagsResTag

	cacheTags, _ := g.Redis().Get(ctx, consts.CacheArticleTags)
	if err := cacheTags.Scan(&resTags); err == nil && resTags != nil {
		return &v1.GetTagsRes{
			Tags: resTags,
		}, nil
	}

	tags, err := dao.Tags.Ctx(ctx).Fields("name, slug").Order("id ASC").All()
	if err != nil {
		r.Response.Status = http.StatusInternalServerError
		return nil, gerror.NewCode(gcode.CodeInternalError, "取得文章標籤失敗")
	}

	for _, tag := range tags {
		resTags = append(resTags, v1.GetTagsResTag{
			Name: tag["name"].String(),
			Slug: tag["slug"].String(),
		})
	}

	_, _ = g.Redis().Set(ctx, consts.CacheArticleTags, resTags)

	return &v1.GetTagsRes{
		Tags: resTags,
	}, nil
}
