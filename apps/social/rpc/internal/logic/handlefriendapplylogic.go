package logic

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/social"
	"github.com/ljp-lachouchou/chan_xin/apps/social/socialmodels"
	"github.com/ljp-lachouchou/chan_xin/deploy/constant"
	"github.com/ljp-lachouchou/chan_xin/pkg/lerr"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	ApplyHasPassErr   = errors.New("此申请已被对方通过")
	ApplyHasRefuseErr = errors.New("此申请已被对方拒绝")
)

type HandleFriendApplyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHandleFriendApplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HandleFriendApplyLogic {
	return &HandleFriendApplyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HandleFriendApplyLogic) HandleFriendApply(in *social.FriendApplyAction) (*social.FriendApplyActionResp, error) {
	findOne, err := l.svcCtx.FriendApplyModel.FindOne(l.ctx, in.ApplyId)
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "social-rpc HandleFriendApply", in.ApplyId)
	}
	switch constant.FriendApplyHandle(findOne.Status) {
	case constant.SuccessHandleApply:
		return nil, errors.WithStack(ApplyHasPassErr)
	case constant.FailHandleApply:
		return nil, errors.WithStack(ApplyHasRefuseErr)
	}
	if in.IsApproved {
		findOne.Status = 1
	} else {
		findOne.Status = 2
	}
	err = l.svcCtx.FriendApplyModel.Tranx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		if err := l.svcCtx.FriendApplyModel.UpdateWithSession(l.ctx, session, findOne); err != nil {
			return lerr.NewWrapError(lerr.NEWDBError(), err, "social-rpc HandleFriendApply Tranx FriendApplyModel.Update", findOne)
		}
		if constant.FriendApplyHandle(findOne.Status) != constant.SuccessHandleApply {
			return nil
		}
		friends := []*socialmodels.FriendRelation{
			{
				UserId:    findOne.ApplicantId,
				FriendId:  findOne.TargetId,
				Remark:    "",
				IsMuted:   0,
				IsTopped:  0,
				IsBlocked: 0,
			},
			{
				UserId:    findOne.TargetId,
				FriendId:  findOne.ApplicantId,
				Remark:    "",
				IsMuted:   0,
				IsTopped:  0,
				IsBlocked: 0,
			},
		}
		_, err2 := l.svcCtx.FriendRelationModel.Insert(l.ctx, session, friends...)
		return err2
	})
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "social-rpc HandleFriendApply Tranx")
	}
	return &social.FriendApplyActionResp{
		IsApproved:  constant.FriendApplyHandle(findOne.Status) == constant.SuccessHandleApply,
		ApplicantId: findOne.ApplicantId,
		TargetId:    findOne.TargetId,
	}, nil
}
