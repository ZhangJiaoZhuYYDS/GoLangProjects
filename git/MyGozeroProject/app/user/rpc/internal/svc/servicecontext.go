package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"miniZhihu/app/user/model"
	"miniZhihu/app/user/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	return &ServiceContext{
		Config: c,
		UserModel: model.NewUserModel(conn,c.CacheRedis),
	}
}
