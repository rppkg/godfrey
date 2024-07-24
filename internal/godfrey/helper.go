package godfrey

import (
	"strings"

	"github.com/rppkg/godfrey/pkg/db"
	"github.com/spf13/viper"
)

const (
	HomeConfigPathDir = ".godfrey"
	DefaultConfigName = "godfrey.yaml"
)

func initConfig() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("GODFREY")

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
}

func initDal() error {
	dbOptions := &db.MySQLOptions{
		Host:                  viper.GetString("GF_SERVE_MYSQL_HOST"),
		Username:              viper.GetString("GF_SERVE_MYSQL_USERNAME"),
		Password:              viper.GetString("GF_SERVE_MYSQL_PASSWORD"),
		Database:              viper.GetString("GF_SERVE_MYSQL_DATABASE"),
		MaxIdleConnections:    viper.GetInt("GF_SERVE_MYSQL_MAX_IDLE_COMM"),
		MaxOpenConnections:    viper.GetInt("GF_SERVE_MYSQL_MAX_OPEN_CONN"),
		MaxConnectionLifeTime: viper.GetDuration("GF_SERVE_MYSQL_MAX_CONN_LIFE_TIME"),
		LogLevel:              viper.GetInt("GF_SERVE_MYSQL_GORM_LOG_LEVEL"),
	}

	err := db.Migrate(dbOptions)
	if err != nil {
		return err
	}

	// ins, err := db.NewMySQL(dbOptions)
	// if err != nil {
	// 	return err
	// }

	// store.Init(ins)

	return nil
}
