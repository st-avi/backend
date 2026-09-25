package v1

import (
	"backend/api"
	"net/http"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/goai"
)

type GetArticelResTag struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type GetArticleReq struct {
	g.Meta `path:"/article" tags:"Article" summary:"取得文章" method:"get"`
	Slug   string `json:"slug" v:"required" dc:"文章 Slug"`
}

type GetArticleRes struct {
	g.Meta      `status:"200" resEg:"api/example/article/v1/article_get.json"`
	Title       string             `json:"title"`
	Summary     string             `json:"summary"`
	Content     string             `json:"content"`
	CoverImage  string             `json:"cover_image"`
	PublishedAt string             `json:"published_at"`
	Category    string             `json:"category"`
	Tags        []GetArticelResTag `json:"tags"`
}

func (r GetArticleRes) EnhanceResponseStatus() (resList map[int]goai.EnhancedStatusType) {
	return map[int]goai.EnhancedStatusType{
		http.StatusNotFound: {
			Response: struct{}{},
			Examples: []interface{}{
				api.CommonRes{
					Code:    gcode.CodeNotFound.Code(),
					Message: "找不到文章",
					Data:    nil,
				},
			},
		},
	}
}
