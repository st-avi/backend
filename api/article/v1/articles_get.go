package v1

import "github.com/gogf/gf/v2/frame/g"

type GetArticlesResArticle struct {
	Title string `json:"title"`
	Slug  string `json:"slug"`
}

type GetArticlesResCategory struct {
	Name     string                  `json:"name"`
	Articles []GetArticlesResArticle `json:"articles"`
}

type GetArticlesReq struct {
	g.Meta `path:"/articles" tags:"Article" summary:"取得所有文章列表" method:"get"`
}

type GetArticlesRes struct {
	g.Meta     `status:"200" resEg:"api/article/v1/example/articles_get.json"`
	Categories []GetArticlesResCategory `json:"categories"`
}
