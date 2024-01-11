// Code generated by MockGen. DO NOT EDIT.
// Source: internal/service/service_interfaces.go
//
// Generated by this command:
//
//	mockgen --source=internal/service/service_interfaces.go --destination=mocks/services_mock.go --package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	models "github.com/GabriellGds/go-orders/internal/models"
	errors "github.com/GabriellGds/go-orders/pkg/errors"
	gomock "go.uber.org/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// CreateItemService mocks base method.
func (m *MockService) CreateItemService(arg0 context.Context, arg1 *models.Item) (*models.Item, *errors.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateItemService", arg0, arg1)
	ret0, _ := ret[0].(*models.Item)
	ret1, _ := ret[1].(*errors.ErrorResponse)
	return ret0, ret1
}

// CreateItemService indicates an expected call of CreateItemService.
func (mr *MockServiceMockRecorder) CreateItemService(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateItemService", reflect.TypeOf((*MockService)(nil).CreateItemService), arg0, arg1)
}

// CreateOrderService mocks base method.
func (m *MockService) CreateOrderService(arg0 context.Context, arg1 *models.Order) (*models.Order, *errors.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrderService", arg0, arg1)
	ret0, _ := ret[0].(*models.Order)
	ret1, _ := ret[1].(*errors.ErrorResponse)
	return ret0, ret1
}

// CreateOrderService indicates an expected call of CreateOrderService.
func (mr *MockServiceMockRecorder) CreateOrderService(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrderService", reflect.TypeOf((*MockService)(nil).CreateOrderService), arg0, arg1)
}

// CreateUserService mocks base method.
func (m *MockService) CreateUserService(arg0 context.Context, arg1 *models.User) (*models.User, *errors.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserService", arg0, arg1)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(*errors.ErrorResponse)
	return ret0, ret1
}

// CreateUserService indicates an expected call of CreateUserService.
func (mr *MockServiceMockRecorder) CreateUserService(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserService", reflect.TypeOf((*MockService)(nil).CreateUserService), arg0, arg1)
}

// DeleteItemService mocks base method.
func (m *MockService) DeleteItemService(arg0 context.Context, arg1 int) *errors.ErrorResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteItemService", arg0, arg1)
	ret0, _ := ret[0].(*errors.ErrorResponse)
	return ret0
}

// DeleteItemService indicates an expected call of DeleteItemService.
func (mr *MockServiceMockRecorder) DeleteItemService(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteItemService", reflect.TypeOf((*MockService)(nil).DeleteItemService), arg0, arg1)
}

// DeleteOrderService mocks base method.
func (m *MockService) DeleteOrderService(ctx context.Context, userID, id int) *errors.ErrorResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrderService", ctx, userID, id)
	ret0, _ := ret[0].(*errors.ErrorResponse)
	return ret0
}

// DeleteOrderService indicates an expected call of DeleteOrderService.
func (mr *MockServiceMockRecorder) DeleteOrderService(ctx, userID, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrderService", reflect.TypeOf((*MockService)(nil).DeleteOrderService), ctx, userID, id)
}

// DeleteUserService mocks base method.
func (m *MockService) DeleteUserService(arg0 context.Context, arg1 int) *errors.ErrorResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserService", arg0, arg1)
	ret0, _ := ret[0].(*errors.ErrorResponse)
	return ret0
}

// DeleteUserService indicates an expected call of DeleteUserService.
func (mr *MockServiceMockRecorder) DeleteUserService(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserService", reflect.TypeOf((*MockService)(nil).DeleteUserService), arg0, arg1)
}

// FindItemService mocks base method.
func (m *MockService) FindItemService(arg0 context.Context, arg1 int) (models.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindItemService", arg0, arg1)
	ret0, _ := ret[0].(models.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindItemService indicates an expected call of FindItemService.
func (mr *MockServiceMockRecorder) FindItemService(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindItemService", reflect.TypeOf((*MockService)(nil).FindItemService), arg0, arg1)
}

// FindOrderService mocks base method.
func (m *MockService) FindOrderService(arg0 context.Context, arg1 int) (models.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrderService", arg0, arg1)
	ret0, _ := ret[0].(models.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrderService indicates an expected call of FindOrderService.
func (mr *MockServiceMockRecorder) FindOrderService(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrderService", reflect.TypeOf((*MockService)(nil).FindOrderService), arg0, arg1)
}

// FindUserService mocks base method.
func (m *MockService) FindUserService(arg0 context.Context, arg1 int) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserService", arg0, arg1)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserService indicates an expected call of FindUserService.
func (mr *MockServiceMockRecorder) FindUserService(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserService", reflect.TypeOf((*MockService)(nil).FindUserService), arg0, arg1)
}

// ListItems mocks base method.
func (m *MockService) ListItems(ctx context.Context) ([]models.Item, *errors.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListItems", ctx)
	ret0, _ := ret[0].([]models.Item)
	ret1, _ := ret[1].(*errors.ErrorResponse)
	return ret0, ret1
}

// ListItems indicates an expected call of ListItems.
func (mr *MockServiceMockRecorder) ListItems(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListItems", reflect.TypeOf((*MockService)(nil).ListItems), ctx)
}

// ListUsers mocks base method.
func (m *MockService) ListUsers(ctx context.Context) ([]models.User, *errors.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsers", ctx)
	ret0, _ := ret[0].([]models.User)
	ret1, _ := ret[1].(*errors.ErrorResponse)
	return ret0, ret1
}

// ListUsers indicates an expected call of ListUsers.
func (mr *MockServiceMockRecorder) ListUsers(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsers", reflect.TypeOf((*MockService)(nil).ListUsers), ctx)
}

// Login mocks base method.
func (m *MockService) Login(arg0 context.Context, arg1 *models.User) (*models.User, string, *errors.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", arg0, arg1)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(*errors.ErrorResponse)
	return ret0, ret1, ret2
}

// Login indicates an expected call of Login.
func (mr *MockServiceMockRecorder) Login(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockService)(nil).Login), arg0, arg1)
}

// UpdateItemSvice mocks base method.
func (m *MockService) UpdateItemSvice(arg0 context.Context, arg1 int, arg2 *models.Item) *errors.ErrorResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateItemSvice", arg0, arg1, arg2)
	ret0, _ := ret[0].(*errors.ErrorResponse)
	return ret0
}

// UpdateItemSvice indicates an expected call of UpdateItemSvice.
func (mr *MockServiceMockRecorder) UpdateItemSvice(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateItemSvice", reflect.TypeOf((*MockService)(nil).UpdateItemSvice), arg0, arg1, arg2)
}

// UpdateUserService mocks base method.
func (m *MockService) UpdateUserService(arg0 context.Context, arg1 int, arg2 *models.User) *errors.ErrorResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserService", arg0, arg1, arg2)
	ret0, _ := ret[0].(*errors.ErrorResponse)
	return ret0
}

// UpdateUserService indicates an expected call of UpdateUserService.
func (mr *MockServiceMockRecorder) UpdateUserService(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserService", reflect.TypeOf((*MockService)(nil).UpdateUserService), arg0, arg1, arg2)
}