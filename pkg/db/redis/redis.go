package redis

import (
	"context"
	"time"

	goredis "github.com/redis/go-redis/v9"
)

type Options struct {
	Addr         string
	Username     string
	Password     string
	Database     int
	MaxRetries   int
	MinIdleConns int
	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	PoolTimeout  time.Duration
	PoolSize     int
}

func Init(opts *Options) (*goredis.Client, error) {
	options := &goredis.Options{
		Addr:         opts.Addr,
		Username:     opts.Username,
		Password:     opts.Password,
		DB:           opts.Database,
		MaxRetries:   opts.MaxRetries,
		MinIdleConns: opts.MinIdleConns,
		DialTimeout:  opts.DialTimeout,
		ReadTimeout:  opts.ReadTimeout,
		WriteTimeout: opts.WriteTimeout,
		PoolTimeout:  opts.PoolTimeout,
		PoolSize:     opts.PoolSize,
	}

	rdb := goredis.NewClient(options)

	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		return nil, err
	}

	return rdb, nil
}