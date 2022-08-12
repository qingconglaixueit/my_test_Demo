package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ BookModel = (*customBookModel)(nil)

type (
	// BookModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBookModel.
	BookModel interface {
		bookModel
	}

	customBookModel struct {
		*defaultBookModel
	}
)

// NewBookModel returns a model for the database table.
func NewBookModel(conn sqlx.SqlConn) BookModel {
	return &customBookModel{
		defaultBookModel: newBookModel(conn),
	}
}
