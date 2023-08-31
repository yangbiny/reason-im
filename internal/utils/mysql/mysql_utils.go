package mysql

import (
	"context"
	"database/sql"
	"reason-im/internal/config/mysql"
	"reason-im/internal/utils/logger"
)

type DatabaseTpl struct {
	Db *sql.DB
}

func (databaseTpl *DatabaseTpl) Insert(ctx context.Context, sql string, opts ...any) int64 {
	connection := mysql.GetConnection(ctx, databaseTpl.Db)
	defer mysql.CloseMysqlConn(connection, ctx)
	prepareContext, err := connection.PrepareContext(ctx, sql)
	if err != nil {
		panic(err)
	}
	result, err := prepareContext.Exec(opts)
	if err != nil {
		logger.Error(ctx, "execute sql has failed", "sql", sql, "opts", opts)
		panic("execute has failed")
	}
	id, _ := result.LastInsertId()
	return id
}

func (databaseTpl *DatabaseTpl) FindOne(ctx context.Context, sql string, renderResult interface{}, opts ...any) interface{} {
	connection := mysql.GetConnection(ctx, databaseTpl.Db)
	queryContext, err := connection.QueryContext(ctx, sql, opts...)
	if err != nil {
		panic(err)
	}
	if !queryContext.Next() {
		return nil
	}
	return mysql.RenderResult(queryContext, renderResult)
}
