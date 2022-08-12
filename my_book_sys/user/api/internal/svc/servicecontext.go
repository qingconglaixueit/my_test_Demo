package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"my_test_demo/my_book_sys/user/api/internal/config"
	"my_test_demo/my_book_sys/user/model"
)

type ServiceContext struct {
	Config        config.Config
	UserHttpModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UserHttpModel: model.NewUserModel(sqlx.NewMysql(c.DataSource)),
	}
}
