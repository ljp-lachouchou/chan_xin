package chat

import (
	"net/http"

	"github.com/ljp-lachouchou/chan_xin/apps/im/api/internal/logic/chat"
	"github.com/ljp-lachouchou/chan_xin/apps/im/api/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/im/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetChatLogHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetChatLogReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := chat.NewGetChatLogLogic(r.Context(), svcCtx)
		resp, err := l.GetChatLog(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
