package v1

import (
	"backend/api"
	"net/http"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/goai"
)

type GetTagData struct {
	Title   string `json:"title"`
	Slug    string `json:"slug"`
	Summary string `json:"summary"`
}

type GetTagReq struct {
	g.Meta `path:"/tag" tags:"Article" summary:"取得該標籤所有文章" method:"GET"`
	Slug   string `json:"slug" v:"required" dc:"標籤 Slug"`
}

type GetTagRes struct {
	g.Meta `status:"200" resEg:"api/example/article/v1/tag_get.json"`
	List   []GetTagData `json:"list"`
}

func (r GetTagRes) EnhanceResponseStatus() (resList map[int]goai.EnhancedStatusType) {
	return map[int]goai.EnhancedStatusType{
		http.StatusNotFound: {
			Response: struct{}{},
			Examples: []any{
				api.CommonRes{
					Code:    gcode.CodeNotFound.Code(),
					Message: "找不到該標籤文章",
					Data:    nil,
				},
			},
		},
	}
}
