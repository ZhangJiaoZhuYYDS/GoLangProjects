package handler

import (
	"miniZhihu/app/article/api/internal/logic"
	"miniZhihu/app/article/api/internal/svc"
	"miniZhihu/common/response" // 1 替换为你真实的response包名
	"net/http"
)

func UploadCoverHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewUploadCoverLogic(r.Context(), svcCtx)
		resp, err := l.UploadCover(r)
		response.Response(w, resp, err) //②  自定义模板内容

	}
}
