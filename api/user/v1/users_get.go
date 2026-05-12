package v1

import (
	"backend/api"
	"net/http"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/goai"
)

type GetUsersResUser struct {
	Id           int    `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	PhoneCountry string `json:"phone_country"`
	PhoneNumber  string `json:"phone_number"`
	Role         string `json:"role"`
	Avatar       string `json:"avatar"`
}

type GetUsersReq struct {
	g.Meta `path:"/users" tags:"Admin" summary:"取得所有使用者資料" method:"get"`
}

type GetUsersRes struct {
	g.Meta `status:"200" resEg:"api/user/v1/example/users_get.json"`
	Users  []GetUsersResUser `json:"users"`
}

func (r GetUsersRes) EnhanceResponseStatus() (resList map[int]goai.EnhancedStatusType) {
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
		http.StatusForbidden: {
			Response: struct{}{},
			Examples: []interface{}{
				api.CommonRes{
					Code:    gcode.CodeNotAuthorized.Code(),
					Message: "需要管理員權限",
					Data:    nil,
				},
			},
		},
	}
}
