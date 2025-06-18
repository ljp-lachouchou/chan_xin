package group

import (
	"net/http"

	"github.com/ljp-lachouchou/chan_xin/apps/social/api/internal/logic/group"
	"github.com/ljp-lachouchou/chan_xin/apps/social/api/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/social/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 邀请某人入群
func InviteToGroupHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GroupInvitation
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := group.NewInviteToGroupLogic(r.Context(), svcCtx)
		err := l.InviteToGroup(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
