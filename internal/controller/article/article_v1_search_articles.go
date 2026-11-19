package article

import (
	v1 "backend/api/article/v1"
	"backend/internal/consts"
	"backend/internal/dao"
	"backend/utility"
	"context"
	"net/http"
	"strings"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
)

const searchMaxKeywords = 5

var likeEscaper = strings.NewReplacer(`\`, `\\`, `%`, `\%`, `_`, `\_`)

func (c *ControllerV1) SearchArticles(ctx context.Context, req *v1.SearchArticlesReq) (res *v1.SearchArticlesRes, err error) {
	r := ghttp.RequestFromCtx(ctx)

	var patterns []string
	seen := make(map[string]struct{})
	for _, kw := range strings.Fields(req.Q) {
		if _, ok := seen[kw]; ok {
			continue
		}
		seen[kw] = struct{}{}
		patterns = append(patterns, "%"+likeEscaper.Replace(kw)+"%")
		if len(patterns) == searchMaxKeywords {
			break
		}
	}
	if len(patterns) == 0 {
		return &v1.SearchArticlesRes{List: []v1.SearchArticlesData{}}, nil
	}

	// 每個關鍵字都需出現在標題、摘要或內文其中之一
	// 標題命中 3 分、摘要 2 分、僅內文 1 分，加總後排序
	var (
		scores    = make([]string, 0, len(patterns))
		conds     = make([]string, 0, len(patterns))
		scoreArgs = make([]any, 0, len(patterns)*2)
		condArgs  = make([]any, 0, len(patterns)*3)
	)
	for _, p := range patterns {
		scores = append(scores, "CASE WHEN a.title ILIKE ? THEN 3 WHEN a.summary ILIKE ? THEN 2 ELSE 1 END")
		scoreArgs = append(scoreArgs, p, p)
		conds = append(conds, "(a.title ILIKE ? OR a.summary ILIKE ? OR a.content ILIKE ?)")
		condArgs = append(condArgs, p, p, p)
	}
	sql := `SELECT a.title, a.slug, a.summary, a.cover_image, a.published_at,
		c.name AS category, c.slug AS category_slug,
		(` + strings.Join(scores, " + ") + `) AS score
	FROM articles a
	LEFT JOIN categories c ON c.id = a.category_id
	WHERE a.status = ? AND ` + strings.Join(conds, " AND ") + `
	ORDER BY score DESC, a.published_at DESC, a.id DESC`
	args := append(scoreArgs, consts.ArticlePublished)
	args = append(args, condArgs...)

	articles, total, err := dao.Articles.Ctx(ctx).Raw(sql, args...).
		Page(req.Page, req.Size).
		AllAndCount(false)
	if err != nil {
		r.Response.Status = http.StatusInternalServerError
		return nil, gerror.NewCode(gcode.CodeInternalError, "搜尋文章失敗")
	}

	list := make([]v1.SearchArticlesData, 0, len(articles))
	for _, article := range articles {
		list = append(list, v1.SearchArticlesData{
			Category:     article["category"].String(),
			CategorySlug: article["category_slug"].String(),
			Title:        article["title"].String(),
			Slug:         article["slug"].String(),
			Summary:      article["summary"].String(),
			CoverImage:   utility.S3FileUrl(article["cover_image"].String()),
			PublishedAt:  article["published_at"].Time().Format("2006-01-02 15:04:05"),
		})
	}

	return &v1.SearchArticlesRes{
		List:  list,
		Total: total,
	}, nil
}
