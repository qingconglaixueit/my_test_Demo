package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"my_test_demo/mymall/order/api/internal/config"
	"my_test_demo/mymall/tenant/rpc/tenant"
)

type ServiceContext struct {
	Config config.Config
	TenantRpc tenant.Tenant
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		TenantRpc:tenant.NewTenant(zrpc.MustNewClient(c.TenantRpc)),
	}
}
