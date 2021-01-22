// Code generated by MockGen. DO NOT EDIT.
// Source: ./domain/message/proto/message.pb.micro.go

// Package test is a generated GoMock package.
package test

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	client "github.com/micro/go-micro/client"
	proto "liaotian/domain/message/proto"
	reflect "reflect"
)

// MockMessageService is a mock of MessageService interface
type MockMessageService struct {
	ctrl     *gomock.Controller
	recorder *MockMessageServiceMockRecorder
}

// MockMessageServiceMockRecorder is the mock recorder for MockMessageService
type MockMessageServiceMockRecorder struct {
	mock *MockMessageService
}

// NewMockMessageService creates a new mock instance
func NewMockMessageService(ctrl *gomock.Controller) *MockMessageService {
	mock := &MockMessageService{ctrl: ctrl}
	mock.recorder = &MockMessageServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMessageService) EXPECT() *MockMessageServiceMockRecorder {
	return m.recorder
}

// Sub mocks base method
func (m *MockMessageService) Sub(ctx context.Context, in *proto.SubRequest, opts ...client.CallOption) (*proto.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Sub", varargs...)
	ret0, _ := ret[0].(*proto.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sub indicates an expected call of Sub
func (mr *MockMessageServiceMockRecorder) Sub(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sub", reflect.TypeOf((*MockMessageService)(nil).Sub), varargs...)
}

// UnSub mocks base method
func (m *MockMessageService) UnSub(ctx context.Context, in *proto.UnSubRequest, opts ...client.CallOption) (*proto.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UnSub", varargs...)
	ret0, _ := ret[0].(*proto.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnSub indicates an expected call of UnSub
func (mr *MockMessageServiceMockRecorder) UnSub(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnSub", reflect.TypeOf((*MockMessageService)(nil).UnSub), varargs...)
}

// Send mocks base method
func (m *MockMessageService) Send(ctx context.Context, in *proto.SendRequest, opts ...client.CallOption) (*proto.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Send", varargs...)
	ret0, _ := ret[0].(*proto.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Send indicates an expected call of Send
func (mr *MockMessageServiceMockRecorder) Send(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockMessageService)(nil).Send), varargs...)
}

// MockMessageHandler is a mock of MessageHandler interface
type MockMessageHandler struct {
	ctrl     *gomock.Controller
	recorder *MockMessageHandlerMockRecorder
}

// MockMessageHandlerMockRecorder is the mock recorder for MockMessageHandler
type MockMessageHandlerMockRecorder struct {
	mock *MockMessageHandler
}

// NewMockMessageHandler creates a new mock instance
func NewMockMessageHandler(ctrl *gomock.Controller) *MockMessageHandler {
	mock := &MockMessageHandler{ctrl: ctrl}
	mock.recorder = &MockMessageHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMessageHandler) EXPECT() *MockMessageHandlerMockRecorder {
	return m.recorder
}

// Sub mocks base method
func (m *MockMessageHandler) Sub(arg0 context.Context, arg1 *proto.SubRequest, arg2 *proto.Response) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sub", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Sub indicates an expected call of Sub
func (mr *MockMessageHandlerMockRecorder) Sub(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sub", reflect.TypeOf((*MockMessageHandler)(nil).Sub), arg0, arg1, arg2)
}

// UnSub mocks base method
func (m *MockMessageHandler) UnSub(arg0 context.Context, arg1 *proto.UnSubRequest, arg2 *proto.Response) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnSub", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnSub indicates an expected call of UnSub
func (mr *MockMessageHandlerMockRecorder) UnSub(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnSub", reflect.TypeOf((*MockMessageHandler)(nil).UnSub), arg0, arg1, arg2)
}

// Send mocks base method
func (m *MockMessageHandler) Send(arg0 context.Context, arg1 *proto.SendRequest, arg2 *proto.Response) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockMessageHandlerMockRecorder) Send(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockMessageHandler)(nil).Send), arg0, arg1, arg2)
}