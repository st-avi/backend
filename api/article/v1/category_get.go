package v1

import (
	"backend/api"
	"net/http"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/goai"
)

type GetCategoryData struct {
	Title   string `json:"title"`
	Slug    string `json:"slug"`
	Summary string `json:"summary"`
}

type GetCategoryReq struct {
	g.Meta `path:"/category" tags:"Article" summary:"取得該系列所有文章" method:"get"`
	Slug   string `json:"slug" v:"required" dc:"系列 Slug"`
}

type GetCategoryRes struct {
	g.Meta `status:"200" resEg:"api/example/article/v1/category_get.json"`
	List   []GetCategoryData `json:"list"`
}

func (r GetCategoryRes) EnhanceResponseStatus() (resList map[int]goai.EnhancedStatusType) {
	return map[int]goai.EnhancedStatusType{
		http.StatusNotFound: {
			Response: struct{}{},
			Examples: []any{
				api.CommonRes{
					Code:    gcode.CodeNotFound.Code(),
					Message: "找不到該系列文章",
					Data:    nil,
				},
			},
		},
	}
}
