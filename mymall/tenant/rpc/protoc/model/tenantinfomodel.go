package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TenantInfoModel = (*customTenantInfoModel)(nil)

type (
	// TenantInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTenantInfoModel.
	TenantInfoModel interface {
		tenantInfoModel
	}

	customTenantInfoModel struct {
		*defaultTenantInfoModel
	}
)

// NewTenantInfoModel returns a model for the database table.
func NewTenantInfoModel(conn sqlx.SqlConn) TenantInfoModel {
	return &customTenantInfoModel{
		defaultTenantInfoModel: newTenantInfoModel(conn),
	}
}
