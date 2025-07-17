package user

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/user/rpc/userservice"
	"github.com/ljp-lachouchou/chan_xin/pkg/ctxdata"
	"github.com/pkg/errors"

	"github.com/ljp-lachouchou/chan_xin/apps/user/api/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 单个用户查询
func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (*types.UserInfoResp, error) {
	// todo: add your logic here and delete this line
	uid := ctxdata.GetUId(l.ctx)
	rpcReq := userservice.GetUserRequest{
		Id: uid,
	}
	rpcResp, err := l.svcCtx.UserService.GetUser(l.ctx, &rpcReq)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	resp := &types.UserInfoResp{
		Info: types.User{
			Id:       rpcResp.Id,
			Phone:    rpcResp.Phone,
			Nickname: rpcResp.Nickname,
			Sex:      byte(rpcResp.Sex),
			Avatar:   rpcResp.Avatar,
		},
	}
	return resp, nil
}
