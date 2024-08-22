package dal

import (
	"sync"

	"gorm.io/gorm"

	"github.com/rppkg/godfrey/internal/apiserver/dal/query"
)

var (
	once sync.Once
	d    IDal
)

//go:generate mockgen -self_package=github.com/rppkg/godfrey/internal/apiserver/dal -destination mock_dal.go -package dal github.com/rppkg/godfrey/internal/apiserver/dal IDal,IUserDal

type IDal interface {
	DB() *gorm.DB
	DBQuery() *query.Query
	Users() IUserDal
}

type D struct {
	db *gorm.DB
	q  *query.Query
}

var _ IDal = (*D)(nil)

func (d *D) DB() *gorm.DB {
	return d.db
}

func (d *D) DBQuery() *query.Query {
	return d.q
}

func Cli() IDal {
	return d
}

func Init(db *gorm.DB) {
	once.Do(func() {
		d = &D{db, query.Use(db)}
	})
}

func (d *D) Users() IUserDal {
	return NewUserDal(d.q)
}
