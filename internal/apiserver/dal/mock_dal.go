package dal

import (
	"context"
	"reflect"

	"github.com/golang/mock/gomock"
	"gorm.io/gorm"

	"github.com/rppkg/godfrey/internal/apiserver/dal/query"
	"github.com/rppkg/godfrey/internal/pkg/models"
)

type MockDal struct {
	ctrl     *gomock.Controller
	recorder *MockDalRecorder
}

type MockDalRecorder struct {
	mock *MockDal
}

func NewMockInitDal(ctrl *gomock.Controller) *MockDal {
	mock := &MockDal{ctrl: ctrl}
	mock.recorder = &MockDalRecorder{mock}
	return mock
}

func (m *MockDal) EXPECT() *MockDalRecorder {
	return m.recorder
}

func (m *MockDal) DB() *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DB")
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *MockDalRecorder) DB() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DB", reflect.TypeOf((*MockDal)(nil).DB))
}

func (m *MockDal) DBQuery() *query.Query {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DBQuery")
	ret0, _ := ret[0].(*query.Query)
	return ret0
}

func (mr *MockDalRecorder) DBQuery() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DBQuery", reflect.TypeOf((*MockDal)(nil).DBQuery))
}

func (m *MockDal) Users() IUserDal {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Users")
	ret0, _ := ret[0].(IUserDal)
	return ret0
}

func (mr *MockDalRecorder) Users() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Users", reflect.TypeOf((*MockDal)(nil).Users))
}

type MockUserDal struct {
	ctrl     *gomock.Controller
	recorder *MockUserDalRecorder
}

type MockUserDalRecorder struct {
	mock *MockUserDal
}

func NewMockUserDal(ctrl *gomock.Controller) *MockUserDal {
	mock := &MockUserDal{ctrl: ctrl}
	mock.recorder = &MockUserDalRecorder{mock}
	return mock
}

func (m *MockUserDal) EXPECT() *MockUserDalRecorder {
	return m.recorder
}

func (m *MockUserDal) Create(arg0 context.Context, arg1 *models.User) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockUserDalRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserDal)(nil).Create), arg0, arg1)
}

func (m *MockUserDal) Get(arg0 context.Context, arg1 string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockUserDalRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUserDal)(nil).Get), arg0, arg1)
}

func (m *MockUserDal) List(arg0 context.Context, arg1, arg2 int) ([]*models.User, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*models.User)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (mr *MockUserDalRecorder) List(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockUserDal)(nil).List), arg0, arg1, arg2)
}

func (m *MockUserDal) Update(arg0 context.Context, arg1 *models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockUserDalRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserDal)(nil).Update), arg0, arg1)
}

func (m *MockUserDal) Delete(arg0 context.Context, arg1 *models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockUserDalRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUserDal)(nil).Delete), arg0, arg1)
}
