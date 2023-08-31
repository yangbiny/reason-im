package mysql

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"reason-im/internal/config"
	"reason-im/internal/utils/logger"
	"reflect"
	"time"
)

var db *sql.DB

func init() {
	mysqlConfig := config.GetImMysqlConfig()
	mysqlUrl := fmt.Sprintf("%s:%s@%s", mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Url)
	var err error
	db, err = sql.Open("mysql", mysqlUrl)
	if err != nil {
		logger.Error(nil, "db mysql has failed", err)
		panic("db mysql has failed ")
	}
	db.SetConnMaxIdleTime(10)
	db.SetMaxOpenConns(100)
}

func GetConnection(ctx context.Context) *sql.Conn {
	conn, err := db.Conn(ctx)
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
	columnTypes, _ := rows.ColumnTypes()
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
	value := reflect.New(of).Elem()
	for i := 0; i < of.NumField(); i++ {
		field := of.Field(i)
		columnName := field.Tag.Get("mysql")
		index := columnKeyIndexMap[columnName]
		columnValue := columnValue[index]
		setValue(value.Field(i), columnValue, columnTypes[i])
	}
	return value.Interface()
}

func setValue(field reflect.Value, value interface{}, valueType *sql.ColumnType) {
	switch valueType.DatabaseTypeName() {
	case "INT", "INTEGER":
		field.SetInt(value.(int64))
	case "VARCHAR", "TEXT":
		field.SetString(string(value.([]byte)))
	case "FLOAT":
		field.SetFloat(value.(float64))
	case "DATETIME", "TIMESTAMP":
		t, err := time.Parse("2006-01-02 15:04:05", string(value.([]byte)))
		if err != nil {
			logger.Error(nil, "时间解析错误", "时间：", string(value.([]byte)))
			panic("时间解析错误")
		}
		field.Set(reflect.ValueOf(t))
	}
}
