package logic

import (
	"context"
	"github.com/go-sql-driver/mysql"
	"github.com/ljp-lachouchou/chan_xin/apps/social/socialmodels"
	"github.com/ljp-lachouchou/chan_xin/deploy/constant"
	"github.com/ljp-lachouchou/chan_xin/pkg/lerr"
	"github.com/ljp-lachouchou/chan_xin/pkg/wuid"
	"github.com/pkg/errors"
	"time"

	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/social"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	GroupIsOkErr          = errors.New("你已加入该群")
	GroupApplyHasExistErr = errors.New("请等待该群通过你的验证")
)

type ApplyGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewApplyGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApplyGroupLogic {
	return &ApplyGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ApplyGroupLogic) ApplyGroup(in *social.GroupApplyReq) (*social.GroupApplyResp, error) {
	hasGroup, err2 := l.svcCtx.GroupMemberModel.FindOneByGroupIdUserId(l.ctx, in.TargetId, in.ApplicantId)
	if err2 != nil && err2 != socialmodels.ErrNotFound {

		return nil, lerr.NewWrapError(lerr.NEWDBError(), err2, "social-rpc ApplyGroup", in.TargetId, in.ApplicantId)
	}
	if hasGroup != nil {
		return nil, errors.WithStack(GroupIsOkErr)
	}
	hasFriendApply, err := l.svcCtx.GroupApplyModel.FindByApplicantIdAndTargetId(l.ctx, in.ApplicantId, in.TargetId)
	if err != nil && err != socialmodels.ErrNotFound {
		return nil, lerr.NewWrapError(lerr.NewSYSTEMError(), err, "social-rpc ApplyGroup", in.ApplicantId, in.TargetId)
	}
	if hasFriendApply != nil {
		switch constant.FriendApplyHandle(hasFriendApply.Status) {
		case constant.FailHandleApply:
			hasFriendApply.Status = 0
			if err := l.svcCtx.GroupApplyModel.UpdateNoSession(l.ctx, hasFriendApply); err != nil {
				return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "social-rpc HandleFriendApply", hasFriendApply.ApplyId)
			}
			return &social.GroupApplyResp{
				ApplyId:   hasFriendApply.ApplyId,
				ApplyTime: time.Now().Unix(),
			}, nil
		default:
		}
		return nil, errors.WithStack(GroupApplyHasExistErr)
	}
	applyId := wuid.GenUid(l.svcCtx.Config.Mysql.DataSource)
	groupApply := socialmodels.GroupApply{
		ApplyId:     applyId,
		ApplicantId: in.ApplicantId,
		TargetId:    in.TargetId,
		GreetMsg:    in.GreetMsg,
		Status:      0,
	}
	_, err = l.svcCtx.GroupApplyModel.Insert(l.ctx, &groupApply)
	if err != nil {
		err2 := errors.Cause(err)
		if v, ok := err2.(*mysql.MySQLError); ok {
			if v.Number == 1062 {
				return nil, errors.WithStack(errors.New("已提交申请,请等待申请通知"))
			}
		}
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "social-rpc ApplyGroup Insert", groupApply)
	}
	return &social.GroupApplyResp{
		ApplyId:   applyId,
		ApplyTime: time.Now().Unix(),
	}, nil

}
