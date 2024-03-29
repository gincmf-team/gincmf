package authAccess

import (
	"net/http"

	"gincmf/service/user/api/internal/logic/user/admin/authAccess"
	"gincmf/service/user/api/internal/svc"
	"gincmf/service/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func StoreHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AccessStore
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := authAccess.NewStoreLogic(r.Context(), svcCtx)
		resp, _ := l.Store(&req)
		httpx.OkJson(w, resp)
	}
}
