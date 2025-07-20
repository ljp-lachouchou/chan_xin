package friend

import (
	"net/http"

	"github.com/ljp-lachouchou/chan_xin/apps/social/api/internal/logic/friend"
	"github.com/ljp-lachouchou/chan_xin/apps/social/api/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/social/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取谁想添加我为好友的列表
func GetHandleFriendApplyListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HandleFriendApplyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := friend.NewGetHandleFriendApplyListLogic(r.Context(), svcCtx)
		resp, err := l.GetHandleFriendApplyList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
