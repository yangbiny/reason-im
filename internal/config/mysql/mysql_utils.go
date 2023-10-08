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

func Datasource() *sql.DB {
	mysqlConfig := config.GetImMysqlConfig()
	mysqlUrl := fmt.Sprintf("%s:%s@%s", mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Url)
	var err error
	db, err := sql.Open("mysql", mysqlUrl)
	if err != nil {
		logger.Error(nil, "db mysql has failed", err)
		panic(err)
	}
	db.SetConnMaxIdleTime(10)
	db.SetMaxOpenConns(100)
	return db
}

func GetConnection(ctx context.Context, db *sql.DB) (*sql.Conn, error) {
	conn, err := db.Conn(ctx)
	if err != nil {
		logger.Error(ctx, "get mysql conn has failed", err)
		return nil, err
	}
	return conn, nil
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
	t := field.Type()
	kind := t.Kind()
	switch kind {
	case reflect.Int64, reflect.Int8, reflect.Int32, reflect.Int16:
		field.SetInt(value.(int64))
	case reflect.String:
		field.SetString(string(value.([]byte)))
	case reflect.Float64, reflect.Float32:
		field.SetFloat(value.(float64))
	}
	if kind == reflect.Struct && t == reflect.TypeOf(time.Time{}) {
		var t time.Time
		if t1, ok := value.(time.Time); ok {
			t = t1
		} else {
			var err error
			t, err = time.Parse("2006-01-02 15:04:05", string(value.([]byte)))
			if err != nil {
				logger.Error(nil, "时间解析错误", "时间：", string(value.([]byte)))
				panic(err)
			}
		}
		field.Set(reflect.ValueOf(t))
	}
}
