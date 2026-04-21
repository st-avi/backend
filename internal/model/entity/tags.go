// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Tags is the golang structure for table tags.
type Tags struct {
	Id        int         `json:"id"        orm:"id"         description:""` //
	Name      string      `json:"name"      orm:"name"       description:""` //
	Slug      string      `json:"slug"      orm:"slug"       description:""` //
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""` //
}
