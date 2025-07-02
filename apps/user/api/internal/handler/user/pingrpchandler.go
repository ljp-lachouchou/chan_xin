package user

import (
	"net/http"

	"github.com/ljp-lachouchou/chan_xin/apps/user/api/internal/logic/user"
	"github.com/ljp-lachouchou/chan_xin/apps/user/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 保持与etcd的连接
func PingRpcHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewPingRpcLogic(r.Context(), svcCtx)
		err := l.PingRpc()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
