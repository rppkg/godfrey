package user

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/rppkg/godfrey/internal/apiserver/dal"
	v1 "github.com/rppkg/godfrey/pkg/api/v1"
)

func BenchmarkService_Get(b *testing.B) {
	// Benchmark test
}

func TestNewService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		dal dal.IDal
	}
	tests := []struct {
		name string
		args args
		want IService
	}{
		{
			name: "TestNewService",
			args: args{dal: dal.NewMockInitDal(ctrl)},
			want: NewService(dal.NewMockInitDal(ctrl)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(tt.args.dal); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserDal := dal.NewMockUserDal(ctrl)

	mockdal := dal.NewMockInitDal(ctrl)
	mockdal.EXPECT().Users().Return(mockUserDal)

	s := &Service{
		dal: mockdal,
	}

	req := &v1.CreateUserRequest{
		Username: "x1",
		Nickname: "x1",
		Email:    "abc@gmail.com",
		Phone:    "18888888888",
		Password: "xxxx",
	}

	got, err := s.Create(context.Background(), req)
	assert.Nil(t, err)
	assert.Equal(t, got, req)
}
