package user

import (
	"net/http"

	"github.com/ljp-lachouchou/chan_xin/apps/user/api/internal/logic/user"
	"github.com/ljp-lachouchou/chan_xin/apps/user/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 单个用户查询
func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
