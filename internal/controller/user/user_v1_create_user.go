package user

import (
	v1 "backend/api/user/v1"
	"backend/internal/consts"
	"backend/internal/dao"
	"backend/internal/logic/cache"
	"backend/internal/logic/user"
	"context"
	"net/http"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	disposable "github.com/rocketlaunchr/anti-disposable-email"
)

func (c *ControllerV1) CreateUser(ctx context.Context, req *v1.CreateUserReq) (res *v1.CreateUserRes, err error) {
	r := ghttp.RequestFromCtx(ctx)

	// 檢查 Email 避免一次性信箱
	parsedEmail, _ := disposable.ParseEmail(req.Email)
	if parsedEmail.Disposable {
		r.Response.Status = http.StatusUnprocessableEntity
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "不允許使用一次性信箱")
	}

	// 檢查 Email 是否已經被註冊
	exist, err := dao.Users.Ctx(ctx).Where("email", req.Email).Exist()
	if err != nil {
		r.Response.Status = http.StatusInternalServerError
		return nil, gerror.NewCode(gcode.CodeInternalError)
	}
	if exist {
		r.Response.Status = http.StatusConflict
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "Email 已經被註冊")
	}

	// 檢查 Email 驗證碼
	verifyCacheKey := consts.CacheKeyVerifyCreateUser + req.Email
	verifyCache, err := cache.Get(verifyCacheKey)
	if err != nil || verifyCache.IsEmpty() {
		r.Response.Status = http.StatusUnprocessableEntity
		return nil, gerror.NewCode(gcode.CodeValidationFailed, "Email 驗證碼已過期或不存在")
	}
	if req.VerifyEmailCode != verifyCache["value"].String() {
		r.Response.Status = http.StatusUnprocessableEntity
		return nil, gerror.NewCode(gcode.CodeValidationFailed, "Email 驗證碼錯誤")
	}

	// 新增使用者
	_, err = user.CreateUser(ctx, *req)
	if err != nil {
		r.Response.Status = http.StatusInternalServerError
		return nil, gerror.NewCode(gcode.CodeInternalError, "新增使用者失敗")
	}

	// 清理 Email 驗證碼
	_ = cache.Del(verifyCacheKey)

	return &v1.CreateUserRes{}, nil
}
