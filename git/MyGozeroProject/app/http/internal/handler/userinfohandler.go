package handler

import (
	"miniZhihu/app/http/internal/logic"
	"miniZhihu/app/http/internal/svc"
	"miniZhihu/common/response" // 1 替换为你真实的response包名
	"net/http"
)

func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo()
		response.Response(w, resp, err) //②  自定义模板内容

	}
}
