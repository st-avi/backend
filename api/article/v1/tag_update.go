package v1

import (
	"backend/api"
	"net/http"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/goai"
)

type UpdateTagReq struct {
	g.Meta `path:"/tag/{id}" tags:"Admin" summary:"編輯標籤" method:"put"`
	Id     int    `json:"id" in:"path" v:"required|min:1" dc:"標籤 ID"`
	Name   string `json:"name" v:"required|max-length:100" dc:"標籤名稱"`
	Slug   string `json:"slug" v:"required|max-length:100|regex:^[a-z0-9]+(-[a-z0-9]+)*$" dc:"標籤 Slug：僅限小寫英數字與連字號"`
}

type UpdateTagRes struct {
	g.Meta `status:"200" resEg:"api/example/article/v1/tag_update.json"`
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Slug   string `json:"slug"`
}

func (r UpdateTagRes) EnhanceResponseStatus() (resList map[int]goai.EnhancedStatusType) {
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
					Message: "找不到標籤",
					Data:    nil,
				},
			},
		},
		http.StatusConflict: {
			Response: struct{}{},
			Examples: []any{
				api.CommonRes{
					Code:    gcode.CodeBusinessValidationFailed.Code(),
					Message: "標籤名稱已存在",
					Data:    nil,
				},
				api.CommonRes{
					Code:    gcode.CodeBusinessValidationFailed.Code(),
					Message: "標籤 Slug 已存在",
					Data:    nil,
				},
			},
		},
	}
}
