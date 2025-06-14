package logic

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/user/rpc/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserLogic) UpdateUser(in *user.UpdateUserRequest) (*user.User, error) {
	// todo: add your logic here and delete this line
	//u, err := l.svcCtx.UsersModel.FindOne(l.ctx, in.UserId)
	//if err != nil {
	//	return nil, lerr.NewWrapError(lerr.NEWDBError(),err,"user-rpc UpdateUser FindOne",in.UserId)
	//}
	//
	//l.svcCtx.UsersModel.Update()

	return &user.User{}, nil
}
