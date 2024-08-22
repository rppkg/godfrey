package user

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"github.com/rppkg/godfrey/internal/apiserver/dal"
	"github.com/rppkg/godfrey/internal/pkg/models"
)

func TestService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDal := dal.NewMockIDal(ctrl)
	mockUserDal := dal.NewMockIUserDal(ctrl)
	mockDal.EXPECT().Users().AnyTimes().Return(mockUserDal)

	u1 := &models.User{
		ID:        "id1",
		Username:  "username1",
		Nickname:  nil,
		Password:  "password1",
		Salt:      "salt1",
		Avatar:    "avatar1",
		Email:     "email1",
		Phone:     "phone1",
		RoleID:    "role_id1",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: gorm.DeletedAt{},
		Role:      nil,
	}

	mockUserDal.EXPECT().Get(gomock.Any(), "godfrey").Return(u1, nil)

	s := &Service{dal: mockDal}
	user, err := s.Get(context.Background(), "godfrey")
	assert.Nil(t, err)
	assert.Equal(t, user.Username, u1.Username)
}
