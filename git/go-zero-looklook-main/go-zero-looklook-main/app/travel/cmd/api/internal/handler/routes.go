// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	homestay "looklook/app/travel/cmd/api/internal/handler/homestay"
	homestayBussiness "looklook/app/travel/cmd/api/internal/handler/homestayBussiness"
	homestayComment "looklook/app/travel/cmd/api/internal/handler/homestayComment"
	"looklook/app/travel/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/homestay/homestayList",
				Handler: homestay.HomestayListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/homestay/businessList",
				Handler: homestay.BusinessListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/homestay/guessList",
				Handler: homestay.GuessListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/homestay/homestayDetail",
				Handler: homestay.HomestayDetailHandler(serverCtx),
			},
		},
		rest.WithPrefix("/travel/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/homestayBussiness/goodBoss",
				Handler: homestayBussiness.GoodBossHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/homestayBussiness/homestayBussinessList",
				Handler: homestayBussiness.HomestayBussinessListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/homestayBussiness/homestayBussinessDetail",
				Handler: homestayBussiness.HomestayBussinessDetailHandler(serverCtx),
			},
		},
		rest.WithPrefix("/travel/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/homestayComment/commentList",
				Handler: homestayComment.CommentListHandler(serverCtx),
			},
		},
		rest.WithPrefix("/travel/v1"),
	)
}