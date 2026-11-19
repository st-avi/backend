package v1

import "github.com/gogf/gf/v2/frame/g"

type SearchArticlesData struct {
	Category     string `json:"category"`
	CategorySlug string `json:"category_slug"`
	Title        string `json:"title"`
	Slug         string `json:"slug"`
	Summary      string `json:"summary"`
	CoverImage   string `json:"cover_image"`
	PublishedAt  string `json:"published_at"`
}

type SearchArticlesReq struct {
	g.Meta `path:"/search" tags:"Article" summary:"搜尋文章" method:"get" dc:"以空白分隔多個關鍵字，每個關鍵字都需出現在標題、摘要或內文其中之一。排序：標題命中 > 摘要命中 > 內文命中，同分依發佈時間新到舊"`
	Q      string `json:"q" in:"query" v:"required|max-length:50" dc:"搜尋關鍵字（最多取前 5 個）"`
	Page   int    `json:"page" in:"query" d:"1" v:"min:1" dc:"頁碼"`
	Size   int    `json:"size" in:"query" d:"10" v:"between:1,50" dc:"每頁筆數（1-50）"`
}

type SearchArticlesRes struct {
	g.Meta `status:"200" resEg:"api/example/article/v1/articles_search.json"`
	List   []SearchArticlesData `json:"list"`
	Total  int                  `json:"total"`
}
