// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Articles is the golang structure of table articles for DAO operations like Where/Data.
type Articles struct {
	g.Meta       `orm:"table:articles, do:true"`
	Id           any         //
	CategoryId   any         //
	Title        any         //
	Slug         any         //
	Summary      any         //
	Content      any         //
	CoverImage   any         //
	Status       any         //
	PublishedAt  *gtime.Time //
	CreatedAt    *gtime.Time //
	UpdatedAt    *gtime.Time //
	SearchVector any         //
}
