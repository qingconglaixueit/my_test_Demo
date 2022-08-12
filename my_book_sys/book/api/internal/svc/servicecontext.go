package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"my_test_demo/my_book_sys/book/api/internal/config"
	"my_test_demo/my_book_sys/book/model"
)

type ServiceContext struct {
	Config config.Config
	BookHttpModel model.BookModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		BookHttpModel:model.NewBookModel(sqlx.NewMysql(c.DataSource)),
	}
}
