// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Cache is the golang structure for table cache.
type Cache struct {
	Key       string      `json:"key"       orm:"key"        description:""` //
	Value     string      `json:"value"     orm:"value"      description:""` //
	ExpiresAt *gtime.Time `json:"expiresAt" orm:"expires_at" description:""` //
}
