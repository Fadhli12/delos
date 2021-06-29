package config

import (
	"github.com/apex/log"
	"github.com/spf13/viper"
)

var Config Env

type Env struct {
	PostgreSQL
}

type PostgreSQL struct {
	USER          string
	HOST          string
	PORT          string
	PASSWORD      string
	DATABASE_NAME string
}

func Environment() {
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	viper.SetConfigName("config.yml")

	env := Env{}
	if err := viper.ReadInConfig(); err != nil {
		log.Errorf("Error reading config file %v", err)
	}
	viper.WatchConfig()

	env.PostgreSQL.HOST = viper.GetString("POSTGRE_HOST")
	env.PostgreSQL.USER = viper.GetString("POSTGRE_USER")
	env.PostgreSQL.PORT = viper.GetString("POSTGRE_PORT")
	env.PostgreSQL.PASSWORD = viper.GetString("POSTGRE_PASSWORD")
	env.PostgreSQL.DATABASE_NAME = viper.GetString("POSTGRE_DATABASE")

	Config = env
}
