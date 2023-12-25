// @Author zhangjiaozhu 2023/10/3 11:02:00
package svc

import (
	"miniZhihu/app/like/mq/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
