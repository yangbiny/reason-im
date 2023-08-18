package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/spf13/viper"
	"reason-im/internal/utils/logger"
)

type MysqlConfig struct {
	User     string ``
	Password string
	Url      string
}

var mysqlConfig *MysqlConfig

func GetConfig() *MysqlConfig {
	if mysqlConfig != nil {
		return mysqlConfig
	}
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err := viper.Unmarshal(&mysqlConfig)
	if err != nil {
		panic("读取配置文件失败")
	}
	mysqlConfig.Url = fmt.Sprintf("%s:%s@%s", mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Url)
	return mysqlConfig
}

func GetConnection(ctx context.Context) *sql.Conn {
	config := GetConfig()
	open, err := sql.Open("mysql", config.Url)
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
