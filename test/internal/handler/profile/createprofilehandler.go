package profile

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-pj/test/internal/logic/profile"
	"go-zero-pj/test/internal/svc"
	"go-zero-pj/test/internal/types"
)

func CreateProfileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := profile.NewCreateProfileLogic(r.Context(), svcCtx)
		err := l.CreateProfile(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
