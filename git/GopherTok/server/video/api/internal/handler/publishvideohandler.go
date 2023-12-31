package handler

import (
	"GopherTok/common/response"
	"GopherTok/server/video/api/internal/logic"
	"GopherTok/server/video/api/internal/svc"
	"GopherTok/server/video/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func PublishVideoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PublishVideoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewPublishVideoLogic(r.Context(), svcCtx)
		resp, err := l.PublishVideo(&req, r)
		response.Response(r, w, resp, err) //②

	}
}
