package logic

import (
	"context"

	"miniZhihu/app/user/rpc/internal/svc"
	"miniZhihu/app/user/rpc/types/service"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendSmsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendSmsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendSmsLogic {
	return &SendSmsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendSmsLogic) SendSms(in *service.SendSmsRequest) (*service.SendSmsResponse, error) {
	// todo: add your logic here and delete this line
	// TODO send sms
	return &service.SendSmsResponse{}, nil
}
