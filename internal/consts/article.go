package consts

type ArticleStatus string

const (
	ArticleDraft     ArticleStatus = "draft"
	ArticlePublished ArticleStatus = "published"
	ArticleTrashed   ArticleStatus = "trashed"
)
