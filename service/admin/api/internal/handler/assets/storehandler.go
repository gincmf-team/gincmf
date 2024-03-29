package assets

import (
	"net/http"

	"gincmf/service/admin/api/internal/logic/assets"
	"gincmf/service/admin/api/internal/svc"
	"gincmf/service/admin/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func StoreHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AssetsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := assets.NewStoreLogic(r.Context(), svcCtx)
		resp, err := l.Store(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
