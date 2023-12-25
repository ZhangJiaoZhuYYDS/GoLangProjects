package logic

import (
	"context"

	"miniZhihu/app/user/rpc/internal/svc"
	"miniZhihu/app/user/rpc/types/service"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindByIdLogic {
	return &FindByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindByIdLogic) FindById(in *service.FindByIdRequest) (*service.FindByIdResponse, error) {
	// todo: add your logic here and delete this line
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		logx.Errorf("查询 用户id：%d , error %v ",in.UserId,err)
		return nil, err
	}
	return &service.FindByIdResponse{
		UserId: user.Id,
		Username: user.Username,
		Avatar: user.Avatar,
	}, nil
}
