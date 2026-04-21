// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Articles is the golang structure for table articles.
type Articles struct {
	Id           int         `json:"id"           orm:"id"            description:""` //
	CategoryId   int         `json:"categoryId"   orm:"category_id"   description:""` //
	Title        string      `json:"title"        orm:"title"         description:""` //
	Slug         string      `json:"slug"         orm:"slug"          description:""` //
	Summary      string      `json:"summary"      orm:"summary"       description:""` //
	Content      string      `json:"content"      orm:"content"       description:""` //
	CoverImage   string      `json:"coverImage"   orm:"cover_image"   description:""` //
	Status       string      `json:"status"       orm:"status"        description:""` //
	PublishedAt  *gtime.Time `json:"publishedAt"  orm:"published_at"  description:""` //
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:""` //
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    description:""` //
	SearchVector string      `json:"searchVector" orm:"search_vector" description:""` //
}
