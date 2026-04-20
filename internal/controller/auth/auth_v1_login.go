package auth

import (
	v1 "backend/api/auth/v1"
	"backend/internal/consts"
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
		case consts.CodeUnauthorized:
			r.Response.Status = http.StatusUnauthorized
		case gcode.CodeInternalError:
			r.Response.Status = http.StatusInternalServerError
		}
		return nil, err
	}

	r.Response.Header().Add(
		"Set-Cookie", "aToken="+aToken+"; HttpOnly; Secure; SameSite=Lax; Max-Age=7200; Path=/;",
	)
	r.Response.Header().Add(
		"Set-Cookie", "rToken="+rToken+"; HttpOnly; Secure; SameSite=Lax; Max-Age=604800; Path=/;",
	)

	return &v1.LoginRes{}, nil
}
