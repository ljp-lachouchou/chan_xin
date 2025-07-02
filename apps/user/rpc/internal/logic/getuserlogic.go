package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/ljp-lachouchou/chan_xin/apps/user/rpc/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/user/rpc/user"
	"github.com/ljp-lachouchou/chan_xin/pkg/lerr"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	ErrUserNotFound = lerr.NewError(int(lerr.SERVICE_COMMON_ERROR), "所查用户没有被找到")
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户资料
func (l *GetUserLogic) GetUser(in *user.GetUserRequest) (*user.User, error) {
	// todo: add your logic here and delete this line
	u, err := l.svcCtx.UsersModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, errors.WithStack(ErrUserNotFound)
	}
	var resp user.User
	err = copier.Copy(&resp, u)
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NewSYSTEMError(), err, "user-rpc GerUser copier.Copy")
	}
	return &resp, nil
}
