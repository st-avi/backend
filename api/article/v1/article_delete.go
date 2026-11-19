package v1

import (
	"backend/api"
	"net/http"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/goai"
)

type DeleteArticleReq struct {
	g.Meta `path:"/article/{id}" tags:"Admin" summary:"刪除文章（移至垃圾桶）" method:"delete" dc:"將文章狀態改為 trashed；已在垃圾桶的文章不做任何變更。還原請用編輯文章 API 將狀態改回 draft 或 published"`
	Id     int `json:"id" in:"path" v:"required|min:1" dc:"文章 ID"`
}

type DeleteArticleRes struct {
	g.Meta `status:"200" resEg:"api/example/article/v1/article_delete.json"`
	Id     int `json:"id"`
}

func (r DeleteArticleRes) EnhanceResponseStatus() (resList map[int]goai.EnhancedStatusType) {
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
