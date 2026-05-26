package auth

import (
	v1 "backend/api/auth/v1"
	"backend/internal/consts"
	"backend/utility"
	"context"
	"net/http"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (c *ControllerV1) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	r := ghttp.RequestFromCtx(ctx)

	rToken := r.Cookie.Get("rToken").String()
	if rToken != "" {
		// 刪除白名單
		rClaims, err := utility.ParseToken(rToken)
		if err != nil {
			r.Response.Status = http.StatusInternalServerError
			return nil, gerror.NewCode(gcode.CodeInternalError)
		}

		_, err = g.Redis().Del(ctx, consts.CacheAuthRTokenJTI+rClaims.ID)
		if err != nil {
			r.Response.Status = http.StatusInternalServerError
			return nil, gerror.NewCode(gcode.CodeInternalError)
		}
	}

	r.Cookie.RemoveCookie("aToken", utility.JwtDomain, "/")
	r.Cookie.RemoveCookie("rToken", utility.JwtDomain, "/")

	return &v1.LogoutRes{}, nil
}
