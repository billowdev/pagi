package configs

import (
	"strconv"

	"github.com/spf13/viper"
)

var (
	APP_DEBUG_MODE bool
	APP_PORT string

	DB_NAME        string
	// DB_USERNAME    string
	// DB_PASSWORD    string
	// DB_HOST        string
	// DB_PORT        string
	// DB_SSL_MODE    string
	// DB_SCHEMA      string
)

func init() {
	InitConfig()

	APP_PORT = viper.GetString("APP_PORT")


	var err error
	APP_DEBUG_MODE, err = strconv.ParseBool(viper.GetString("APP_DEBUG_MODE"))
	if err != nil {
		APP_DEBUG_MODE = false
	}

	DB_NAME = viper.GetString("DB_NAME")
	// DB_USERNAME = viper.GetString("DB_USERNAME")
	// DB_PASSWORD = viper.GetString("DB_PASSWORD")
	// DB_HOST = viper.GetString("DB_HOST")
	// DB_PORT = viper.GetString("DB_PORT")
	// DB_SSL_MODE = viper.GetString("DB_SSL_MODE")
	// if DB_SSL_MODE == "" {
	// 	DB_SSL_MODE = "default"
	// }
	// DB_SCHEMA = viper.GetString("DB_SCHEMA")
}
