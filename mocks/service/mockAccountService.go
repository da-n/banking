// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/da-n/banking/service (interfaces: AccountService)

// Package service is a generated GoMock package.
package service

import (
	reflect "reflect"

	errs "github.com/da-n/banking-lib/errs"
	dto "github.com/da-n/banking/dto"
	gomock "github.com/golang/mock/gomock"
)

// MockAccountService is a mock of AccountService interface.
type MockAccountService struct {
	ctrl     *gomock.Controller
	recorder *MockAccountServiceMockRecorder
}

// MockAccountServiceMockRecorder is the mock recorder for MockAccountService.
type MockAccountServiceMockRecorder struct {
	mock *MockAccountService
}

// NewMockAccountService creates a new mock instance.
func NewMockAccountService(ctrl *gomock.Controller) *MockAccountService {
	mock := &MockAccountService{ctrl: ctrl}
	mock.recorder = &MockAccountServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountService) EXPECT() *MockAccountServiceMockRecorder {
	return m.recorder
}

// MakeTransaction mocks base method.
func (m *MockAccountService) MakeTransaction(arg0 dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeTransaction", arg0)
	ret0, _ := ret[0].(*dto.TransactionResponse)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// MakeTransaction indicates an expected call of MakeTransaction.
func (mr *MockAccountServiceMockRecorder) MakeTransaction(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeTransaction", reflect.TypeOf((*MockAccountService)(nil).MakeTransaction), arg0)
}

// NewAccount mocks base method.
func (m *MockAccountService) NewAccount(arg0 dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewAccount", arg0)
	ret0, _ := ret[0].(*dto.NewAccountResponse)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// NewAccount indicates an expected call of NewAccount.
func (mr *MockAccountServiceMockRecorder) NewAccount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAccount", reflect.TypeOf((*MockAccountService)(nil).NewAccount), arg0)
}