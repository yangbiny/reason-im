package config

import (
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type Config struct {
	Mysql MysqlConfig `yaml:"mysql"`
}

type MysqlConfig struct {
	ImDatabaseConfig DatabasesConfig `yaml:"im"`
}

type DatabasesConfig struct {
	User     string `yaml:"username"`
	Password string `yaml:"password"`
	Url      string `yaml:"url"`
}

var config = &Config{}

func withTagName(tag string) viper.DecoderConfigOption {
	return func(config *mapstructure.DecoderConfig) {
		config.TagName = tag
	}
}

func init() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("config")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	opts := []viper.DecoderConfigOption{
		withTagName("yaml"),
	}
	err := viper.Unmarshal(config, opts...)
	if err != nil {
		panic("读取配置文件失败")
	}
}

func GetImMysqlConfig() *DatabasesConfig {
	return &config.Mysql.ImDatabaseConfig
}
