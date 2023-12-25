package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"miniZhihu/app/article/api/internal/logic"
	"miniZhihu/app/article/api/internal/svc"
	"miniZhihu/app/article/api/internal/types"
	"miniZhihu/common/response" // 1 替换为你真实的response包名
	"net/http"
)

func PublishHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PublishRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewPublishLogic(r.Context(), svcCtx)
		resp, err := l.Publish(&req)
		response.Response(w, resp, err) //②  自定义模板内容

	}
}
