package dal

import (
	"sync"

	"gorm.io/gorm"
)

var (
	once sync.Once
	dal  *Dal
)

type IDal interface {
	DB() *gorm.DB
}

type Dal struct {
	gdb *gorm.DB
}

var _ IDal = (*Dal)(nil)

func (ms *Dal) DB() *gorm.DB {
	return ms.gdb
}

func GetDal() *Dal {
	return dal
}

func InitDB(gdb *gorm.DB) {
	once.Do(func() {
		dal = &Dal{gdb}
	})
}
