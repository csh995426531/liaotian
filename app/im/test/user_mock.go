// Code generated by MockGen. DO NOT EDIT.
// Source: ./domain/user/proto/user.pb.micro.go

// Package handler is a generated GoMock package.
package test

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	client "github.com/micro/go-micro/client"
	proto "liaotian/domain/user/proto"
	reflect "reflect"
)

// MockUserService is a mock of UserService interface
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// CreateUserInfo mocks base method
func (m *MockUserService) CreateUserInfo(ctx context.Context, in *proto.Request, opts ...client.CallOption) (*proto.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateUserInfo", varargs...)
	ret0, _ := ret[0].(*proto.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUserInfo indicates an expected call of CreateUserInfo
func (mr *MockUserServiceMockRecorder) CreateUserInfo(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserInfo", reflect.TypeOf((*MockUserService)(nil).CreateUserInfo), varargs...)
}

// GetUserInfo mocks base method
func (m *MockUserService) GetUserInfo(ctx context.Context, in *proto.Request, opts ...client.CallOption) (*proto.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUserInfo", varargs...)
	ret0, _ := ret[0].(*proto.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserInfo indicates an expected call of GetUserInfo
func (mr *MockUserServiceMockRecorder) GetUserInfo(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserInfo", reflect.TypeOf((*MockUserService)(nil).GetUserInfo), varargs...)
}

// UpdateUserInfo mocks base method
func (m *MockUserService) UpdateUserInfo(ctx context.Context, in *proto.Request, opts ...client.CallOption) (*proto.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateUserInfo", varargs...)
	ret0, _ := ret[0].(*proto.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserInfo indicates an expected call of UpdateUserInfo
func (mr *MockUserServiceMockRecorder) UpdateUserInfo(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserInfo", reflect.TypeOf((*MockUserService)(nil).UpdateUserInfo), varargs...)
}

// CheckUserPwd mocks base method
func (m *MockUserService) CheckUserPwd(ctx context.Context, in *proto.Request, opts ...client.CallOption) (*proto.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CheckUserPwd", varargs...)
	ret0, _ := ret[0].(*proto.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckUserPwd indicates an expected call of CheckUserPwd
func (mr *MockUserServiceMockRecorder) CheckUserPwd(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUserPwd", reflect.TypeOf((*MockUserService)(nil).CheckUserPwd), varargs...)
}

// MockUserHandler is a mock of UserHandler interface
type MockUserHandler struct {
	ctrl     *gomock.Controller
	recorder *MockUserHandlerMockRecorder
}

// MockUserHandlerMockRecorder is the mock recorder for MockUserHandler
type MockUserHandlerMockRecorder struct {
	mock *MockUserHandler
}

// NewMockUserHandler creates a new mock instance
func NewMockUserHandler(ctrl *gomock.Controller) *MockUserHandler {
	mock := &MockUserHandler{ctrl: ctrl}
	mock.recorder = &MockUserHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserHandler) EXPECT() *MockUserHandlerMockRecorder {
	return m.recorder
}

// CreateUserInfo mocks base method
func (m *MockUserHandler) CreateUserInfo(arg0 context.Context, arg1 *proto.Request, arg2 *proto.Response) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserInfo", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUserInfo indicates an expected call of CreateUserInfo
func (mr *MockUserHandlerMockRecorder) CreateUserInfo(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserInfo", reflect.TypeOf((*MockUserHandler)(nil).CreateUserInfo), arg0, arg1, arg2)
}

// GetUserInfo mocks base method
func (m *MockUserHandler) GetUserInfo(arg0 context.Context, arg1 *proto.Request, arg2 *proto.Response) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserInfo", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetUserInfo indicates an expected call of GetUserInfo
func (mr *MockUserHandlerMockRecorder) GetUserInfo(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserInfo", reflect.TypeOf((*MockUserHandler)(nil).GetUserInfo), arg0, arg1, arg2)
}

// UpdateUserInfo mocks base method
func (m *MockUserHandler) UpdateUserInfo(arg0 context.Context, arg1 *proto.Request, arg2 *proto.Response) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserInfo", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserInfo indicates an expected call of UpdateUserInfo
func (mr *MockUserHandlerMockRecorder) UpdateUserInfo(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserInfo", reflect.TypeOf((*MockUserHandler)(nil).UpdateUserInfo), arg0, arg1, arg2)
}

// CheckUserPwd mocks base method
func (m *MockUserHandler) CheckUserPwd(arg0 context.Context, arg1 *proto.Request, arg2 *proto.Response) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUserPwd", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckUserPwd indicates an expected call of CheckUserPwd
func (mr *MockUserHandlerMockRecorder) CheckUserPwd(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUserPwd", reflect.TypeOf((*MockUserHandler)(nil).CheckUserPwd), arg0, arg1, arg2)
}
