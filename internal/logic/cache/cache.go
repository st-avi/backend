package cache

import (
	"encoding/json"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

func Set(key string, value any, ttl time.Duration) error {
	b, err := json.Marshal(value)
	if err != nil {
		return err
	}
	expiresAt := gtime.Now().Add(ttl)
	_, err = g.Model("cache").Data(g.Map{
		"key":        key,
		"value":      string(b),
		"expires_at": expiresAt,
	}).OnConflict("key").Save()
	return err
}

func MSet(data g.Map) error {
	_, err := g.Model("cache").Data(g.List{
		data,
	}).OnConflict("key").Save()
	return err
}

func Get(key string) (gdb.Record, error) {
	return g.Model("cache").Where("key", key).Where("expires_at >", gtime.Now()).One()
}

func MGet(key []string) ([]gdb.Record, error) {
	return g.Model("cache").Where("key", key).Where("expires_at >", gtime.Now()).All()
}

func Del(key string) error {
	_, err := g.Model("cache").Where("key", key).Delete()
	return err
}

func Exists(key string) bool {
	exist, err := g.Model("cache").Where("key", key).Where("expires_at >", gtime.Now()).Exist()
	if err != nil {
		return false
	}
	return exist
}

func Expire(key string, ttl time.Duration) error {
	if ttl <= 0 {
		return Del(key)
	}
	_, err := g.Model("cache").Where("key", key).Data(g.Map{
		"expires_at": gtime.Now().Add(ttl),
	}).Update()
	return err
}
