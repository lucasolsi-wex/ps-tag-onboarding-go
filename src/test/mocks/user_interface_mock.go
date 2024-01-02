// Code generated by MockGen. DO NOT EDIT.
// Source: src/model/service/user_interface.go
//
// Generated by this command:
//
//	mockgen -source=src/model/service/user_interface.go -destination=src/test/mocks/user_interface_mock.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	custom_errors "github.com/lucasolsi-wex/go-crud/src/config/custom_errors"
	model "github.com/lucasolsi-wex/go-crud/src/model"
	gomock "go.uber.org/mock/gomock"
)

// MockUserDomainService is a mock of UserDomainService interface.
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
func (m *MockUserDomainService) CreateUser(domainInterface model.UserDomainInterface) (model.UserDomainInterface, *custom_errors.CustomErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", domainInterface)
	ret0, _ := ret[0].(model.UserDomainInterface)
	ret1, _ := ret[1].(*custom_errors.CustomErr)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserDomainServiceMockRecorder) CreateUser(domainInterface any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserDomainService)(nil).CreateUser), domainInterface)
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
func (m *MockUserDomainService) FindUserById(id string) (model.UserDomainInterface, *custom_errors.CustomErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserById", id)
	ret0, _ := ret[0].(model.UserDomainInterface)
	ret1, _ := ret[1].(*custom_errors.CustomErr)
	return ret0, ret1
}

// FindUserById indicates an expected call of FindUserById.
func (mr *MockUserDomainServiceMockRecorder) FindUserById(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserById", reflect.TypeOf((*MockUserDomainService)(nil).FindUserById), id)
}