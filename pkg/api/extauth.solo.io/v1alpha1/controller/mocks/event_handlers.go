// Code generated by MockGen. DO NOT EDIT.
// Source: ./event_handlers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/solo-io/ext-auth-service/pkg/api/extauth.solo.io/v1alpha1"
	controller "github.com/solo-io/ext-auth-service/pkg/api/extauth.solo.io/v1alpha1/controller"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockExtAuthConfigEventHandler is a mock of ExtAuthConfigEventHandler interface
type MockExtAuthConfigEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockExtAuthConfigEventHandlerMockRecorder
}

// MockExtAuthConfigEventHandlerMockRecorder is the mock recorder for MockExtAuthConfigEventHandler
type MockExtAuthConfigEventHandlerMockRecorder struct {
	mock *MockExtAuthConfigEventHandler
}

// NewMockExtAuthConfigEventHandler creates a new mock instance
func NewMockExtAuthConfigEventHandler(ctrl *gomock.Controller) *MockExtAuthConfigEventHandler {
	mock := &MockExtAuthConfigEventHandler{ctrl: ctrl}
	mock.recorder = &MockExtAuthConfigEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockExtAuthConfigEventHandler) EXPECT() *MockExtAuthConfigEventHandlerMockRecorder {
	return m.recorder
}

// CreateExtAuthConfig mocks base method
func (m *MockExtAuthConfigEventHandler) CreateExtAuthConfig(obj *v1alpha1.ExtAuthConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateExtAuthConfig", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateExtAuthConfig indicates an expected call of CreateExtAuthConfig
func (mr *MockExtAuthConfigEventHandlerMockRecorder) CreateExtAuthConfig(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateExtAuthConfig", reflect.TypeOf((*MockExtAuthConfigEventHandler)(nil).CreateExtAuthConfig), obj)
}

// UpdateExtAuthConfig mocks base method
func (m *MockExtAuthConfigEventHandler) UpdateExtAuthConfig(old, new *v1alpha1.ExtAuthConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateExtAuthConfig", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateExtAuthConfig indicates an expected call of UpdateExtAuthConfig
func (mr *MockExtAuthConfigEventHandlerMockRecorder) UpdateExtAuthConfig(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateExtAuthConfig", reflect.TypeOf((*MockExtAuthConfigEventHandler)(nil).UpdateExtAuthConfig), old, new)
}

// DeleteExtAuthConfig mocks base method
func (m *MockExtAuthConfigEventHandler) DeleteExtAuthConfig(obj *v1alpha1.ExtAuthConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteExtAuthConfig", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteExtAuthConfig indicates an expected call of DeleteExtAuthConfig
func (mr *MockExtAuthConfigEventHandlerMockRecorder) DeleteExtAuthConfig(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteExtAuthConfig", reflect.TypeOf((*MockExtAuthConfigEventHandler)(nil).DeleteExtAuthConfig), obj)
}

// GenericExtAuthConfig mocks base method
func (m *MockExtAuthConfigEventHandler) GenericExtAuthConfig(obj *v1alpha1.ExtAuthConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericExtAuthConfig", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericExtAuthConfig indicates an expected call of GenericExtAuthConfig
func (mr *MockExtAuthConfigEventHandlerMockRecorder) GenericExtAuthConfig(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericExtAuthConfig", reflect.TypeOf((*MockExtAuthConfigEventHandler)(nil).GenericExtAuthConfig), obj)
}

// MockExtAuthConfigEventWatcher is a mock of ExtAuthConfigEventWatcher interface
type MockExtAuthConfigEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockExtAuthConfigEventWatcherMockRecorder
}

// MockExtAuthConfigEventWatcherMockRecorder is the mock recorder for MockExtAuthConfigEventWatcher
type MockExtAuthConfigEventWatcherMockRecorder struct {
	mock *MockExtAuthConfigEventWatcher
}

// NewMockExtAuthConfigEventWatcher creates a new mock instance
func NewMockExtAuthConfigEventWatcher(ctrl *gomock.Controller) *MockExtAuthConfigEventWatcher {
	mock := &MockExtAuthConfigEventWatcher{ctrl: ctrl}
	mock.recorder = &MockExtAuthConfigEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockExtAuthConfigEventWatcher) EXPECT() *MockExtAuthConfigEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method
func (m *MockExtAuthConfigEventWatcher) AddEventHandler(ctx context.Context, h controller.ExtAuthConfigEventHandler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, h}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddEventHandler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEventHandler indicates an expected call of AddEventHandler
func (mr *MockExtAuthConfigEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockExtAuthConfigEventWatcher)(nil).AddEventHandler), varargs...)
}
