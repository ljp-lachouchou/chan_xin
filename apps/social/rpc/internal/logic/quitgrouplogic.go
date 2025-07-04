package logic

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/pkg/lerr"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"

	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/social"

	"github.com/zeromicro/go-zero/core/logx"
)

type QuitGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQuitGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuitGroupLogic {
	return &QuitGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QuitGroupLogic) QuitGroup(in *social.GroupQuitRequest) (*social.GroupQuitResp, error) {
	findOne, err := l.svcCtx.GroupInfoModel.FindOne(l.ctx, in.GroupId)
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "social-rpc QuitGroupLogic.QuitGroup FindOne", in.GroupId)
	}
	if strings.Compare(findOne.OwnerId, in.UserId) == 0 {
		if err := l.svcCtx.GroupInfoModel.Delete(l.ctx, in.GroupId); err != nil {
			return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "social-rpc QuitGroupLogic.QuitGroup Delete", in.GroupId)
		}
		return &social.GroupQuitResp{}, nil
	}
	err = l.svcCtx.GroupMemberModel.Transx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		if err := l.svcCtx.GroupApplyModel.DeleteByGIdAndUId(l.ctx, session, in.GroupId, in.UserId); err != nil {
			return lerr.NewWrapError(lerr.NEWDBError(), err, "social-rpc HandleFriendApply", in.GroupId, in.UserId)
		}
		return l.svcCtx.GroupMemberModel.DeleteByGIdAndUId(l.ctx, in.GroupId, in.UserId)
	})
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "social-rpc QuitGroup GroupMemberModel.Transx")
	}
	return &social.GroupQuitResp{}, nil
}
