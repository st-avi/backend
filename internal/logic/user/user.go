package user

import (
	v1 "backend/api/user/v1"
	"backend/internal/consts"
	"backend/internal/dao"
	"backend/utility"
	"context"
	"database/sql"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func CreateUser(ctx context.Context, req v1.CreateUserReq) (result sql.Result, err error) {
	hashPWD, err := utility.HashPWD(req.Password)
	if err != nil {
		return nil, gerror.NewCode(gcode.CodeInternalError, "密碼 Hash 失敗")
	}
	return dao.Users.Ctx(ctx).Insert(g.Map{
		"username":      req.Username,
		"email":         req.Email,
		"password":      hashPWD,
		"phone_country": req.PhoneCountry,
		"phone_number":  req.PhoneNumber,
		"role_id":       consts.RoleUser,
	})
}
