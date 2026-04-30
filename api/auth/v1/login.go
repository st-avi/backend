package v1

import (
	"backend/api"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/goai"
)

type LoginReq struct {
	g.Meta   `path:"/login" tags:"Auth" sm:"登入" dc:"使用帳密、TOTP 登入" method:"post"`
	Email    string `json:"email" v:"required|email" dc:"電子郵件地址"`
	Password string `json:"password" v:"required" dc:"密碼"`
	Totp     string `json:"totp" v:"required|length:6,6" dc:"TOTP"`
}

type LoginRes struct {
	g.Meta `status:"200" resEg:"api/auth/v1/example/login.json"`
}

func (r LoginRes) EnhanceResponseStatus() (resList map[int]goai.EnhancedStatusType) {
	return map[int]goai.EnhancedStatusType{
		401: {
			Response: struct{}{},
			Examples: []interface{}{
				api.CommonRes{
					Code:    61,
					Message: "信箱/密碼/TOTP錯誤",
					Data:    nil,
				},
			},
		},
	}
}
