package mysql

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"reason-im/internal/config"
	"reason-im/internal/utils/logger"
)

func GetConnection(ctx context.Context) *sql.Conn {
	mysqlConfig := config.GetImMysqlConfig()
	open, err := sql.Open("mysql", mysqlConfig.Url)
	if err != nil {
		logger.Warn(ctx, "open mysql has failed", err)
		panic("open mysql has failed ")
	}
	conn, err := open.Conn(ctx)
	if err != nil {
		logger.Warn(ctx, "get mysql conn has failed", err)
		panic("get mysql conn has failed")
	}
	return conn
}

func CloseMysqlConn(conn *sql.Conn, context context.Context) {
	if err := conn.Close(); err != nil {
		logger.Warn(context, "close mysql connection has failed")
	}
}
