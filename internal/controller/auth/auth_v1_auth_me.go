package auth

import (
	v1 "backend/api/auth/v1"
	"backend/internal/consts"
	"backend/internal/dao"
	"backend/internal/model/entity"
	"context"
	"net/http"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (c *ControllerV1) AuthMe(ctx context.Context, req *v1.AuthMeReq) (res *v1.AuthMeRes, err error) {
	r := ghttp.RequestFromCtx(ctx)

	var user entity.Users
	err = dao.Users.Ctx(ctx).Where("id", r.GetCtxVar("userId").Int()).Scan(&user)
	if err != nil {
		r.Response.Status = http.StatusInternalServerError
		return nil, gerror.NewCode(gcode.CodeInternalError)
	}

	return &v1.AuthMeRes{
		Username: user.Username,
		Email:    user.Email,
		Role:     consts.Role(user.RoleId).String(),
		Avatar:   user.Avatar,
	}, nil
}
