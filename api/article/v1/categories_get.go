package v1

import "github.com/gogf/gf/v2/frame/g"

type GetCategoriesResCategory struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type GetCategoriesReq struct {
	g.Meta `path:"/categories" tags:"Article" summary:"取得所有文章系列" method:"get"`
}

type GetCategoriesRes struct {
	g.Meta     `status:"200" resEg:"api/example/article/v1/categories_get.json"`
	Categories []GetCategoriesResCategory `json:"categories"`
}
