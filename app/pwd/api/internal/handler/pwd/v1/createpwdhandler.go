package v1

import (
	"net/http"

	v1 "go-zero-pwd/app/pwd/api/internal/logic/pwd/v1"
	"go-zero-pwd/app/pwd/api/internal/svc"
	"go-zero-pwd/app/pwd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreatePwdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreatePwdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := v1.NewCreatePwdLogic(r.Context(), svcCtx)
		resp, err := l.CreatePwd(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
