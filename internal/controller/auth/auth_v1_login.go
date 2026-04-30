package auth

import (
	v1 "backend/api/auth/v1"
	"backend/internal/logic/users"
	"context"
	"net/http"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	r := ghttp.RequestFromCtx(ctx)

	aToken, rToken, err := users.Login(ctx, req.Email, req.Password, req.Totp)
	if err != nil {
		switch gerror.Code(err) {
		case gcode.CodeNotAuthorized:
			r.Response.Status = http.StatusUnauthorized
		}
		return nil, err
	}

	r.Response.Header().Add(
		"Set-Cookie", "aToken="+aToken+"; HttpOnly; Secure; Max-Age=7200; Path=/; Domain=stavi.tw;",
	)
	r.Response.Header().Add(
		"Set-Cookie", "rToken="+rToken+"; HttpOnly; Secure; Max-Age=604800; Path=/; Domain=stavi.tw;",
	)

	return &v1.LoginRes{}, nil
}
