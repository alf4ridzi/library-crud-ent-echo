package config

import (
	"github.com/spf13/viper"
)

type env struct {
	MysqlConfigEnv `mapstructure:",squash"` // Gunakan squash untuk struct bersarang
}

var (
	AppConfig *env
)

type MysqlConfigEnv struct {
	DBHost           string `mapstructure:"MYSQL_DB_HOST"`
	DBName           string `mapstructure:"MYSQL_DB_NAME"`
	DBUserName       string `mapstructure:"MYSQL_DB_USERNAME"`
	DBPassword       string `mapstructure:"MYSQL_DB_PASSWORD"`
	DBPort           int    `mapstructure:"MYSQL_DB_PORT"`
	JwtAuthSecret    string `mapstructure:"JWT_AUTH_SECRET"`
	JwtRefreshSecret string `mapstructure:"JWT_REFRESH_SECRET"`
}

func LoadEnv() error {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	tempEnv := &env{}
	err := viper.Unmarshal(&tempEnv.MysqlConfigEnv)
	if err != nil {
		return err
	}

	AppConfig = tempEnv
	return nil
}
