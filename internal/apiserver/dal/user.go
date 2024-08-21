package dal

import (
	"context"

	"github.com/rppkg/godfrey/internal/apiserver/dal/query"
	"github.com/rppkg/godfrey/internal/pkg/models"
)

type IUserDal interface {
	Create(ctx context.Context, user *models.User) (*models.User, error)
	Get(ctx context.Context, username string) (*models.User, error)
	List(ctx context.Context, offset, limit int) ([]*models.User, int64, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, user *models.User) error
}

type UserDal struct {
	q *query.Query
}

var _ IUserDal = (*UserDal)(nil)

func NewUserDal(q *query.Query) IUserDal {
	return &UserDal{q}
}

func (ud *UserDal) Create(ctx context.Context, user *models.User) (*models.User, error) {
	err := ud.q.User.WithContext(ctx).Create(user)
	if err != nil {
		return nil, err
	}

	item, err := ud.Get(ctx, user.Username)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (ud *UserDal) Get(ctx context.Context, username string) (*models.User, error) {
	user, err := ud.q.User.WithContext(ctx).Where(ud.q.User.Username.Eq(username)).First()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ud *UserDal) List(ctx context.Context, offset, limit int) ([]*models.User, int64, error) {
	list, total, err := ud.q.User.WithContext(ctx).Order(ud.q.User.ID.Desc()).FindByPage(offset, limit)
	if err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func (ud *UserDal) Update(ctx context.Context, user *models.User) error {
	if err := ud.q.User.WithContext(ctx).Save(user); err != nil {
		return err
	}

	return nil
}

func (ud *UserDal) Delete(ctx context.Context, user *models.User) error {
	info, err := ud.q.User.WithContext(ctx).Delete(user)
	if err != nil {
		return err
	}

	if info.Error != nil {
		return info.Error
	}

	return nil
}
