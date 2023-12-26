// Code generated by MockGen. DO NOT EDIT.
// Source: src/repository/user_repository.go
//
// Generated by this command:
//
//	mockgen -source=src/repository/user_repository.go -destination=test/mocks/user_repository_mock.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	custom_errors "github.com/lucasolsi-wex/go-crud/src/config/custom_errors"
	model "github.com/lucasolsi-wex/go-crud/src/model"
	gomock "go.uber.org/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserRepository) CreateUser(domainInterface model.UserDomainInterface) (model.UserDomainInterface, *custom_errors.CustomErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", domainInterface)
	ret0, _ := ret[0].(model.UserDomainInterface)
	ret1, _ := ret[1].(*custom_errors.CustomErr)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserRepositoryMockRecorder) CreateUser(domainInterface any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserRepository)(nil).CreateUser), domainInterface)
}

// ExistsByFirstNameAndLastName mocks base method.
func (m *MockUserRepository) ExistsByFirstNameAndLastName(firstName, lastName string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExistsByFirstNameAndLastName", firstName, lastName)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ExistsByFirstNameAndLastName indicates an expected call of ExistsByFirstNameAndLastName.
func (mr *MockUserRepositoryMockRecorder) ExistsByFirstNameAndLastName(firstName, lastName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExistsByFirstNameAndLastName", reflect.TypeOf((*MockUserRepository)(nil).ExistsByFirstNameAndLastName), firstName, lastName)
}

// FindUserById mocks base method.
func (m *MockUserRepository) FindUserById(id string) (model.UserDomainInterface, *custom_errors.CustomErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserById", id)
	ret0, _ := ret[0].(model.UserDomainInterface)
	ret1, _ := ret[1].(*custom_errors.CustomErr)
	return ret0, ret1
}

// FindUserById indicates an expected call of FindUserById.
func (mr *MockUserRepositoryMockRecorder) FindUserById(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserById", reflect.TypeOf((*MockUserRepository)(nil).FindUserById), id)
}
