package oauth

import (
	"gincmf/service/user/api/internal/types"
	"net/http"

	"gincmf/service/user/api/internal/logic/user/oauth"
	"gincmf/service/user/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RefreshHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req types.RefreshReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := oauth.NewRefreshLogic(r.Context(), svcCtx)
		resp, _ := l.Refresh(&req)
		httpx.OkJson(w, resp)
	}
}
