package article

import (
	"net/http"

	"gincmf/service/portal/api/internal/logic/portal/admin/article"
	"gincmf/service/portal/api/internal/svc"
	"gincmf/service/portal/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func StoreHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ArticleSaveReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := article.NewStoreLogic(r.Context(), svcCtx)
		resp := l.Store(&req)
		httpx.OkJson(w, resp)
	}
}
