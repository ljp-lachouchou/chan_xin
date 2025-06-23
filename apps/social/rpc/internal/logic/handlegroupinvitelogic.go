package logic

import (
	"context"
	"github.com/go-sql-driver/mysql"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/social"
	"github.com/ljp-lachouchou/chan_xin/apps/social/socialmodels"
	"github.com/ljp-lachouchou/chan_xin/deploy/constant"
	"github.com/ljp-lachouchou/chan_xin/pkg/lerr"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/zeromicro/go-zero/core/logx"
)

type HandleGroupInviteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHandleGroupInviteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HandleGroupInviteLogic {
	return &HandleGroupInviteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HandleGroupInviteLogic) HandleGroupInvite(in *social.GroupInviteAction) (*social.GroupInviteActionResp, error) {
	findOne, err := l.svcCtx.GroupApplyModel.FindOne(l.ctx, in.InviteId)
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "social-rpc HandleGroupInvite", in.InviteId)
	}
	switch constant.FriendApplyHandle(findOne.Status) {
	case constant.SuccessHandleApply:
		return nil, errors.WithStack(ApplyHasPassErr)
	case constant.FailHandleApply:
		return nil, errors.WithStack(ApplyHasRefuseErr)
	}
	if in.IsAccepted {
		findOne.Status = 1
	} else {
		findOne.Status = 2
	}
	err = l.svcCtx.GroupApplyModel.Transx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		if err := l.svcCtx.GroupApplyModel.Update(l.ctx, session, findOne); err != nil {
			return lerr.NewWrapError(lerr.NEWDBError(), err, "social-rpc HandleGroupInvite Tranx GroupApplyModel.Update", findOne)
		}
		if constant.FriendApplyHandle(findOne.Status) != constant.SuccessHandleApply {
			return nil
		}

		groupName, _ := l.svcCtx.GroupInfoModel.GetNicknameByGid(l.ctx, findOne.ApplicantId)
		member := []*socialmodels.GroupMember{
			{
				GroupId:       findOne.ApplicantId,
				UserId:        findOne.TargetId,
				GroupNickname: groupName,
				ShowNickname:  1,
				IsAdmin:       0,
				IsMuted:       0,
				IsTopped:      0,
				Remark:        groupName,
			},
		}
		_, err := l.svcCtx.GroupMemberModel.InsertMembers(l.ctx, session, member...)
		return err
	})
	if err != nil {
		err2 := errors.Cause(err)
		if v, ok := err2.(*mysql.MySQLError); ok {
			switch v.Number {
			case 1062:
				return nil, errors.WithStack(errors.New("用户ID为" + findOne.TargetId + "的用户已经在本群中"))
			case 1452:
				return nil, errors.WithStack(errors.New("该群已被解散"))
			}
		}
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "social-rpc HandleGroupInvite Tranx")
	}
	return &social.GroupInviteActionResp{
		IsApproved:    in.IsAccepted,
		ApplicationId: findOne.TargetId,
		TargetId:      findOne.ApplicantId,
	}, nil
}
