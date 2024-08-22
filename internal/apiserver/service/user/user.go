package user

import (
	"context"
	"errors"
	"log/slog"

	"github.com/rppkg/godfrey/internal/pkg/models"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"github.com/rppkg/godfrey/internal/apiserver/dal"
	"github.com/rppkg/godfrey/internal/pkg/auth"
	"github.com/rppkg/godfrey/internal/pkg/core"
	v1 "github.com/rppkg/godfrey/pkg/api/v1"
	"github.com/rppkg/godfrey/pkg/log"
	"github.com/rppkg/godfrey/pkg/token"
)

//go:generate mockgen -destination mock/mock_user.go -package user github.com/rppkg/godfrey/internal/apiserver/service/user IService

type IService interface {
	Create(ctx context.Context, r *v1.CreateUserRequest) (*v1.CreateUserResponse, error)
	Login(ctx context.Context, r *v1.LoginUserRequest) (*v1.LoginUserResponse, error)
	Get(ctx context.Context, username string) (*v1.GetUserResponse, error)
	List(ctx context.Context, r *v1.ListUserRequest) (*v1.ListUserResponse, error)
	Update(ctx context.Context, username string, r *v1.UpdateUserRequest) (*v1.UpdateUserResponse, error)
	Delete(ctx context.Context, username string) (*v1.DeleteUserResponse, error)
}

type Service struct {
	dal dal.IDal
}

var _ IService = (*Service)(nil)

func NewService(dal dal.IDal) IService {
	return &Service{dal: dal}
}

func (s *Service) Create(ctx context.Context, r *v1.CreateUserRequest) (*v1.CreateUserResponse, error) {
	var user models.User
	_ = copier.Copy(&user, r)

	item, err := s.dal.Users().Create(ctx, &user)
	if err != nil {
		log.Error("创建用户错误", slog.Any("error", err))
		return nil, core.HTTP500.SetMessage("创建用户错误")
	}

	var resp v1.CreateUserResponse
	_ = copier.Copy(&resp, item)

	return &resp, nil
}

func (s *Service) Login(ctx context.Context, r *v1.LoginUserRequest) (*v1.LoginUserResponse, error) {
	user, err := s.dal.Users().Get(ctx, r.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, core.HTTP404.SetMessage("不存的用户")
		}
		log.Error("获取用户错误", slog.Any("error", err))
		return nil, core.HTTP500.SetMessage("获取用户错误")
	}

	if user.Password != auth.SignPwdWithSalt(r.Password, user.Salt) {
		log.Error("密码错误", slog.Any("error", err))
		return nil, core.HTTP500.SetMessage("密码错误")
	}

	tokenStr, err := token.Sign(r.Username)
	if err != nil {
		log.Error("签发Token错误", slog.Any("error", err))
		return nil, core.HTTP500.SetMessage("签发Token错误")
	}

	return &v1.LoginUserResponse{Token: tokenStr}, nil
}

func (s *Service) Get(ctx context.Context, username string) (*v1.GetUserResponse, error) {
	item, err := s.dal.Users().Get(ctx, username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, core.HTTP404.SetMessage("不存的用户")
		}

		log.Error("获取用户错误", slog.Any("error", err))
		return nil, core.HTTP500.SetMessage("获取用户错误")
	}

	var resp v1.GetUserResponse
	_ = copier.Copy(&resp, item)

	resp.CreatedAt = item.CreatedAt.Format("2006-01-02 15:04:05")
	resp.UpdatedAt = item.UpdatedAt.Format("2006-01-02 15:04:05")

	return &resp, nil
}

func (s *Service) List(ctx context.Context, r *v1.ListUserRequest) (*v1.ListUserResponse, error) {
	items, total, err := s.dal.Users().List(ctx, r.Offset, r.Limit)
	if err != nil {
		log.Error("获取用户列表错误", slog.Any("error", err))
		return nil, core.HTTP500.SetMessage("获取用户列表错误")
	}

	users := make([]*v1.User, 0, len(items))
	for _, item := range items {
		var u v1.User
		_ = copier.Copy(&u, item)
		u.CreatedAt = item.CreatedAt.Format("2006-01-02 15:04:05")
		u.UpdatedAt = item.UpdatedAt.Format("2006-01-02 15:04:05")

		users = append(users, &u)
	}

	return &v1.ListUserResponse{Users: users, Total: total}, nil
}

func (s *Service) Update(ctx context.Context, username string, r *v1.UpdateUserRequest) (*v1.UpdateUserResponse, error) {
	item, err := s.dal.Users().Get(ctx, username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, core.HTTP404.SetMessage("不存的用户")
		}

		log.Error("获取用户错误", slog.Any("error", err))
		return nil, core.HTTP500.SetMessage("获取用户错误")
	}

	if r.Nickname != nil {
		*item.Nickname = *r.Nickname
	}
	if r.Email != nil {
		item.Email = *r.Email
	}
	if r.Phone != nil {
		item.Phone = *r.Phone
	}

	err = s.dal.Users().Update(ctx, item)
	if err != nil {
		log.Error("更新用户错误", slog.Any("error", err))
		return nil, core.HTTP500.SetMessage("更新用户错误")
	}

	var resp v1.UpdateUserResponse
	_ = copier.Copy(&resp, item)

	return &resp, nil
}

func (s *Service) Delete(ctx context.Context, username string) (*v1.DeleteUserResponse, error) {
	item, err := s.dal.Users().Get(ctx, username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, core.HTTP404.SetMessage("不存的用户")
		}

		log.Error("获取用户错误", slog.Any("error", err))
		return nil, core.HTTP500.SetMessage("获取用户错误")
	}

	err = s.dal.Users().Delete(ctx, item)
	if err != nil {
		log.Error("删除用户错误", slog.Any("error", err))
		return nil, core.HTTP500.SetMessage("删除用户错误")
	}

	var resp v1.DeleteUserResponse
	_ = copier.Copy(&resp, item)

	return &resp, nil
}
