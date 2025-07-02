package group

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/social/rpc/socialservice"

	"github.com/ljp-lachouchou/chan_xin/apps/social/api/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/social/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取群信息
func NewGetGroupInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupInfoLogic {
	return &GetGroupInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGroupInfoLogic) GetGroupInfo(req *types.GroupInfoRequest) (resp *types.GroupInfo, err error) {
	in := &socialservice.GroupInfoRequest{
		UserId:  req.UserId,
		GroupId: req.GroupId,
	}
	info, err := l.svcCtx.SocialService.GetGroupInfo(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return &types.GroupInfo{
		GroupId:    info.GroupId,
		Name:       info.Name,
		OwnerId:    info.OwnerId,
		AdminIds:   info.AdminIds,
		MemberIds:  info.MemberIds,
		MaxMembers: info.MaxMembers,
	}, nil
}
