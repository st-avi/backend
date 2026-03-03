// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Users is the golang structure for table users.
type Users struct {
	Id           int         `json:"id"           orm:"id"            description:""` //
	Username     string      `json:"username"     orm:"username"      description:""` //
	Email        string      `json:"email"        orm:"email"         description:""` //
	Password     string      `json:"password"     orm:"password"      description:""` //
	PhoneCountry string      `json:"phoneCountry" orm:"phone_country" description:""` //
	PhoneNumber  string      `json:"phoneNumber"  orm:"phone_number"  description:""` //
	RoleId       int         `json:"roleId"       orm:"role_id"       description:""` //
	Avatar       string      `json:"avatar"       orm:"avatar"        description:""` //
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:""` //
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    description:""` //
	DeletedAt    *gtime.Time `json:"deletedAt"    orm:"deleted_at"    description:""` //
}
