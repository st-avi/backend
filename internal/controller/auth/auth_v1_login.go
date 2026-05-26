package auth

import (
	v1 "backend/api/auth/v1"
	"backend/internal/logic/user"
	"backend/utility"
	"context"
	"net/http"
	"time"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	r := ghttp.RequestFromCtx(ctx)

	// 登入、生成 Token
	aToken, rToken, err := user.Login(ctx, req.Email, req.Password, req.Totp)
	if err != nil {
		switch gerror.Code(err) {
		case gcode.CodeNotAuthorized:
			r.Response.Status = http.StatusUnauthorized
		}
		return nil, err
	}

	// 新增白名單
	err = user.SetCacheAuthRTokenJTI(ctx, rToken)
	if err != nil {
		r.Response.Status = http.StatusInternalServerError
		return nil, gerror.NewCode(gcode.CodeInternalError)
	}

	r.Cookie.SetCookie("aToken", aToken, utility.JwtDomain, "/", 1*time.Hour, ghttp.CookieOptions{
		SameSite: http.SameSiteStrictMode,
		Secure:   true,
		HttpOnly: true,
	})
	r.Cookie.SetCookie("rToken", rToken, utility.JwtDomain, "/", 7*24*time.Hour, ghttp.CookieOptions{
		SameSite: http.SameSiteStrictMode,
		Secure:   true,
		HttpOnly: true,
	})

	return &v1.LoginRes{}, nil
}
