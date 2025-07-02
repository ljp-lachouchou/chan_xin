package logic

import (
	"context"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/social"
	"github.com/ljp-lachouchou/chan_xin/apps/social/socialmodels"
	"github.com/ljp-lachouchou/chan_xin/pkg/lerr"
	"github.com/ljp-lachouchou/chan_xin/pkg/wuid"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	ErrGroupApplyHasSubmit = errors.New("你已经提交过相关申请")
)

type InviteToGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInviteToGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InviteToGroupLogic {
	return &InviteToGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InviteToGroupLogic) InviteToGroup(in *social.GroupInvitation) (*social.GroupInvitationResp, error) {
	err := l.svcCtx.GroupApplyModel.Transx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		var applyList []*socialmodels.GroupApply
		for _, v := range in.TargetIds {
			applyId := wuid.GenUid(l.svcCtx.Config.Mysql.DataSource)
			apply := &socialmodels.GroupApply{
				ApplyId:     applyId,
				ApplicantId: in.GroupId,
				TargetId:    v,
				GreetMsg:    "向你发出入群邀请",
				Status:      0,
			}
			applyList = append(applyList, apply)
		}
		_, err := l.svcCtx.GroupApplyModel.InsertByTargetIdList(l.ctx, session, applyList...)
		return err
	})
	if err != nil {
		errs := errors.Cause(err)
		if err1, ok := errs.(*mysql.MySQLError); ok {
			if err1.Number == 1062 {
				return nil, errors.WithStack(ErrGroupApplyHasSubmit)
			}
		}
		fmt.Printf("ss", errs)
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "social-rpc InviteToGroup tranx")
	}
	return &social.GroupInvitationResp{}, nil
}
