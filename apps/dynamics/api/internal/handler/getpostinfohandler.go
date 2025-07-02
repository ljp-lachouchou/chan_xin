package handler

import (
	"net/http"

	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/api/internal/logic"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/api/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 单个post信息
func getPostInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetPostInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetPostInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetPostInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
