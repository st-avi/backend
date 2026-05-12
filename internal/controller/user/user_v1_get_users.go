package user

import (
	v1 "backend/api/user/v1"
	"backend/internal/consts"
	"backend/internal/dao"
	"backend/internal/model/entity"
	"context"
	"net/http"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (c *ControllerV1) GetUsers(ctx context.Context, req *v1.GetUsersReq) (res *v1.GetUsersRes, err error) {
	r := ghttp.RequestFromCtx(ctx)

	var users []entity.Users
	err = dao.Users.Ctx(ctx).Scan(&users)
	if err != nil {
		r.Response.Status = http.StatusInternalServerError
		return nil, gerror.NewCode(gcode.CodeInternalError, "取得使用者資料失敗")
	}

	resUsers := make([]v1.GetUsersResUser, 0)
	for _, user := range users {
		resUsers = append(resUsers, v1.GetUsersResUser{
			Id:           user.Id,
			Username:     user.Username,
			Email:        user.Email,
			PhoneCountry: user.PhoneCountry,
			PhoneNumber:  user.PhoneNumber,
			Role:         consts.Role(user.RoleId).String(),
			Avatar:       user.Avatar,
		})
	}

	return &v1.GetUsersRes{
		Users: resUsers,
	}, nil
}
