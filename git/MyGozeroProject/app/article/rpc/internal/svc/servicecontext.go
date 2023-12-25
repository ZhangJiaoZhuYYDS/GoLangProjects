package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"miniZhihu/app/article/model/article"
	"miniZhihu/app/article/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	ArticleModel article.ArticleModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		ArticleModel: article.NewArticleModel(sqlx.NewMysql(c.Database,)),
	}
}
