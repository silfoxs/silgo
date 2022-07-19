package handler

import (
	"net/http"

	"silgo/internal/logic"
	"silgo/internal/svc"
	"silgo/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func SilgoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewSilgoLogic(r.Context(), svcCtx)
		resp, err := l.Silgo(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
