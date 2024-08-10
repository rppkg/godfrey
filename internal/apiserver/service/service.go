package service

import (
	"github.com/rppkg/godfrey/internal/apiserver/dal"
	"github.com/rppkg/godfrey/internal/apiserver/service/user"
)

type IService interface {
	Users() user.IService
}

var _ IService = (*Service)(nil)

type Service struct {
	dal dal.IDal
}

func New(dal dal.IDal) *Service {
	return &Service{dal: dal}
}

func (s *Service) Users() user.IService {
	return user.New(s.dal)
}
