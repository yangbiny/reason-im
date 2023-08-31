package mysql

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"reason-im/internal/config"
	"reason-im/internal/utils/logger"
	"reflect"
)

func GetConnection(ctx context.Context) *sql.Conn {
	mysqlConfig := config.GetImMysqlConfig()
	mysqlUrl := fmt.Sprintf("%s:%s@%s", mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Url)
	open, err := sql.Open("mysql", mysqlUrl)
	if err != nil {
		logger.Error(ctx, "open mysql has failed", err)
		panic("open mysql has failed ")
	}
	conn, err := open.Conn(ctx)
	if err != nil {
		logger.Error(ctx, "get mysql conn has failed", err)
		panic("get mysql conn has failed")
	}
	return conn
}

func CloseMysqlConn(conn *sql.Conn, context context.Context) {
	if err := conn.Close(); err != nil {
		logger.Error(context, "close mysql connection has failed")
	}
}

func RenderResult(rows *sql.Rows, resultType interface{}) interface{} {
	of := reflect.TypeOf(resultType)
	if of.Kind() != reflect.Struct {
		panic("请传入结构体")
	}
	columns, _ := rows.Columns()
	columnsSize := len(columns)
	columnValue := make([]interface{}, columnsSize)
	valuePointers := make([]interface{}, columnsSize)
	for i := range valuePointers {
		valuePointers[i] = &columnValue[i]
	}
	err := rows.Scan(valuePointers...)
	if err != nil {
		panic(err)
	}
	columnKeyIndexMap := make(map[string]int)
	for i := range columns {
		columnKeyIndexMap[columns[i]] = i
	}
	value := reflect.New(of)
	for i := 0; i < value.NumField(); i++ {
		field := of.Field(0)
		columnName := field.Tag.Get("mysql")
		index := columnKeyIndexMap[columnName]
		columnValue := valuePointers[index]
		value.Field(i).SetInt(columnValue.(int64))
	}
	return value
}
