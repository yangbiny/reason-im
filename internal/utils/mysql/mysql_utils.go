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

type Transaction = *sql.Tx

func NewDatabaseTpl(db *sql.DB) *DatabaseTpl {
	return &DatabaseTpl{
		Db: db,
	}
}

func (databaseTpl *DatabaseTpl) prepareContext(ctx *context.Context, sqlStr string) (*sql.Stmt, error) {
	tx := (*ctx).Value("tx")
	if tx != nil {
		transaction := tx.(Transaction)
		return transaction.PrepareContext(*ctx, sqlStr)
	} else {
		connection, err := mysql.GetConnection(*ctx, databaseTpl.Db)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return connection.PrepareContext(*ctx, sqlStr)
	}
}

func (databaseTpl *DatabaseTpl) Insert(ctx *context.Context, sqlStr string, opts ...any) (int64, error) {
	prepareContext, err := databaseTpl.prepareContext(ctx, sqlStr)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	result, err := prepareContext.Exec(opts...)
	if err != nil {
		logger.Error(*ctx, "execute sqlStr has failed", "sqlStr", sqlStr, "opts", opts)
		return 0, errors.WithStack(err)
	}
	id, _ := result.LastInsertId()
	return id, nil
}

func (databaseTpl *DatabaseTpl) FindOne(ctx *context.Context, sql string, renderResult interface{}, opts ...any) (interface{}, error) {
	prepareContext, err := databaseTpl.prepareContext(ctx, sql)
	defer func() {
		if rec := recover(); rec != nil {
			var err error
			if as := errors.As(rec.(error), &err); as {
				logger.ErrorWithErr(*ctx, "execute has failed", errors.WithStack(err))
			}
		}
	}()
	queryContext, err := prepareContext.QueryContext(*ctx, opts...)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if !queryContext.Next() {
		return nil, nil
	}
	return mysql.RenderResult(queryContext, renderResult)
}

func (databaseTpl *DatabaseTpl) Update(ctx *context.Context, sql string, ops ...any) (int64, error) {
	prepareContext, err := databaseTpl.prepareContext(ctx, sql)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	result, err := prepareContext.Exec(ops...)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	return result.RowsAffected()
}

func (databaseTpl *DatabaseTpl) FindList(ctx *context.Context, sql string, renderResult interface{}, opts ...any) ([]interface{}, error) {
	prepareContext, err := databaseTpl.prepareContext(ctx, sql)
	queryContext, err := prepareContext.QueryContext(*ctx, opts...)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var result = make([]interface{}, 0)
	for queryContext.Next() {
		i, err := mysql.RenderResult(queryContext, renderResult)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		result = append(result, i)
	}
	return result, nil
}

func (databaseTpl *DatabaseTpl) WithTransaction(ctx *context.Context, f func(ctx *context.Context) error) error {
	tx, err := databaseTpl.Db.Begin()
	if err != nil {
		return errors.WithStack(err)
	}
	value := context.WithValue(*ctx, "tx", tx)
	err = f(&value)
	if err != nil {
		_ = tx.Rollback()
		return errors.WithStack(err)
	}
	err = tx.Commit()
	return err
}
