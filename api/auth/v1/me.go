package v1

import (
	"backend/api"
	"net/http"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/goai"
)

type AuthMeReq struct {
	g.Meta `path:"/auth/me" tags:"Auth" sm:"取得登入者的基本資料" dc:"用於前端 middleware 取得登入者的基本資料" method:"get"`
	AToken string `json:"aToken" in:"cookie" v:"required" dc:"存取用 Token"`
}

type AuthMeRes struct {
	g.Meta   `status:"200" resEg:"api/auth/v1/example/me.json"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Avatar   string `json:"avatar"`
}

func (r AuthMeRes) EnhanceResponseStatus() (resList map[int]goai.EnhancedStatusType) {
	return map[int]goai.EnhancedStatusType{
		http.StatusUnauthorized: {
			Response: struct{}{},
			Examples: []interface{}{
				api.CommonRes{
					Code:    gcode.CodeNotAuthorized.Code(),
					Message: "無效的 token",
					Data:    nil,
				},
			},
		},
	}
}
