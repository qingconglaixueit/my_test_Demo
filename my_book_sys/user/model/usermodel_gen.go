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
	userFieldNames          = builder.RawFieldNames(&User{})
	userRows                = strings.Join(userFieldNames, ",")
	userRowsExpectAutoSet   = strings.Join(stringx.Remove(userFieldNames, "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	userRowsWithPlaceHolder = strings.Join(stringx.Remove(userFieldNames, "`stu_id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"
)

type (
	userModel interface {
		Insert(ctx context.Context, data *User) (sql.Result, error)
		FindOne(ctx context.Context, stuId string) (*User, error)
		FindOneByName(ctx context.Context, name string) (*User, error)
		Update(ctx context.Context, data *User) error
		Delete(ctx context.Context, stuId string) error
	}

	defaultUserModel struct {
		conn  sqlx.SqlConn
		table string
	}

	User struct {
		StuId    string `db:"stu_id"`   // stu_id
		Name     string `db:"name"`     // name
		Password string `db:"password"` // password
		Gender   string `db:"gender"`   // gender
	}
)

func newUserModel(conn sqlx.SqlConn) *defaultUserModel {
	return &defaultUserModel{
		conn:  conn,
		table: "`user`",
	}
}

func (m *defaultUserModel) Delete(ctx context.Context, stuId string) error {
	query := fmt.Sprintf("delete from %s where `stu_id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, stuId)
	return err
}

func (m *defaultUserModel) FindOne(ctx context.Context, stuId string) (*User, error) {
	query := fmt.Sprintf("select %s from %s where `stu_id` = ? limit 1", userRows, m.table)
	var resp User
	err := m.conn.QueryRowCtx(ctx, &resp, query, stuId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) FindOneByName(ctx context.Context, name string) (*User, error) {
	query := fmt.Sprintf("select %s from %s where `name` = ? limit 1", userRows, m.table)
	var resp User
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

func (m *defaultUserModel) Insert(ctx context.Context, data *User) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, userRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.StuId, data.Name, data.Password, data.Gender)
	return ret, err
}

func (m *defaultUserModel) Update(ctx context.Context, data *User) error {
	query := fmt.Sprintf("update %s set %s where `stu_id` = ?", m.table, userRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Name, data.Password, data.Gender, data.StuId)
	return err
}

func (m *defaultUserModel) tableName() string {
	return m.table
}
