package cmd

import (
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
		r.Response.Status = http.StatusUnauthorized
		r.Response.WriteJsonExit(ghttp.DefaultHandlerResponse{
			Code:    gcode.CodeNotAuthorized.Code(),
			Message: "無效的 token",
		})
	}
	r.SetCtxVar("userId", claims.Subject)
	r.Middleware.Next()
}
