// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Cache is the golang structure of table cache for DAO operations like Where/Data.
type Cache struct {
	g.Meta    `orm:"table:cache, do:true"`
	Key       any         //
	Value     any         //
	ExpiresAt *gtime.Time //
}
