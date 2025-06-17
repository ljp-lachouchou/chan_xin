package logic

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/social"
	"github.com/ljp-lachouchou/chan_xin/apps/social/socialmodels"
	"github.com/ljp-lachouchou/chan_xin/deploy/constant"
	"github.com/ljp-lachouchou/chan_xin/pkg/lerr"
	"github.com/ljp-lachouchou/chan_xin/pkg/wuid"
	"github.com/pkg/errors"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	FriendApplyHasExistErr = errors.New("你对对方已经提交过好友申请")
	FriendIsOkErr          = errors.New("你与对方已经是好友")
)

type ApplyFriendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewApplyFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApplyFriendLogic {
	return &ApplyFriendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// === 好友管理接口 ===
func (l *ApplyFriendLogic) ApplyFriend(in *social.FriendApplyRequest) (*social.FriendApplyResponse, error) {
	hasFriend, err2 := l.svcCtx.FriendRelationModel.FindOneByUserIdFriendId(l.ctx, in.ApplicantId, in.TargetId)
	if err2 != nil && err2 != socialmodels.ErrNotFound {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err2, "social-rpc ApplyFriend", in.ApplicantId, in.TargetId)
	}
	if hasFriend != nil {
		return &social.FriendApplyResponse{}, errors.WithStack(FriendIsOkErr)
	}
	hasFriendApply, err := l.svcCtx.FriendApplyModel.FindByApplicantIdAndTargetId(l.ctx, in.ApplicantId, in.TargetId)
	if err != nil && err != socialmodels.ErrNotFound {
		return nil, lerr.NewWrapError(lerr.NewSYSTEMError(), err, "social-rpc ApplyFriend", in.ApplicantId, in.TargetId)
	}
	if hasFriendApply != nil {
		switch constant.FriendApplyHandle(hasFriendApply.Status) {
		case constant.FailHandleApply:
			hasFriendApply.Status = 0
			if err := l.svcCtx.FriendApplyModel.Update(l.ctx, hasFriendApply); err != nil {
				return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "social-rpc HandleFriendApply", hasFriendApply.ApplyId)
			}
			return &social.FriendApplyResponse{
				ApplyId:   hasFriendApply.ApplyId,
				ApplyTime: time.Now().Unix(),
			}, nil
		default:
		}
		return nil, errors.WithStack(FriendApplyHasExistErr)
	}
	applyId := wuid.GenUid(l.svcCtx.Config.Mysql.DataSource)
	friendApply := socialmodels.FriendApply{
		ApplyId:     applyId,
		ApplicantId: in.ApplicantId,
		TargetId:    in.TargetId,
		GreetMsg:    in.GreetMsg,
		Status:      0,
	}
	_, err = l.svcCtx.FriendApplyModel.Insert(l.ctx, &friendApply)
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "social-rpc ApplyFriend Insert", friendApply)
	}
	return &social.FriendApplyResponse{
		ApplyId:   applyId,
		ApplyTime: time.Now().Unix(),
	}, nil
}
