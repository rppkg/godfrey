package db

import (
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	"gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	migratemysql "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rppkg/godfrey/pkg/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm/logger"
)

type MySQLOptions struct {
	Host                  string
	Username              string
	Password              string
	Database              string
	MaxIdleConnections    int
	MaxOpenConnections    int
	MaxConnectionLifeTime time.Duration
	LogLevel              int
}

func (o *MySQLOptions) DSN() string {
	return fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s`,
		o.Username,
		o.Password,
		o.Host,
		o.Database,
		true,
		"Local")
}

func NewMySQL(opts *MySQLOptions) (*gorm.DB, error) {
	logLevel := logger.Silent
	if opts.LogLevel != 0 {
		logLevel = logger.LogLevel(opts.LogLevel)
	}
	db, err := gorm.Open(mysql.Open(opts.DSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(opts.MaxOpenConnections)
	sqlDB.SetConnMaxLifetime(opts.MaxConnectionLifeTime)
	sqlDB.SetMaxIdleConns(opts.MaxIdleConnections)

	return db, nil
}

func Migrate(opts *MySQLOptions) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?multiStatements=true",
		opts.Username, opts.Password, opts.Host, opts.Database,
	)

	log.Info("dsn", "msg", dsn)

	db, _ := sql.Open("mysql", dsn)
	driver, _ := migratemysql.WithInstance(db, &migratemysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		opts.Database,
		driver,
	)

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Error("migrate up failed", slog.Any("error", err))
		return fmt.Errorf("migration failed: %w", err)
	}

	return nil
}
