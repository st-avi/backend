package auth

import (
	v1 "backend/api/auth/v1"
	"backend/internal/consts"
	"backend/internal/logic/user"
	"backend/utility"
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (c *ControllerV1) Refresh(ctx context.Context, req *v1.RefreshReq) (res *v1.RefreshRes, err error) {
	r := ghttp.RequestFromCtx(ctx)

	claims, err := utility.ParseToken(r.Cookie.Get("rToken").String())
	if err != nil || claims.Purpose != utility.JwtPurposeRefresh {
		r.Cookie.RemoveCookie("rToken", utility.JwtDomain, "/")
		r.Response.Status = http.StatusUnauthorized
		return nil, gerror.NewCode(gcode.CodeNotAuthorized, "無效的 token")
	}

	// 檢查白名單
	exists, err := g.Redis().Exists(ctx, consts.CacheAuthRTokenJTI+claims.ID)
	if err != nil || exists == 0 {
		r.Cookie.RemoveCookie("rToken", utility.JwtDomain, "/")
		r.Response.Status = http.StatusUnauthorized
		return nil, gerror.NewCode(gcode.CodeNotAuthorized, "無效的 token")
	}

	userId, err := strconv.Atoi(claims.Subject)
	if err != nil {
		r.Response.Status = http.StatusInternalServerError
		return nil, gerror.NewCode(gcode.CodeInternalError)
	}

	aToken, rToken, err := user.GenAuthToken(userId)
	if err != nil {
		r.Response.Status = http.StatusInternalServerError
		return nil, err
	}

	// 刪除舊的白名單
	_, err = g.Redis().Del(ctx, consts.CacheAuthRTokenJTI+claims.ID)
	if err != nil {
		r.Response.Status = http.StatusInternalServerError
		return nil, gerror.NewCode(gcode.CodeInternalError)
	}

	// 新增新的白名單
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

	return &v1.RefreshRes{}, nil
}
