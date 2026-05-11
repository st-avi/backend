package utility

import (
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"
)

type DBCfg struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Pass     string `json:"pass"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Timezone string `json:"timezone"`
}

var DBDefaultCfg DBCfg
var JwtSecret []byte
var JwtDomain string
var CORSAllowDomain []string

func init() {
	ctx := gctx.New()

	secret, err := gcfg.Instance().Get(ctx, "jwt.secret")
	if err != nil {
		panic("jwt.secret 讀取失敗: " + err.Error())
	}
	JwtSecret = secret.Bytes()

	jwtDomain, err := gcfg.Instance().Get(ctx, "jwt.domain")
	if err != nil {
		panic("jwt.domain 讀取失敗: " + err.Error())
	}
	JwtDomain = jwtDomain.String()

	dbDefault, err := gcfg.Instance().Get(ctx, "database.default")
	if err != nil {
		panic("database.default 讀取失敗: " + err.Error())
	}
	err = dbDefault.MapToMap(&DBDefaultCfg)
	if err != nil {
		panic("database.default 解析失敗: " + err.Error())
	}

	corsAllowDomain, err := gcfg.Instance().Get(ctx, "cors.allowDomain")
	if err != nil {
		panic("cors.allowDomain 讀取失敗: " + err.Error())
	}
	CORSAllowDomain = corsAllowDomain.Strings()
}
