package auth

import (
	v1 "backend/api/auth/v1"
	"backend/utility"
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (c *ControllerV1) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	r := ghttp.RequestFromCtx(ctx)

	r.Cookie.RemoveCookie("aToken", utility.JwtDomain, "/")
	r.Cookie.RemoveCookie("rToken", utility.JwtDomain, "/")

	return &v1.LogoutRes{}, nil
}
