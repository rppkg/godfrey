package dal

import (
	"context"

	"github.com/rppkg/godfrey/internal/apiserver/dal/query"
	"github.com/rppkg/godfrey/internal/pkg/models"
)

type IUserDal interface {
	Get(ctx context.Context, username string) (*models.User, error)
}

type UserDal struct {
	q *query.Query
}

var _ IUserDal = (*UserDal)(nil)

func newUsers(q *query.Query) *UserDal {
	return &UserDal{q}
}

func (ud *UserDal) Get(ctx context.Context, username string) (*models.User, error) {
	Q := ud.q.User.WithContext(ctx)

	user, err := Q.Where(ud.q.User.Username.Eq(username)).First()
	if err != nil {
		return nil, err
	}

	return user, nil
}
