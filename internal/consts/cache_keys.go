package consts

const (
	CacheKeyVerifyCreateUser = "verify:create_user:" // 新增使用者時 Email 驗證碼
	CacheAuthRTokenJTI       = "auth:rToken:"        // rToken 白名單 JTI

	CacheArticleList10  = "article:list:10"  // 首頁最新 10 篇文章
	CacheArticleListALL = "article:list:all" // 文章頁面
	CacheArticleTags    = "article:tags"     // 文章標籤
	CacheArticleInfo    = "article:info:"    // 文章資訊
)
