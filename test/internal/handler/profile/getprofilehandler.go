package profile

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-pj/test/internal/logic/profile"
	"go-zero-pj/test/internal/svc"
	"go-zero-pj/test/internal/types"
)

func GetProfileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := profile.NewGetProfileLogic(r.Context(), svcCtx)
		resp, err := l.GetProfile(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
