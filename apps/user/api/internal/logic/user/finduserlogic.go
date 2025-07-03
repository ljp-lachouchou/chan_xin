package user

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/user/rpc/userservice"

	"github.com/ljp-lachouchou/chan_xin/apps/user/api/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户
func NewFindUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserLogic {
	return &FindUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindUserLogic) FindUser(req *types.FindUserReq) (*types.FindUserResp, error) {
	user, err := l.svcCtx.FindUser(l.ctx, &userservice.FindUserReq{
		Name:  req.Name,
		Phone: req.Phone,
		Ids:   req.Ids,
	})
	if err != nil {
		return nil, err
	}
	var resp []types.User
	for _,v := range user.User {
		resp = append(resp,types.User{
			Id:       v.Id,
			Phone:    v.Phone,
			Nickname: v.Nickname,
			Sex:      byte(v.Sex),
			Avatar:   v.Avatar,
		})
	}
	return &types.FindUserResp{
		Infos: resp,
	},nil
}
