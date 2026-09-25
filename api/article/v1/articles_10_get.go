package v1

import "github.com/gogf/gf/v2/frame/g"

type GetArticles10Data struct {
	Category   string `json:"category"`
	Title      string `json:"title"`
	Slug       string `json:"slug"`
	Summary    string `json:"summary"`
	CoverImage string `json:"cover_image"`
}
type GetArticles10Req struct {
	g.Meta `path:"/articles/10" tags:"Article" summary:"取得最新10篇文章" method:"get"`
}

type GetArticles10Res struct {
	g.Meta `status:"200" resEg:"api/example/article/v1/articles_10_get.json"`
	List   []GetArticles10Data `json:"list"`
}
