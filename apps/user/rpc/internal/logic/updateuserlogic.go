package logic

import (
	"context"
	"database/sql"
	"github.com/ljp-lachouchou/chan_xin/apps/user/rpc/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/user/rpc/user"
	"github.com/ljp-lachouchou/chan_xin/pkg/lerr"
	"time"

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
	u, err := l.svcCtx.UsersModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "user-rpc UpdateUser FindOne", in.Id)
	}
	if in.Sex != nil {
		u.Sex = sql.NullInt64{
			Int64: int64(*in.Sex),
			Valid: true,
		}
	}
	if in.Sex != nil || in.Nickname != nil || in.Avatar != nil {
		u.UpdatedAt = sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		}
	}
	if in.Avatar != nil {
		u.Avatar = *in.Avatar
	}
	if in.Nickname != nil {
		u.Nickname = *in.Nickname
	}
	if err := l.svcCtx.UsersModel.Update(l.ctx, u); err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "user-rpc UpdateUser update")
	}
	return &user.User{
		Id:       u.Id,
		Nickname: u.Nickname,
		Avatar:   u.Avatar,
		Phone:    u.Phone,
		Status:   int32(u.Status.Int64),
		Sex:      int32(u.Sex.Int64),
	}, nil
}
