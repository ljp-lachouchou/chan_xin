package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/ljp-lachouchou/chan_xin/apps/user/usermodels"
	"github.com/ljp-lachouchou/chan_xin/pkg/lerr"
	"github.com/pkg/errors"

	"github.com/ljp-lachouchou/chan_xin/apps/user/rpc/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserLogic {
	return &FindUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindUserLogic) FindUser(in *user.FindUserReq) (*user.FindUserResp, error) {
	var (
		userEntitys []*usermodels.Users
		err         error
	)
	if in.Phone != "" {
		userEntity, err := l.svcCtx.UsersModel.FindByPhone(l.ctx, in.Phone)
		if err == nil {
			userEntitys = append(userEntitys, userEntity)

		}
	} else if in.Name != "" {
		userEntitys, err = l.svcCtx.UsersModel.ListByName(l.ctx, in.Name)
		if err != nil {
			logx.Error(err)
		}
	} else if len(in.Ids) > 0 {
		userEntitys, err = l.svcCtx.UsersModel.ListByIds(l.ctx, in.Ids)
		if err != nil {
			logx.Error(err)
		}
	}
	if err != nil {
		return nil, errors.Wrapf(lerr.NEWDBError(), "db find err: %v req:%v ,%v, %v", err, in.Phone, in.Name, in.Ids)
	}
	var resp []*user.User
	err = copier.Copy(&resp, userEntitys)
	if err != nil {
		return nil, errors.Wrapf(lerr.NewSYSTEMError(), "copy err: %v", err)
	}
	return &user.FindUserResp{
		User: resp,
	}, nil
	return &user.FindUserResp{}, nil
}
