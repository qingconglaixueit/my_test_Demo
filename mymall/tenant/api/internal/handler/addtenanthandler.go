package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"my_test_demo/mymall/tenant/api/internal/logic"
	"my_test_demo/mymall/tenant/api/internal/svc"
	"my_test_demo/mymall/tenant/api/internal/types"
)

func addTenantHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddTenantReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAddTenantLogic(r.Context(), svcCtx)
		resp, err := l.AddTenant(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
