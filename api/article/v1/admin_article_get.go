package v1

import (
	"backend/api"
	"net/http"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/goai"
	"github.com/gogf/gf/v2/os/gtime"
)

type GetAdminArticleReq struct {
	g.Meta `path:"/article/{id}" tags:"Admin" summary:"取得後台文章（編輯用）" method:"get"`
	Id     int `json:"id" in:"path" v:"required|min:1" dc:"文章 ID"`
}

type GetAdminArticleRes struct {
	g.Meta      `status:"200" resEg:"api/example/article/v1/admin_article_get.json"`
	Id          int         `json:"id"`
	Title       string      `json:"title"`
	Slug        string      `json:"slug"`
	Summary     string      `json:"summary"`
	Content     string      `json:"content"`
	CoverImage  string      `json:"cover_image"`
	CategoryId  int         `json:"category_id" dc:"系列 ID，0 表示無系列"`
	TagIds      []int       `json:"tag_ids"`
	Status      string      `json:"status"`
	PublishedAt *gtime.Time `json:"published_at"`
	CreatedAt   *gtime.Time `json:"created_at"`
	UpdatedAt   *gtime.Time `json:"updated_at"`
}

func (r GetAdminArticleRes) EnhanceResponseStatus() (resList map[int]goai.EnhancedStatusType) {
	return map[int]goai.EnhancedStatusType{
		http.StatusUnauthorized: {
			Response: struct{}{},
			Examples: []any{
				api.CommonRes{
					Code:    gcode.CodeNotAuthorized.Code(),
					Message: "無效的 token",
					Data:    nil,
				},
			},
		},
		http.StatusForbidden: {
			Response: struct{}{},
			Examples: []any{
				api.CommonRes{
					Code:    gcode.CodeNotAuthorized.Code(),
					Message: "需要管理員權限",
					Data:    nil,
				},
			},
		},
		http.StatusNotFound: {
			Response: struct{}{},
			Examples: []any{
				api.CommonRes{
					Code:    gcode.CodeNotFound.Code(),
					Message: "找不到文章",
					Data:    nil,
				},
			},
		},
	}
}
