package auth

import (
	v1 "backend/api/auth/v1"
	"backend/internal/consts"
	"backend/internal/logic/users"
	"backend/utility"
	"context"
	"net/http"
	"strconv"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (c *ControllerV1) Refresh(ctx context.Context, req *v1.RefreshReq) (res *v1.RefreshRes, err error) {
	r := ghttp.RequestFromCtx(ctx)

	claims, err := utility.ParseToken(r.Cookie.Get("rToken").String())
	if err != nil || claims.Purpose != utility.JwtPurposeRefresh {
		r.Response.WriteStatus(http.StatusUnauthorized)
		return nil, gerror.NewCode(consts.CodeUnauthorized, "無效的 token")
	}

	userId, err := strconv.Atoi(claims.Subject)
	if err != nil {
		r.Response.WriteStatus(http.StatusInternalServerError)
		return nil, gerror.NewCode(gcode.CodeInternalError)
	}

	aToken, rToken, err := users.GenAuthToken(userId)
	if err != nil {
		r.Response.WriteStatus(http.StatusInternalServerError)
		return nil, err
	}

	r.Response.Header().Add(
		"Set-Cookie", "aToken="+aToken+"; HttpOnly; Secure; Max-Age=7200; Path=/;",
	)
	r.Response.Header().Add(
		"Set-Cookie", "rToken="+rToken+"; HttpOnly; Secure; Max-Age=604800; Path=/;",
	)

	return &v1.RefreshRes{}, nil
}
