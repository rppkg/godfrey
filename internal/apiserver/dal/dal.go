package dal

import (
	"sync"

	"gorm.io/gorm"
)

var (
	once   sync.Once
	client IClient
)

type IClient interface {
	DB() *gorm.DB
}

type Client struct {
	gdb *gorm.DB
}

var _ IClient = (*Client)(nil)

func (ms *Client) DB() *gorm.DB {
	return ms.gdb
}

func Cli() IClient {
	return client
}

func InitDB(gdb *gorm.DB) {
	once.Do(func() {
		client = &Client{gdb}
	})
}
