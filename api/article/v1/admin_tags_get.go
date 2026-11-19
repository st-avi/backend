package v1

import (
	"backend/api"
	"net/http"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/goai"
)

type GetAdminTagsResTag struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Slug         string `json:"slug"`
	ArticleCount int    `json:"article_count" dc:"文章數（含草稿與垃圾桶）"`
}

type GetAdminTagsReq struct {
	g.Meta `path:"/tags" tags:"Admin" summary:"取得後台標籤列表" method:"get" dc:"回傳所有標籤（含尚無文章的），供文章表單選擇與標籤管理使用"`
}

type GetAdminTagsRes struct {
	g.Meta `status:"200" resEg:"api/example/article/v1/admin_tags_get.json"`
	Tags   []GetAdminTagsResTag `json:"tags"`
}

func (r GetAdminTagsRes) EnhanceResponseStatus() (resList map[int]goai.EnhancedStatusType) {
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
