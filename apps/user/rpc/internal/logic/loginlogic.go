package logic

import (
	"context"
	"fmt"
	"github.com/ljp-lachouchou/chan_xin/apps/user/rpc/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/user/rpc/user"
	"github.com/ljp-lachouchou/chan_xin/pkg/ctxdata"
	"github.com/ljp-lachouchou/chan_xin/pkg/lerr"
	"github.com/ljp-lachouchou/chan_xin/pkg/lhash"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
)

var (
	ErrPhoneNotRegister = lerr.NewError(int(lerr.SERVICE_COMMON_ERROR), "手机号没有被注册")
	ErrUserPwdError     = lerr.NewError(int(lerr.SERVICE_COMMON_ERROR), "密码错误")
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户鉴权
func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	// todo: add your logic here and delete this line
	hasRegister, err := l.svcCtx.UsersModel.FindByPhone(l.ctx, in.Phone)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, errors.WithStack(ErrPhoneNotRegister)
		}
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "user-rpc Login UsersModel.FindByPhone", in.Phone)
	}
	//if _, err := l.svcCtx.Redis.Get(ldefault.SYSTEM_REDIS_UID); err == nil {
	//	return l.LoginResp(hasRegister.Id)
	//}
	if !lhash.ComparePasswords(hasRegister.Password.String, in.Password) {
		return nil, errors.WithStack(ErrUserPwdError)
	}

	return l.LoginResp(hasRegister.Id)
}
func (l *LoginLogic) LoginResp(id string) (*user.LoginResponse, error) {
	iat := time.Now().Unix()
	token, err := ctxdata.GetToken(l.svcCtx.Config.Jwt.AccessSecret, iat, l.svcCtx.Config.Jwt.AccessExpire, id)
	redisKey := fmt.Sprintf("user_token:%s", id)

	if err := l.svcCtx.Redis.Setex(redisKey, token, int(l.svcCtx.Config.Jwt.AccessExpire)); err != nil {
		return nil, lerr.NewWrapError(lerr.NewCOMMONError(), err, "user-rpc SetRootToken")
	}
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "user-rpc Register GetToken")
	}
	return &user.LoginResponse{
		Id:        id,
		AuthToken: token,
		ExpiresAt: iat + l.svcCtx.Config.Jwt.AccessExpire,
	}, nil
}
