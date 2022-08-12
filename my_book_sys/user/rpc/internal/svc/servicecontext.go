package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"my_test_demo/my_book_sys/user/model"
	"my_test_demo/my_book_sys/user/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	UserRpcModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		UserRpcModel:model.NewUserModel(sqlx.NewMysql(c.DataSource)),
	}
}
