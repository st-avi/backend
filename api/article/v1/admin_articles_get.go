package v1

import (
	"backend/api"
	"net/http"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/goai"
	"github.com/gogf/gf/v2/os/gtime"
)

type GetAdminArticlesData struct {
	Id           int         `json:"id"`
	Title        string      `json:"title"`
	Slug         string      `json:"slug"`
	Status       string      `json:"status"`
	Category     string      `json:"category"`
	CategorySlug string      `json:"category_slug"`
	PublishedAt  *gtime.Time `json:"published_at"`
	UpdatedAt    *gtime.Time `json:"updated_at"`
}

type GetAdminArticlesReq struct {
	g.Meta `path:"/articles" tags:"Admin" summary:"取得後台文章列表" method:"get"`
	Status string `json:"status" in:"query" v:"in:draft,published,trashed" dc:"文章狀態篩選：不填為 draft + published（不含垃圾桶）"`
	Page   int    `json:"page" in:"query" d:"1" v:"min:1" dc:"頁碼"`
	Size   int    `json:"size" in:"query" d:"20" v:"between:1,100" dc:"每頁筆數（1-100）"`
}

type GetAdminArticlesRes struct {
	g.Meta `status:"200" resEg:"api/example/article/v1/admin_articles_get.json"`
	List   []GetAdminArticlesData `json:"list"`
	Total  int                    `json:"total"`
}

func (r GetAdminArticlesRes) EnhanceResponseStatus() (resList map[int]goai.EnhancedStatusType) {
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
	}
}
