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

	aToken, rToken, err := user.Login(ctx, req.Email, req.Password, req.Totp)
	if err != nil {
		switch gerror.Code(err) {
		case gcode.CodeNotAuthorized:
			r.Response.Status = http.StatusUnauthorized
		}
		return nil, err
	}

	r.Cookie.SetCookie("aToken", aToken, utility.JwtDomain, "/", 2*time.Hour, ghttp.CookieOptions{
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
