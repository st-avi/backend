package v1

import (
	"backend/api"
	"net/http"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/goai"
)

type CreateUserReq struct {
	g.Meta          `path:"/user" tags:"User" summary:"新增使用者" method:"post"`
	Username        string `json:"username" v:"required|min-length:3" dc:"使用者名稱"`
	Email           string `json:"email" v:"required|email" dc:"電子郵件地址"`
	VerifyEmailCode string `json:"verify_email_code" v:"required|length:6,6" dc:"電子信箱驗證碼"`
	Password        string `json:"password" v:"required|min-length:12|password3" dc:"使用者密碼：長度 12~18、必須包含大小寫字母、數字和特殊符號"`
	PhoneCountry    string `json:"phone_country" v:"length:2,5" dc:"電話國碼"`
	PhoneNumber     string `json:"phone_number" v:"length:1,20" dc:"電話號碼"`
}

type CreateUserRes struct {
	g.Meta `status:"200" resEg:"api/user/v1/example/user_create.json"`
}

func (r CreateUserRes) EnhanceResponseStatus() (resList map[int]goai.EnhancedStatusType) {
	return map[int]goai.EnhancedStatusType{
		http.StatusConflict: {
			Response: struct{}{},
			Examples: []interface{}{
				api.CommonRes{
					Code:    gcode.CodeBusinessValidationFailed.Code(),
					Message: "Email 已經被註冊",
					Data:    nil,
				},
			},
		},
		http.StatusUnprocessableEntity: {
			Response: struct{}{},
			Examples: []interface{}{
				api.CommonRes{
					Code:    gcode.CodeBusinessValidationFailed.Code(),
					Message: "不允許使用一次性信箱",
					Data:    nil,
				},
				api.CommonRes{
					Code:    gcode.CodeValidationFailed.Code(),
					Message: "Email 驗證碼已過期或不存在",
					Data:    nil,
				},
				api.CommonRes{
					Code:    gcode.CodeValidationFailed.Code(),
					Message: "Email 驗證碼錯誤",
					Data:    nil,
				},
			},
		},
	}
}
