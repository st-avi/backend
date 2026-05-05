package jobs

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

func CleanCacheTable() error {
	_, err := g.Model("cache").Where("expires_at <", gtime.Now()).Delete()
	return err
}
