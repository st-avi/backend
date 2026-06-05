package v1

import "github.com/gogf/gf/v2/frame/g"

type GetTagsResTag struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type GetTagsReq struct {
	g.Meta `path:"/tags" tags:"Article" summary:"取得所有文章標籤" method:"get"`
}

type GetTagsRes struct {
	g.Meta `status:"200" resEg:"api/article/v1/example/tags_get.json"`
	Tags   []GetTagsResTag `json:"tags"`
}
