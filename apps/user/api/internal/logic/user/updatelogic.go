package user

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/ljp-lachouchou/chan_xin/apps/user/rpc/userservice"
	"github.com/ljp-lachouchou/chan_xin/pkg/ctxdata"
	"github.com/ljp-lachouchou/chan_xin/pkg/lerr"
	"github.com/pkg/errors"

	"github.com/ljp-lachouchou/chan_xin/apps/user/api/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新用户
func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.UpdateReq) (resp *types.UpdateResp, err error) {
	// todo: add your logic here and delete this line
	var rpcReq userservice.UpdateUserRequest
	err = copier.Copy(&rpcReq, req)
	uid := ctxdata.GetUId(l.ctx)
	rpcReq.Id = uid
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NewSYSTEMError(), err, "user-api update copier.Copy()")
	}
	rpcResp, err := l.svcCtx.UserService.UpdateUser(l.ctx, &rpcReq)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &types.UpdateResp{
		Info: types.User{
			Id:       rpcResp.Id,
			Phone:    rpcResp.Phone,
			Nickname: rpcResp.Nickname,
			Sex:      byte(rpcResp.Sex),
			Avatar:   rpcResp.Avatar,
		},
	}, nil
}
