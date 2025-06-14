package user

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/ljp-lachouchou/chan_xin/apps/user/rpc/userservice"
	"github.com/ljp-lachouchou/chan_xin/pkg/lerr"

	"github.com/ljp-lachouchou/chan_xin/apps/user/api/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户登入
func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (*types.LoginResp, error) {
	// todo: add your logic here and delete this line
	var rpcReq userservice.LoginRequest
	err := copier.Copy(&rpcReq, req)
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NewCOMMONError(), err, "user-api login copier.Copy failed")
	}
	rpcResp, err := l.svcCtx.Login(l.ctx, &rpcReq)
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NewCOMMONError(), err, "user-api login failed")
	}
	resp := types.LoginResp{
		Id:     rpcResp.Id,
		Token:  rpcResp.AuthToken,
		Expire: rpcResp.ExpiresAt,
	}
	return &resp, nil
}
