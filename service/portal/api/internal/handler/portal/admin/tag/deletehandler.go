package tag

import (
	"net/http"

	"gincmf/service/portal/api/internal/logic/portal/admin/tag"
	"gincmf/service/portal/api/internal/svc"
	"gincmf/service/portal/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OneReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := tag.NewDeleteLogic(r.Context(), svcCtx)
		resp := l.Delete(&req)
		httpx.OkJson(w, resp)
	}
}
