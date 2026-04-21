// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TagsDao is the data access object for the table tags.
type TagsDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  TagsColumns        // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// TagsColumns defines and stores column names for the table tags.
type TagsColumns struct {
	Id        string //
	Name      string //
	Slug      string //
	CreatedAt string //
}

// tagsColumns holds the columns for the table tags.
var tagsColumns = TagsColumns{
	Id:        "id",
	Name:      "name",
	Slug:      "slug",
	CreatedAt: "created_at",
}

// NewTagsDao creates and returns a new DAO object for table data access.
func NewTagsDao(handlers ...gdb.ModelHandler) *TagsDao {
	return &TagsDao{
		group:    "default",
		table:    "tags",
		columns:  tagsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TagsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TagsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TagsDao) Columns() TagsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TagsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TagsDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *TagsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
