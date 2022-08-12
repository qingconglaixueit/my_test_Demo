package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"my_test_demo/my_book_sys/book/api/internal/logic"
	"my_test_demo/my_book_sys/book/api/internal/svc"
	"my_test_demo/my_book_sys/book/api/internal/types"
)

func searchHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SearchReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewSearchLogic(r.Context(), svcCtx)
		resp, err := l.Search(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
