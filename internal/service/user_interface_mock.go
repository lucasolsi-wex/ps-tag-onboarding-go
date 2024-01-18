// Code generated by MockGen. DO NOT EDIT.
// Source: internal/service/user_interface.go
//
// Generated by this command:
//
//	mockgen -source=internal/service/user_interface.go -destination=internal/service/user_interface_mock.go -package=service
//

// Package service is a generated GoMock package.
package service

import (
	reflect "reflect"

	models "github.com/lucasolsi-wex/go-crud/internal/models"
	gomock "go.uber.org/mock/gomock"
)

// MockUserDomainService is a mock of UserInterfaceService interface.
type MockUserDomainService struct {
	ctrl     *gomock.Controller
	recorder *MockUserDomainServiceMockRecorder
}

// MockUserDomainServiceMockRecorder is the mock recorder for MockUserDomainService.
type MockUserDomainServiceMockRecorder struct {
	mock *MockUserDomainService
}

// NewMockUserDomainService creates a new mock instance.
func NewMockUserDomainService(ctrl *gomock.Controller) *MockUserDomainService {
	mock := &MockUserDomainService{ctrl: ctrl}
	mock.recorder = &MockUserDomainServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserDomainService) EXPECT() *MockUserDomainServiceMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserDomainService) CreateUser(request models.UserRequest) (*models.UserResponse, *models.CustomErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", request)
	ret0, _ := ret[0].(*models.UserResponse)
	ret1, _ := ret[1].(*models.CustomErr)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserDomainServiceMockRecorder) CreateUser(request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserDomainService)(nil).CreateUser), request)
}

// ExistsByFirstNameAndLastName mocks base method.
func (m *MockUserDomainService) ExistsByFirstNameAndLastName(firstName, lastName string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExistsByFirstNameAndLastName", firstName, lastName)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ExistsByFirstNameAndLastName indicates an expected call of ExistsByFirstNameAndLastName.
func (mr *MockUserDomainServiceMockRecorder) ExistsByFirstNameAndLastName(firstName, lastName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExistsByFirstNameAndLastName", reflect.TypeOf((*MockUserDomainService)(nil).ExistsByFirstNameAndLastName), firstName, lastName)
}

// FindUserById mocks base method.
func (m *MockUserDomainService) FindUserById(id string) (models.UserResponse, *models.CustomErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserById", id)
	ret0, _ := ret[0].(models.UserResponse)
	ret1, _ := ret[1].(*models.CustomErr)
	return ret0, ret1
}

// FindUserById indicates an expected call of FindUserById.
func (mr *MockUserDomainServiceMockRecorder) FindUserById(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserById", reflect.TypeOf((*MockUserDomainService)(nil).FindUserById), id)
}
