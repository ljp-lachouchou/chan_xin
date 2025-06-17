package logic

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/pkg/lerr"

	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/social"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateGroupStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateGroupStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGroupStatusLogic {
	return &UpdateGroupStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 群内某一个成员的自己更改对于群的状态
func (l *UpdateGroupStatusLogic) UpdateGroupStatus(in *social.GroupStatusUpdate) (*social.GroupStatusUpdateResp, error) {
	member, err := l.svcCtx.GroupMemberModel.FindOneByGroupIdUserId(l.ctx, in.GroupId, in.UserId)
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "social-rpc UpdateGroupStatus GroupMemberModel.FindOneByGroupIdUserId", in.GroupId, in.UserId)
	}
	if in.Status == nil {
		return &social.GroupStatusUpdateResp{}, nil
	}
	if in.Status.IsMuted != nil {
		if *in.Status.IsMuted {
			member.IsMuted = 1
		} else {
			member.IsMuted = 0
		}
	}
	if in.Status.IsTopped != nil {
		if *in.Status.IsTopped {
			member.IsTopped = 1
		} else {
			member.IsTopped = 0
		}
	}
	if in.Status.Remark != nil {
		member.Remark = *in.Status.Remark
	}
	err = l.svcCtx.GroupMemberModel.Update(l.ctx, member)
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "social-rpc UpdateGroupStatus GroupMemberModel.Update", member)
	}
	return &social.GroupStatusUpdateResp{}, nil
}
