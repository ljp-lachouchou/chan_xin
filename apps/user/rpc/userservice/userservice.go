// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: user.proto

package userservice

import (
	"context"

	"github.com/ljp-lachouchou/chan_xin/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	FindUserReq       = user.FindUserReq
	FindUserResp      = user.FindUserResp
	GetUserRequest    = user.GetUserRequest
	LoginRequest      = user.LoginRequest
	LoginResponse     = user.LoginResponse
	PingReq           = user.PingReq
	PingResp          = user.PingResp
	RegisterReq       = user.RegisterReq
	RegisterResp      = user.RegisterResp
	UpdateUserRequest = user.UpdateUserRequest
	User              = user.User

	UserService interface {
		// 用户鉴权
		Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
		Ping(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingResp, error)
		// 用户资料
		GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error)
		UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error)
		Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error)
		FindUser(ctx context.Context, in *FindUserReq, opts ...grpc.CallOption) (*FindUserResp, error)
	}

	defaultUserService struct {
		cli zrpc.Client
	}
)

func NewUserService(cli zrpc.Client) UserService {
	return &defaultUserService{
		cli: cli,
	}
}

// 用户鉴权
func (m *defaultUserService) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultUserService) Ping(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingResp, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}

// 用户资料
func (m *defaultUserService) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.GetUser(ctx, in, opts...)
}

func (m *defaultUserService) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.UpdateUser(ctx, in, opts...)
}

func (m *defaultUserService) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultUserService) FindUser(ctx context.Context, in *FindUserReq, opts ...grpc.CallOption) (*FindUserResp, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.FindUser(ctx, in, opts...)
}
