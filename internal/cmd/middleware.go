package cmd

import (
	"backend/internal/logic/user"
	"backend/utility"
	"net/http"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/net/ghttp"
)

func MiddlewareCORS(r *ghttp.Request) {
	corsOptions := r.Response.DefaultCORSOptions()
	corsOptions.AllowDomain = utility.CORSAllowDomain
	if !r.Response.CORSAllowedOrigin(corsOptions) {
		r.Response.Status = http.StatusForbidden
		r.Response.WriteJsonExit(ghttp.DefaultHandlerResponse{
			Code:    gcode.CodeNotAuthorized.Code(),
			Message: "CORS origin not allowed",
		})
	}
	r.Response.CORS(corsOptions)
	r.Middleware.Next()
}

func MiddlewareAuth(r *ghttp.Request) {
	aToken := r.Cookie.Get("aToken").String()
	claims, err := utility.ParseToken(aToken)
	if err != nil || claims.Purpose != utility.JwtPurposeAccess {
		r.Cookie.Remove("aToken")
		r.Response.Status = http.StatusUnauthorized
		r.Response.WriteJsonExit(ghttp.DefaultHandlerResponse{
			Code:    gcode.CodeNotAuthorized.Code(),
			Message: "無效的 token",
		})
	}
	r.SetCtxVar("userId", claims.Subject)
	r.Middleware.Next()
}

func MiddlewareAdmin(r *ghttp.Request) {
	userId := r.GetCtxVar("userId").Int()
	isAdmin, err := user.IsAdmin(userId)
	if err != nil || !isAdmin {
		r.Response.Status = http.StatusForbidden
		r.Response.WriteJsonExit(ghttp.DefaultHandlerResponse{
			Code:    gcode.CodeNotAuthorized.Code(),
			Message: "需要管理員權限",
		})
	}
	r.Middleware.Next()
}
