package logic

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/pkg/lerr"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/social"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveAdminLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveAdminLogic {
	return &RemoveAdminLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveAdminLogic) RemoveAdmin(in *social.RemoveAdminReq) (*social.RemoveAdminResp, error) {
	member, err := l.svcCtx.GroupMemberModel.FindOneByGroupIdUserId(l.ctx, in.GroupId, in.TargetId)
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "social-rpc RemoveAdmin GroupMemberModel.FindOneByGroupIdUserId", in.GroupId, in.TargetId)
	}
	err = l.svcCtx.GroupOperationModel.Transx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		member.IsAdmin = 0
		err := l.svcCtx.GroupMemberModel.UpdateWithSession(l.ctx, session, member)
		if err != nil {
			return err
		}
		return l.svcCtx.GroupOperationModel.DeleteByGIdOidTid(l.ctx, session, in.GroupId, in.OperatorId, in.TargetId)

	})
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "social-rpc RemoveAdmin transx")
	}
	return &social.RemoveAdminResp{}, nil
}
