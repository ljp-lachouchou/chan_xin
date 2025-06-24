package chat

import (
	"net/http"

	"github.com/ljp-lachouchou/chan_xin/apps/im/api/internal/logic/chat"
	"github.com/ljp-lachouchou/chan_xin/apps/im/api/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/im/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SetUpUserConversationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SetUpUserConversationReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := chat.NewSetUpUserConversationLogic(r.Context(), svcCtx)
		err := l.SetUpUserConversation(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
