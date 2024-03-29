package account

import (
	"net/http"

	"gincmf/service/user/api/internal/logic/user/admin/account"
	"gincmf/service/user/api/internal/svc"
	"gincmf/service/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func EditHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AdminSaveReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := account.NewEditLogic(r.Context(), svcCtx)
		resp, _ := l.Edit(&req)
		httpx.OkJson(w, resp)
	}
}
