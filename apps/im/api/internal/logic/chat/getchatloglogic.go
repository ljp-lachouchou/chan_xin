package chat

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/apps/im/rpc/im"

	"github.com/ljp-lachouchou/chan_xin/apps/im/api/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/im/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetChatLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetChatLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetChatLogLogic {
	return &GetChatLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetChatLogLogic) GetChatLog(req *types.GetChatLogReq) (*types.GetChatLogResp, error) {
	logResp, err := l.svcCtx.Im.GetChatLog(l.ctx, &im.GetChatLogReq{
		ConversationId: req.ConversationId,
		StartSendTime:  req.StartSendTime,
		EndSendTime:    req.EndSendTime,
		Count:          req.Count,
		MsgId:          req.MsgId,
	})
	if err != nil {
		return nil, err
	}
	var list []types.ChatLog
	for _, v := range logResp.List {
		list = append(list, types.ChatLog{
			Id:             v.Id,
			ConversationId: v.ConversationId,
			SendId:         v.SendId,
			RecvId:         v.RecvId,
			MsgType:        v.MsgType,
			MsgContent:     v.MsgContent,
			ChatType:       v.ChatType,
			SendTime:       v.SendTime,
		})
	}
	return &types.GetChatLogResp{
		List: list,
	}, nil

}
