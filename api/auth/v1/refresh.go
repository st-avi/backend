package v1

import (
	"backend/api"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/goai"
)

type RefreshReq struct {
	g.Meta `path:"/refresh" tags:"Auth" sm:"刷新 Token" dc:"使用 rToken 來獲取新的 aToken、rToken" method:"post"`
	RToken string `json:"rToken" in:"cookie" v:"required" dc:"刷新用 Token"`
}

type RefreshRes struct {
	g.Meta `status:"200" resEg:"api/auth/v1/example/refresh.json"`
}

func (r RefreshRes) EnhanceResponseStatus() (resList map[int]goai.EnhancedStatusType) {
	return map[int]goai.EnhancedStatusType{
		401: {
			Response: struct{}{},
			Examples: []interface{}{
				api.CommonRes{
					Code:    61,
					Message: "無效的 token",
					Data:    nil,
				},
			},
		},
	}
}
