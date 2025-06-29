package logic

import (
	"context"
	"database/sql"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/dynamicsmodels"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/dynamics"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/pkg/lerr"
	"github.com/ljp-lachouchou/chan_xin/pkg/wuid"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetCoverLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetCoverLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetCoverLogic {
	return &SetCoverLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 设置个人动态封面（用于个人主页）
func (l *SetCoverLogic) SetCover(in *dynamics.SetCoverRequest) (*dynamics.Empty, error) {
	findOne, err := l.svcCtx.SocialCircleModel.FindByUserId(l.ctx, in.UserId)
	if err != nil {
		if err == dynamicsmodels.ErrNotFound {
			_, err := l.svcCtx.SocialCircleModel.Insert(l.ctx, &dynamicsmodels.SocialCircle{
				Id:     wuid.GenUid(l.svcCtx.Config.Mysql.DataSource),
				UserId: in.UserId,
				CoverUrl: sql.NullString{
					Valid:  true,
					String: in.CoverUrl,
				},
			})
			if err != nil {
				return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "dynamic-rpc SetCover insert failed ")
			}
			return &dynamics.Empty{}, nil
		}
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "dynamic-rpc SetCover FindByUserId failed ", in.UserId)
	}

	findOne.CoverUrl = sql.NullString{
		String: in.CoverUrl,
		Valid:  true,
	}
	err = l.svcCtx.SocialCircleModel.Update(l.ctx, findOne)
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "dynamic-rpc SetCover update failed ", findOne)
	}
	return &dynamics.Empty{}, nil
}
