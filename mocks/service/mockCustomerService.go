// Code generated by MockGen. DO NOT EDIT.
// Source: red/service (interfaces: CustomerService)

// Package service is a generated GoMock package.
package service

import (
	dto "red/dto"
	errs "red/errs"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCustomerService is a mock of CustomerService interface.
type MockCustomerService struct {
	ctrl     *gomock.Controller
	recorder *MockCustomerServiceMockRecorder
}

// MockCustomerServiceMockRecorder is the mock recorder for MockCustomerService.
type MockCustomerServiceMockRecorder struct {
	mock *MockCustomerService
}

// NewMockCustomerService creates a new mock instance.
func NewMockCustomerService(ctrl *gomock.Controller) *MockCustomerService {
	mock := &MockCustomerService{ctrl: ctrl}
	mock.recorder = &MockCustomerServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCustomerService) EXPECT() *MockCustomerServiceMockRecorder {
	return m.recorder
}

// GetAllCustomer mocks base method.
func (m *MockCustomerService) GetAllCustomer(arg0 string) ([]dto.CustomerResponse, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllCustomer", arg0)
	ret0, _ := ret[0].([]dto.CustomerResponse)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// GetAllCustomer indicates an expected call of GetAllCustomer.
func (mr *MockCustomerServiceMockRecorder) GetAllCustomer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCustomer", reflect.TypeOf((*MockCustomerService)(nil).GetAllCustomer), arg0)
}

// GetCustomer mocks base method.
func (m *MockCustomerService) GetCustomer(arg0 string) (*dto.CustomerResponse, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCustomer", arg0)
	ret0, _ := ret[0].(*dto.CustomerResponse)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// GetCustomer indicates an expected call of GetCustomer.
func (mr *MockCustomerServiceMockRecorder) GetCustomer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCustomer", reflect.TypeOf((*MockCustomerService)(nil).GetCustomer), arg0)
}

// SaveCustomer mocks base method.
func (m *MockCustomerService) SaveCustomer(arg0 dto.CustomerRequest) (*dto.CustomerResponse, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveCustomer", arg0)
	ret0, _ := ret[0].(*dto.CustomerResponse)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// SaveCustomer indicates an expected call of SaveCustomer.
func (mr *MockCustomerServiceMockRecorder) SaveCustomer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveCustomer", reflect.TypeOf((*MockCustomerService)(nil).SaveCustomer), arg0)
}
