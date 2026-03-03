package cmd

import (
	"backend/utility"
	"net/http"

	"github.com/gogf/gf/v2/net/ghttp"
)

func MiddlewareCORS(r *ghttp.Request) {
	corsOptions := r.Response.DefaultCORSOptions()
	corsOptions.AllowDomain = utility.CORSAllowDomain
	if !r.Response.CORSAllowedOrigin(corsOptions) {
		r.Response.WriteStatusExit(http.StatusForbidden)
	}
	r.Response.CORS(corsOptions)
	r.Middleware.Next()
}

func MiddlewareAuth(r *ghttp.Request) {
	aToken := r.Cookie.Get("aToken").String()
	claims, err := utility.ParseToken(aToken)
	if err != nil || claims.Purpose != utility.JwtPurposeAccess {
		r.Response.WriteStatusExit(http.StatusUnauthorized, "無效的 token")
	}
	r.SetCtxVar("userId", claims.Subject)
	r.Middleware.Next()
}
