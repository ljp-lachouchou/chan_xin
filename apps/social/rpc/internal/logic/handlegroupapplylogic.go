package logic

import (
	"context"
	"github.com/google/uuid"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/social"
	"github.com/ljp-lachouchou/chan_xin/apps/social/socialmodels"
	"github.com/ljp-lachouchou/chan_xin/deploy/constant"
	"github.com/ljp-lachouchou/chan_xin/pkg/ldefault"
	"github.com/ljp-lachouchou/chan_xin/pkg/lerr"
	"github.com/ljp-lachouchou/chan_xin/pkg/redislock"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	OtherAdminHandlerErr = errors.New("其他管理员正在处理")
)

type HandleGroupApplyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	redislock.DistributedLock
}

func NewHandleGroupApplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HandleGroupApplyLogic {
	l := &HandleGroupApplyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
	l.DistributedLock = redislock.NewRedisLock(l.svcCtx.Redis, ldefault.DEFAULT_REDIS_LOCK_KEY, uuid.NewString(), ldefault.DEFAULT_REDIS_LOCK_EXPIRE)
	return l
}

func (l *HandleGroupApplyLogic) HandleGroupApply(in *social.GroupApplyAction) (*social.GroupApplyActionResp, error) {
	findOne, err := l.svcCtx.GroupApplyModel.FindOne(l.ctx, in.ApplyId)
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "social-rpc HandleGroupApply", in.ApplyId)
	}
	switch constant.FriendApplyHandle(findOne.Status) {
	case constant.SuccessHandleApply:
		return nil, errors.WithStack(ApplyHasPassErr)
	case constant.FailHandleApply:
		return nil, errors.WithStack(ApplyHasRefuseErr)
	}
	//分布式锁
	acquire, err := l.DistributedLock.Acquire()
	defer l.DistributedLock.Release()
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NewCOMMONError(), err, "distributed lock acquire err")
	}
	if !acquire {
		return nil, errors.WithStack(OtherAdminHandlerErr)
	}
	if in.IsApproved {
		findOne.Status = 1
	} else {
		findOne.Status = 2
	}
	err = l.svcCtx.GroupApplyModel.Transx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		if err := l.svcCtx.GroupApplyModel.Update(l.ctx, session, findOne); err != nil {
			return lerr.NewWrapError(lerr.NEWDBError(), err, "social-rpc HandleGroupApply Tranx GroupApplyModel.Update", findOne)
		}
		if constant.FriendApplyHandle(findOne.Status) != constant.SuccessHandleApply {
			return nil
		}
		member := &socialmodels.GroupMember{
			GroupId:       findOne.TargetId,
			UserId:        findOne.ApplicantId,
			GroupNickname: "",
			ShowNickname:  1,
			IsAdmin:       0,
			IsMuted:       0,
			IsTopped:      0,
			Remark:        "",
		}
		_, err2 := l.svcCtx.GroupMemberModel.Insert(l.ctx, member)

		return err2

	})
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "social-rpc HandleGroupApply Transx")
	}
	return &social.GroupApplyActionResp{}, nil
}
func containsId(id string, list []*socialmodels.GroupMember) bool {
	for _, v := range list {
		if strings.Compare(id, v.UserId) == 0 {
			return true
		}
	}
	return false
}
