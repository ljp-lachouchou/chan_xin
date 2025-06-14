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

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户注册
func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (*types.RegisterResp, error) {
	// todo: add your logic here and delete this line
	var rpcReq userservice.RegisterReq
	err := copier.Copy(&rpcReq, req)
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NewSYSTEMError(), err, "user-api register copier.Copy()")
	}
	rpcResp, err := l.svcCtx.UserService.Register(l.ctx, &rpcReq)
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NewCOMMONError(), err, "user-api UserService.Register()")
	}
	resp := &types.RegisterResp{
		Token:  rpcResp.Token,
		Expire: rpcResp.Expire,
	}
	return resp, nil
}
