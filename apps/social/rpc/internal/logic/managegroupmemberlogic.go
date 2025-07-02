package logic

import (
	"context"
	"database/sql"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/social"
	"github.com/ljp-lachouchou/chan_xin/apps/social/socialmodels"
	"github.com/ljp-lachouchou/chan_xin/pkg/lerr"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/zeromicro/go-zero/core/logx"
)

type ManageGroupMemberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewManageGroupMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ManageGroupMemberLogic {
	return &ManageGroupMemberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ManageGroupMemberLogic) ManageGroupMember(in *social.GroupMemberManage) (*social.GroupMemberManageResp, error) {
	member, err := l.svcCtx.GroupMemberModel.FindOneByGroupIdUserId(l.ctx, in.GroupId, in.TargetId)
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "social-rpc ManageGroupMember GroupMemberModel.FindOneByGroupIdUserId", in.GroupId, in.TargetId)
	}
	err = l.svcCtx.GroupOperationModel.Transx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		member.IsAdmin = 1
		err := l.svcCtx.GroupMemberModel.UpdateWithSession(l.ctx, session, member)
		if err != nil {
			return err
		}
		data := &socialmodels.GroupOperation{
			GroupId:    in.GroupId,
			OperatorId: in.OperatorId,
			TargetId: sql.NullString{
				String: in.TargetId,
				Valid:  true,
			},
			ActionType: in.Action.String(),
		}
		_, err = l.svcCtx.GroupOperationModel.InertWithSession(l.ctx, session, data)
		return err
	})
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "social-rpc ManageGroupMember transx")
	}
	return &social.GroupMemberManageResp{}, nil
}
