package apiserver

import (
	"strings"

	"github.com/spf13/viper"

	"github.com/rppkg/godfrey/internal/apiserver/dal"
	"github.com/rppkg/godfrey/internal/pkg/models"
	"github.com/rppkg/godfrey/pkg/db/mysql"
)

var cfg string

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
		Host:                  viper.GetString("APISERVER_DB_HOST"),
		Username:              viper.GetString("APISERVER_DB_USERNAME"),
		Password:              viper.GetString("APISERVER_DB_PASSWORD"),
		Database:              viper.GetString("APISERVER_DB_DATABASE"),
		MaxIdleConnections:    viper.GetInt("APISERVER_DB_MAX_IDLE_COMM"),
		MaxOpenConnections:    viper.GetInt("APISERVER_DB_MAX_OPEN_CONN"),
		MaxConnectionLifeTime: viper.GetDuration("APISERVER_DB_MAX_CONN_LIFE_TIME"),
		LogLevel:              viper.GetInt("APISERVER_DB_LOG_LEVEL"),
	}

	gormDB, err := mysql.InitDB(dbOptions)
	if err != nil {
		return err
	}

	// NOTE: maybe it's better to use atlas
	err = gormDB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&models.User{},
		&models.Role{},
	)
	if err != nil {
		return err
	}

	dal.InitDB(gormDB)

	return nil
}
