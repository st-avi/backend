package consts

const (
	CacheKeyVerifyCreateUser = "verify:create_user:" // 新增使用者時 Email 驗證碼
	CacheAuthRTokenJTI       = "auth:rToken:"        // rToken 白名單 JTI

	CacheArticleListAside = "article:list:aside" // 文章頁面側邊欄列表
	CacheArticleList10    = "article:list:10"    // 文章頁面最新 10 篇文章
	CacheArticleTags      = "article:tags"       // 文章標籤
	CacheArticleInfo      = "article:info:"      // 文章資訊
)
