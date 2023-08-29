package config

import "github.com/spf13/viper"

type Config struct {
	mysql *MysqlConfig `yaml:"mysql"`
}

type MysqlConfig struct {
	ImDatabaseConfig *DatabasesConfig `yaml:"im"`
}

type DatabasesConfig struct {
	User     string `yaml:"username"`
	Password string `yaml:"password"`
	Url      string `yaml:"url"`
}

var config *Config

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.SetConfigFile("/Volumes/workspace/admin/reason-im/config/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic("读取配置文件失败")
	}
	err := viper.Unmarshal(&config)
	if err != nil {
		panic("读取配置文件失败")
	}
}

func GetImMysqlConfig() *DatabasesConfig {
	return config.mysql.ImDatabaseConfig
}
