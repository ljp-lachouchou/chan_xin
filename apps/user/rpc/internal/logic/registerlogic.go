package logic

import (
	"context"
	"database/sql"
	"github.com/ljp-lachouchou/chan_xin/apps/user/usermodels"
	"github.com/ljp-lachouchou/chan_xin/pkg/ctxdata"
	"github.com/ljp-lachouchou/chan_xin/pkg/lerr"
	"github.com/ljp-lachouchou/chan_xin/pkg/lhash"
	"github.com/ljp-lachouchou/chan_xin/pkg/wuid"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"time"

	"github.com/ljp-lachouchou/chan_xin/apps/user/rpc/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	PhoneHasRegister = errors.New("手机号已经被注册")
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.RegisterResp, error) {
	// todo: add your logic here and delete this line
	hasRegister, err := l.svcCtx.UsersModel.FindByPhone(l.ctx, in.Phone)
	if err != nil && err != sqlc.ErrNotFound {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "user-rpc Register UsersModel.FindByPhone", in.Phone)
	}
	if hasRegister != nil {
		return nil, errors.WithStack(PhoneHasRegister)
	}
	//手机号没被注册
	hasRegister = &usermodels.Users{
		Id:       wuid.GenUid(l.svcCtx.Config.Mysql.DataSource),
		Avatar:   in.Avatar,
		Nickname: in.Nickname,
		Phone:    in.Phone,
		Sex: sql.NullInt64{
			Int64: int64(in.Sex),
			Valid: true,
		},
	}
	if err := lhash.ValidatePassword(in.Password); err != nil {
		return nil, errors.WithStack(err)
	}
	hashPwd, err := lhash.GenPasswordHash(in.Password)
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NewSYSTEMError(), err, "user-rpc Register lhash.GenPasswordHash", in.Password)
	}
	hasRegister.Password = sql.NullString{
		String: hashPwd,
		Valid:  true,
	}
	if _, err := l.svcCtx.UsersModel.Insert(l.ctx, hasRegister); err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "user-rpc Register UsersModel.Insert")
	}
	iat := time.Now().Unix()
	token, err := ctxdata.GetToken(l.svcCtx.Config.Jwt.AccessSecret, iat, l.svcCtx.Config.Jwt.AccessExpire, hasRegister.Id)
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "user-rpc Register GetToken")
	}
	return &user.RegisterResp{
		Token:  token,
		Expire: iat + l.svcCtx.Config.Jwt.AccessExpire,
	}, nil
}
