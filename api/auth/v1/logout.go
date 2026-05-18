package v1

import (
	"backend/api"
	"net/http"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/goai"
)

type LogoutReq struct {
	g.Meta `path:"/logout" tags:"Auth" sm:"登出" dc:"登出使用者" method:"post"`
}

type LogoutRes struct {
	g.Meta `status:"200" resEg:"api/auth/v1/example/logout.json"`
}

func (r LogoutRes) EnhanceResponseStatus() (resList map[int]goai.EnhancedStatusType) {
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
