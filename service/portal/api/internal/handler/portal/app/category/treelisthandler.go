package category

import (
	"net/http"

	"gincmf/service/portal/api/internal/logic/portal/app/category"
	"gincmf/service/portal/api/internal/svc"
	"gincmf/service/portal/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func TreeListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OneReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := category.NewTreeListLogic(r.Context(), svcCtx)
		resp := l.TreeList(&req)
		httpx.OkJson(w, resp)
	}
}
