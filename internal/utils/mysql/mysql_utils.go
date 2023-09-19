package mysql

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"reason-im/internal/config/mysql"
	"reason-im/internal/utils/logger"
)

type DatabaseTpl struct {
	Db *sql.DB
}

func NewDatabaseTpl(db *sql.DB) *DatabaseTpl {
	return &DatabaseTpl{
		Db: db,
	}
}

func (databaseTpl *DatabaseTpl) Insert(ctx context.Context, sql string, opts ...any) (int64, error) {
	connection, err := mysql.GetConnection(ctx, databaseTpl.Db)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	defer mysql.CloseMysqlConn(connection, ctx)
	prepareContext, err := connection.PrepareContext(ctx, sql)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	result, err := prepareContext.Exec(opts...)
	if err != nil {
		logger.Error(ctx, "execute sql has failed", "sql", sql, "opts", opts)
		return 0, errors.WithStack(err)
	}
	id, _ := result.LastInsertId()
	return id, nil
}

func (databaseTpl *DatabaseTpl) FindOne(ctx context.Context, sql string, renderResult interface{}, opts ...any) (interface{}, error) {
	connection, err := mysql.GetConnection(ctx, databaseTpl.Db)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	queryContext, err := connection.QueryContext(ctx, sql, opts...)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if !queryContext.Next() {
		return nil, nil
	}
	return mysql.RenderResult(queryContext, renderResult), nil
}

func (databaseTpl *DatabaseTpl) Update(ctx context.Context, sql string, ops ...any) (int64, error) {
	connection, err := mysql.GetConnection(ctx, databaseTpl.Db)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	defer mysql.CloseMysqlConn(connection, ctx)
	prepareContext, err := connection.PrepareContext(ctx, sql)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	result, err := prepareContext.Exec(ops...)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	return result.RowsAffected()
}

func (databaseTpl *DatabaseTpl) FindList(ctx context.Context, sql string, renderResult interface{}, opts ...any) ([]interface{}, error) {
	connection, err := mysql.GetConnection(ctx, databaseTpl.Db)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	queryContext, err := connection.QueryContext(ctx, sql, opts...)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var result = make([]interface{}, 0)
	for queryContext.Next() {
		i := mysql.RenderResult(queryContext, renderResult)
		result = append(result, i)
	}
	return result, nil
}
