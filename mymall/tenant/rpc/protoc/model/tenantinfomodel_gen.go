// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	tenantInfoFieldNames          = builder.RawFieldNames(&TenantInfo{})
	tenantInfoRows                = strings.Join(tenantInfoFieldNames, ",")
	tenantInfoRowsExpectAutoSet   = strings.Join(stringx.Remove(tenantInfoFieldNames, "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	tenantInfoRowsWithPlaceHolder = strings.Join(stringx.Remove(tenantInfoFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"
)

type (
	tenantInfoModel interface {
		Insert(ctx context.Context, data *TenantInfo) (sql.Result, error)
		FindOne(ctx context.Context, id string) (*TenantInfo, error)
		Update(ctx context.Context, data *TenantInfo) error
		Delete(ctx context.Context, id string) error
		FindOneByName(ctx context.Context, name string) (*TenantInfo, error)
	}

	defaultTenantInfoModel struct {
		conn  sqlx.SqlConn
		table string
	}

	TenantInfo struct {
		Id   string `db:"id"`   // tenant id
		Name string `db:"name"` // tenant name
		Addr string `db:"addr"` // tenant addr
	}
)

func newTenantInfoModel(conn sqlx.SqlConn) *defaultTenantInfoModel {
	return &defaultTenantInfoModel{
		conn:  conn,
		table: "`tenant_info`",
	}
}

func (m *defaultTenantInfoModel) Delete(ctx context.Context, id string) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultTenantInfoModel) FindOne(ctx context.Context, id string) (*TenantInfo, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", tenantInfoRows, m.table)
	var resp TenantInfo
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
func (m *defaultTenantInfoModel) FindOneByName(ctx context.Context, name string) (*TenantInfo, error) {
	query := fmt.Sprintf("select %s from %s where `name` = ? limit 1", tenantInfoRows, m.table)
	var resp TenantInfo
	err := m.conn.QueryRowCtx(ctx, &resp, query, name)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTenantInfoModel) Insert(ctx context.Context, data *TenantInfo) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, tenantInfoRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Id, data.Name, data.Addr)
	return ret, err
}

func (m *defaultTenantInfoModel) Update(ctx context.Context, data *TenantInfo) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, tenantInfoRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Name, data.Addr, data.Id)
	return err
}

func (m *defaultTenantInfoModel) tableName() string {
	return m.table
}
