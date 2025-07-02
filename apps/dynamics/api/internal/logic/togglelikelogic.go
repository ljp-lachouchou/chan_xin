package logic

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/dynamics"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/dynamicsclient"

	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/api/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ToggleLikeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 点赞/取消点赞
func NewToggleLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ToggleLikeLogic {
	return &ToggleLikeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ToggleLikeLogic) ToggleLike(req *types.LikeAction) (*types.Empty, error) {
	postInfo, err2 := l.svcCtx.Dynamics.GetPostInfo(l.ctx, &dynamicsclient.GetPostInfoReq{
		PostId: req.PostId,
	})
	if err2 != nil {
		return nil, err2
	}
	_, err := l.svcCtx.Dynamics.ToggleLike(l.ctx, &dynamics.LikeAction{
		PostId:   req.PostId,
		LikerId:  req.LikerId,
		IsCancel: req.IsCancel,
	})
	if err != nil {
		return nil, err
	}
	_, err2 = l.svcCtx.Dynamics.CreateNotification(l.ctx, &dynamicsclient.CreateNotificationReq{
		UserId:        postInfo.UserId,
		Type:          0,
		TriggerUserId: req.LikerId,
		PostId:        req.PostId,
	})
	if err2 != nil {
		return nil, err2
	}
	return &types.Empty{}, nil
}
