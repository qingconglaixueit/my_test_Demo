package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"my_test_demo/mymall/tenant/rpc/internal/config"
	"my_test_demo/mymall/tenant/rpc/protoc/model"
)

type ServiceContext struct {
	Config config.Config
	Model model.TenantInfoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model:model.NewTenantInfoModel(sqlx.NewMysql(c.DataSource)),
	}
}
