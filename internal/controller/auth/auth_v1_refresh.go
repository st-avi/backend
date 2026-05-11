package auth

import (
	v1 "backend/api/auth/v1"
	"backend/internal/logic/user"
	"backend/utility"
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (c *ControllerV1) Refresh(ctx context.Context, req *v1.RefreshReq) (res *v1.RefreshRes, err error) {
	r := ghttp.RequestFromCtx(ctx)

	claims, err := utility.ParseToken(r.Cookie.Get("rToken").String())
	if err != nil || claims.Purpose != utility.JwtPurposeRefresh {
		r.Cookie.Remove("rToken")
		r.Response.WriteStatus(http.StatusUnauthorized)
		return nil, gerror.NewCode(gcode.CodeNotAuthorized, "無效的 token")
	}

	userId, err := strconv.Atoi(claims.Subject)
	if err != nil {
		r.Response.WriteStatus(http.StatusInternalServerError)
		return nil, gerror.NewCode(gcode.CodeInternalError)
	}

	aToken, rToken, err := user.GenAuthToken(userId)
	if err != nil {
		r.Response.WriteStatus(http.StatusInternalServerError)
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

	return &v1.RefreshRes{}, nil
}
