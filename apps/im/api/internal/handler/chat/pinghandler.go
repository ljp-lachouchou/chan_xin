package chat

import (
	"net/http"

	"github.com/ljp-lachouchou/chan_xin/apps/im/api/internal/logic/chat"
	"github.com/ljp-lachouchou/chan_xin/apps/im/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func PingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := chat.NewPingLogic(r.Context(), svcCtx)
		err := l.Ping()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
