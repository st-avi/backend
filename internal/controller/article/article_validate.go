package article

import (
	"backend/internal/dao"
	"context"
	"net/http"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
)

func checkCategory(ctx context.Context, r *ghttp.Request, categoryId int) (slug string, err error) {
	if categoryId <= 0 {
		return "", nil
	}
	category, err := dao.Categories.Ctx(ctx).Fields("slug").Where("id", categoryId).One()
	if err != nil {
		r.Response.Status = http.StatusInternalServerError
		return "", gerror.NewCode(gcode.CodeInternalError)
	}
	if category.IsEmpty() {
		r.Response.Status = http.StatusUnprocessableEntity
		return "", gerror.NewCode(gcode.CodeBusinessValidationFailed, "找不到系列")
	}
	return category["slug"].String(), nil
}

func checkTags(ctx context.Context, r *ghttp.Request, ids []int) (tagIds []int, tagSlugs []string, err error) {
	tagIds = make([]int, 0, len(ids))
	seen := make(map[int]struct{}, len(ids))
	for _, id := range ids {
		if _, ok := seen[id]; !ok {
			seen[id] = struct{}{}
			tagIds = append(tagIds, id)
		}
	}
	if len(tagIds) == 0 {
		return tagIds, nil, nil
	}

	tags, err := dao.Tags.Ctx(ctx).Fields("slug").WhereIn("id", tagIds).All()
	if err != nil {
		r.Response.Status = http.StatusInternalServerError
		return nil, nil, gerror.NewCode(gcode.CodeInternalError)
	}
	if len(tags) != len(tagIds) {
		r.Response.Status = http.StatusUnprocessableEntity
		return nil, nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "找不到部分標籤")
	}
	for _, tag := range tags {
		tagSlugs = append(tagSlugs, tag["slug"].String())
	}
	return tagIds, tagSlugs, nil
}

func checkNameSlugExist(r *ghttp.Request, m *gdb.Model, name, slug, label string) error {
	rows, err := m.Fields("name, slug").Where("name", name).WhereOr("slug", slug).All()
	if err != nil {
		r.Response.Status = http.StatusInternalServerError
		return gerror.NewCode(gcode.CodeInternalError)
	}
	for _, row := range rows {
		if row["name"].String() == name {
			r.Response.Status = http.StatusConflict
			return gerror.NewCode(gcode.CodeBusinessValidationFailed, label+"名稱已存在")
		}
	}
	if len(rows) > 0 {
		r.Response.Status = http.StatusConflict
		return gerror.NewCode(gcode.CodeBusinessValidationFailed, label+" Slug 已存在")
	}
	return nil
}
