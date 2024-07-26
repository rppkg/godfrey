package godfrey

import (
	"strings"

	"github.com/spf13/viper"

	"github.com/rppkg/godfrey/internal/godfrey/dal"
	"github.com/rppkg/godfrey/pkg/db/mysql"
)

func initConfig() {
	if cfg != "" {
		viper.SetConfigFile(cfg)
	}

	viper.AutomaticEnv()
	viper.SetEnvPrefix("GODFREY")

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
}

func initDal() error {
	dbOptions := &mysql.Options{
		Host:                  viper.GetString("GF_SERVE_MYSQL_HOST"),
		Username:              viper.GetString("GF_SERVE_MYSQL_USERNAME"),
		Password:              viper.GetString("GF_SERVE_MYSQL_PASSWORD"),
		Database:              viper.GetString("GF_SERVE_MYSQL_DATABASE"),
		MaxIdleConnections:    viper.GetInt("GF_SERVE_MYSQL_MAX_IDLE_COMM"),
		MaxOpenConnections:    viper.GetInt("GF_SERVE_MYSQL_MAX_OPEN_CONN"),
		MaxConnectionLifeTime: viper.GetDuration("GF_SERVE_MYSQL_MAX_CONN_LIFE_TIME"),
		LogLevel:              viper.GetInt("GF_SERVE_MYSQL_GORM_LOG_LEVEL"),
	}

	err := mysql.Migrate(dbOptions)
	if err != nil {
		return err
	}

	gdb, err := mysql.InitDB(dbOptions)
	if err != nil {
		return err
	}
	dal.InitDB(gdb)

	return nil
}
