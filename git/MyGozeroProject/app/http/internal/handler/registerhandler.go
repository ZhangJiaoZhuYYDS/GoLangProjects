package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"miniZhihu/app/http/internal/logic"
	"miniZhihu/app/http/internal/svc"
	"miniZhihu/app/http/internal/types"
	"miniZhihu/common/response" // 1 替换为你真实的response包名
	"net/http"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(),w,err)
		}else {
			httpx.OkJsonCtx(r.Context(),w,resp)
		}
		response.Response(w, resp, err) //②  自定义模板内容
	}
}