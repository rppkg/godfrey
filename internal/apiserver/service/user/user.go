package user

import (
	"context"

	"github.com/rppkg/godfrey/internal/pkg/core"

	"github.com/rppkg/godfrey/internal/apiserver/dal"
	"github.com/rppkg/godfrey/internal/pkg/auth"
	v1 "github.com/rppkg/godfrey/pkg/api/v1"
	"github.com/rppkg/godfrey/pkg/token"
)

type IService interface {
	Login(ctx context.Context, r *v1.LoginUserRequest) (*v1.LoginUserResponse, error)
}

type Service struct {
	dal dal.IDal
}

var _ IService = (*Service)(nil)

func New(dal dal.IDal) *Service {
	return &Service{dal: dal}
}

func (s *Service) Login(ctx context.Context, r *v1.LoginUserRequest) (*v1.LoginUserResponse, error) {
	user, err := s.dal.Users().Get(ctx, r.Username)
	if err != nil {
		return nil, core.HTTP500.SetMessage(err.Error())
	}

	if user.Password != auth.SignPwdWithSalt(r.Password, user.Salt) {
		return nil, core.HTTP500.SetMessage("密码错误")
	}

	tokenStr, err := token.Sign(r.Username)
	if err != nil {
		return nil, core.HTTP500.SetMessage(err.Error())
	}

	return &v1.LoginUserResponse{Token: tokenStr}, nil
}
