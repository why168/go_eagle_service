package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go_eagle_service/internal/logic"
	"go_eagle_service/internal/svc"
	"go_eagle_service/internal/types"
)

func eagleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewEagLogic(r.Context(), svcCtx)
		resp, err := l.GoEagleService(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
