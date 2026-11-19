package v1

import (
	"backend/api"
	"net/http"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/goai"
)

type CreateCategoryReq struct {
	g.Meta `path:"/category" tags:"Admin" summary:"新增系列" method:"post"`
	Name   string `json:"name" v:"required|max-length:100" dc:"系列名稱"`
	Slug   string `json:"slug" v:"required|max-length:100|regex:^[a-z0-9]+(-[a-z0-9]+)*$" dc:"系列 Slug：僅限小寫英數字與連字號"`
}

type CreateCategoryRes struct {
	g.Meta `status:"200" resEg:"api/example/article/v1/category_create.json"`
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Slug   string `json:"slug"`
}

func (r CreateCategoryRes) EnhanceResponseStatus() (resList map[int]goai.EnhancedStatusType) {
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
		http.StatusConflict: {
			Response: struct{}{},
			Examples: []any{
				api.CommonRes{
					Code:    gcode.CodeBusinessValidationFailed.Code(),
					Message: "系列名稱已存在",
					Data:    nil,
				},
				api.CommonRes{
					Code:    gcode.CodeBusinessValidationFailed.Code(),
					Message: "系列 Slug 已存在",
					Data:    nil,
				},
			},
		},
	}
}
