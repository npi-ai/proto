// Code generated by protoc-gen-go-grpc-mock. DO NOT EDIT.
// source: api/api.proto

package server

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockChatServerClient is a mock of ChatServerClient interface.
type MockChatServerClient struct {
	ctrl     *gomock.Controller
	recorder *MockChatServerClientMockRecorder
}

// MockChatServerClientMockRecorder is the mock recorder for MockChatServerClient.
type MockChatServerClientMockRecorder struct {
	mock *MockChatServerClient
}

// NewMockChatServerClient creates a new mock instance.
func NewMockChatServerClient(ctrl *gomock.Controller) *MockChatServerClient {
	mock := &MockChatServerClient{ctrl: ctrl}
	mock.recorder = &MockChatServerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChatServerClient) EXPECT() *MockChatServerClientMockRecorder {
	return m.recorder
}

// Chat mocks base method.
func (m *MockChatServerClient) Chat(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Chat", varargs...)
	ret0, _ := ret[0].(*Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Chat indicates an expected call of Chat.
func (mr *MockChatServerClientMockRecorder) Chat(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Chat", reflect.TypeOf((*MockChatServerClient)(nil).Chat), varargs...)
}

// MockChatServerServer is a mock of ChatServerServer interface.
type MockChatServerServer struct {
	ctrl     *gomock.Controller
	recorder *MockChatServerServerMockRecorder
}

// MockChatServerServerMockRecorder is the mock recorder for MockChatServerServer.
type MockChatServerServerMockRecorder struct {
	mock *MockChatServerServer
}

// NewMockChatServerServer creates a new mock instance.
func NewMockChatServerServer(ctrl *gomock.Controller) *MockChatServerServer {
	mock := &MockChatServerServer{ctrl: ctrl}
	mock.recorder = &MockChatServerServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChatServerServer) EXPECT() *MockChatServerServerMockRecorder {
	return m.recorder
}

// Chat mocks base method.
func (m *MockChatServerServer) Chat(ctx context.Context, in *Request) (*Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Chat", ctx, in)
	ret0, _ := ret[0].(*Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Chat indicates an expected call of Chat.
func (mr *MockChatServerServerMockRecorder) Chat(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Chat", reflect.TypeOf((*MockChatServerServer)(nil).Chat), ctx, in)
}