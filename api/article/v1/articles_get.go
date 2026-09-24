package v1

import "github.com/gogf/gf/v2/frame/g"

type GetArticlesData struct {
	Category string `json:"category"`
	Title    string `json:"title"`
	Slug     string `json:"slug"`
}

type GetArticlesReq struct {
	g.Meta `path:"/articles" tags:"Article" summary:"取得所有文章列表" method:"get"`
}

type GetArticlesRes struct {
	g.Meta `status:"200" resEg:"api/article/v1/example/articles_get.json"`
	List   []GetArticlesData `json:"list"`
}
