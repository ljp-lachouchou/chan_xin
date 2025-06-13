package logic

import (
	"context"

	"github.com/ljp-lachouchou/chan_xin/apps/user/rpc/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
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

	return &user.User{}, nil
}
