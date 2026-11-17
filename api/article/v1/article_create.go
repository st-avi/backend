package v1

import (
	"backend/api"
	"net/http"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/goai"
)

type CreateArticleReq struct {
	g.Meta     `path:"/article" tags:"Admin" summary:"新增文章" method:"post"`
	Title      string `json:"title" v:"required|max-length:255" dc:"文章標題"`
	Slug       string `json:"slug" v:"required|max-length:255|regex:^[a-z0-9]+(-[a-z0-9]+)*$" dc:"文章 Slug：僅限小寫英數字與連字號，例如 c-01"`
	Summary    string `json:"summary" dc:"文章摘要"`
	Content    string `json:"content" v:"required" dc:"文章內容（Markdown）"`
	CoverImage string `json:"cover_image" v:"max-length:500" dc:"封面圖片 S3 路徑（不含網域）"`
	CategoryId int    `json:"category_id" v:"min:0" dc:"系列 ID，不填或 0 表示無系列"`
	TagIds     []int  `json:"tag_ids" dc:"標籤 ID 列表"`
	Status     string `json:"status" d:"draft" v:"in:draft,published" dc:"文章狀態：draft 草稿、published 發佈"`
}

type CreateArticleRes struct {
	g.Meta `status:"200" resEg:"api/example/article/v1/article_create.json"`
	Id     int    `json:"id"`
	Slug   string `json:"slug"`
}

func (r CreateArticleRes) EnhanceResponseStatus() (resList map[int]goai.EnhancedStatusType) {
	return map[int]goai.EnhancedStatusType{
		http.StatusUnauthorized: {
			Response: struct{}{},
			Examples: []interface{}{
				api.CommonRes{
					Code:    gcode.CodeNotAuthorized.Code(),
					Message: "無效的 token",
					Data:    nil,
				},
			},
		},
		http.StatusForbidden: {
			Response: struct{}{},
			Examples: []interface{}{
				api.CommonRes{
					Code:    gcode.CodeNotAuthorized.Code(),
					Message: "需要管理員權限",
					Data:    nil,
				},
			},
		},
		http.StatusConflict: {
			Response: struct{}{},
			Examples: []interface{}{
				api.CommonRes{
					Code:    gcode.CodeBusinessValidationFailed.Code(),
					Message: "文章 Slug 已存在",
					Data:    nil,
				},
			},
		},
		http.StatusUnprocessableEntity: {
			Response: struct{}{},
			Examples: []interface{}{
				api.CommonRes{
					Code:    gcode.CodeBusinessValidationFailed.Code(),
					Message: "找不到系列",
					Data:    nil,
				},
				api.CommonRes{
					Code:    gcode.CodeBusinessValidationFailed.Code(),
					Message: "找不到部分標籤",
					Data:    nil,
				},
			},
		},
	}
}
