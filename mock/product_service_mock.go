// Code generated by MockGen. DO NOT EDIT.
// Source: product_service.go

// Package mock is a generated GoMock package.
package mock

import (
	model "mutlicontainer/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	echo "github.com/labstack/echo/v4"
)

// MockProductsService is a mock of ProductsService interface.
type MockProductsService struct {
	ctrl     *gomock.Controller
	recorder *MockProductsServiceMockRecorder
}

// MockProductsServiceMockRecorder is the mock recorder for MockProductsService.
type MockProductsServiceMockRecorder struct {
	mock *MockProductsService
}

// NewMockProductsService creates a new mock instance.
func NewMockProductsService(ctrl *gomock.Controller) *MockProductsService {
	mock := &MockProductsService{ctrl: ctrl}
	mock.recorder = &MockProductsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductsService) EXPECT() *MockProductsServiceMockRecorder {
	return m.recorder
}

// CreateProduct mocks base method.
func (m *MockProductsService) CreateProduct(ctx echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProduct", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateProduct indicates an expected call of CreateProduct.
func (mr *MockProductsServiceMockRecorder) CreateProduct(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProduct", reflect.TypeOf((*MockProductsService)(nil).CreateProduct), ctx)
}

// GetDetails mocks base method.
func (m *MockProductsService) GetDetails(ctx echo.Context) (model.ProductDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDetails", ctx)
	ret0, _ := ret[0].(model.ProductDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDetails indicates an expected call of GetDetails.
func (mr *MockProductsServiceMockRecorder) GetDetails(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDetails", reflect.TypeOf((*MockProductsService)(nil).GetDetails), ctx)
}
