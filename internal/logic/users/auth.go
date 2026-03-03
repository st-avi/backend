package users

import (
	"backend/internal/consts"
	"backend/internal/dao"
	"backend/internal/model/entity"
	"backend/utility"
	"context"
	"time"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

func Login(ctx context.Context, email, password string) (aToken, rToken string, err error) {
	var user entity.Users
	err = dao.Users.Ctx(ctx).Where("email", email).Scan(&user)
	if err != nil {
		return "", "", gerror.NewCode(consts.CodeUnauthorized, "信箱或密碼錯誤")
	}

	pwdCorrect, err := utility.ComparePWD(password, user.Password)
	if err != nil || !pwdCorrect {
		return "", "", gerror.NewCode(consts.CodeUnauthorized, "信箱或密碼錯誤")
	}

	return GenAuthToken(user.Id)
}

func GenAuthToken(userId int) (aToken, rToken string, err error) {
	aToken, aErr := utility.GenToken(userId, utility.JwtPurposeAccess, time.Hour*2)
	rToken, rErr := utility.GenToken(userId, utility.JwtPurposeRefresh, time.Hour*24*7)
	if aErr != nil || rErr != nil {
		return "", "", gerror.NewCode(gcode.CodeInternalError, "Token 生成失敗")
	}
	return aToken, rToken, nil
}
