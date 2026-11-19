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

func (c *ControllerV1) GetAdminTags(ctx context.Context, req *v1.GetAdminTagsReq) (res *v1.GetAdminTagsRes, err error) {
	r := ghttp.RequestFromCtx(ctx)

	tags, err := dao.Tags.Ctx(ctx).As("t").
		Fields("t.id, t.name, t.slug, COUNT(at.article_id) AS article_count").
		LeftJoin("article_tags at", "at.tag_id = t.id").
		Group("t.id").
		Order("t.id ASC").
		All()
	if err != nil {
		r.Response.Status = http.StatusInternalServerError
		return nil, gerror.NewCode(gcode.CodeInternalError, "取得標籤列表失敗")
	}

	list := make([]v1.GetAdminTagsResTag, 0, len(tags))
	for _, tag := range tags {
		list = append(list, v1.GetAdminTagsResTag{
			Id:           tag["id"].Int(),
			Name:         tag["name"].String(),
			Slug:         tag["slug"].String(),
			ArticleCount: tag["article_count"].Int(),
		})
	}

	return &v1.GetAdminTagsRes{
		Tags: list,
	}, nil
}
