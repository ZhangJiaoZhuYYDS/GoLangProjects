package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"miniZhihu/app/http/internal/config"
	"miniZhihu/app/user/rpc/user"
)

type ServiceContext struct {
	Config config.Config
	UserRPC user.User
	BizRedis *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	userRpc := zrpc.MustNewClient(c.UserRPC)

	return &ServiceContext{
		Config: c,
		UserRPC: user.NewUser(userRpc),
		BizRedis: redis.New(c.BizRedis.Host,redis.WithPass(c.BizRedis.Pass)),
	}
}
