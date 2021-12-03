// Code generated by MockGen. DO NOT EDIT.
// Source: ./event_handlers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/solo-io/solo-apis/pkg/api/fed.solo.io/v1"
	controller "github.com/solo-io/solo-apis/pkg/api/fed.solo.io/v1/controller"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockGlooInstanceEventHandler is a mock of GlooInstanceEventHandler interface.
type MockGlooInstanceEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockGlooInstanceEventHandlerMockRecorder
}

// MockGlooInstanceEventHandlerMockRecorder is the mock recorder for MockGlooInstanceEventHandler.
type MockGlooInstanceEventHandlerMockRecorder struct {
	mock *MockGlooInstanceEventHandler
}

// NewMockGlooInstanceEventHandler creates a new mock instance.
func NewMockGlooInstanceEventHandler(ctrl *gomock.Controller) *MockGlooInstanceEventHandler {
	mock := &MockGlooInstanceEventHandler{ctrl: ctrl}
	mock.recorder = &MockGlooInstanceEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGlooInstanceEventHandler) EXPECT() *MockGlooInstanceEventHandlerMockRecorder {
	return m.recorder
}

// CreateGlooInstance mocks base method.
func (m *MockGlooInstanceEventHandler) CreateGlooInstance(obj *v1.GlooInstance) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGlooInstance", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateGlooInstance indicates an expected call of CreateGlooInstance.
func (mr *MockGlooInstanceEventHandlerMockRecorder) CreateGlooInstance(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGlooInstance", reflect.TypeOf((*MockGlooInstanceEventHandler)(nil).CreateGlooInstance), obj)
}

// DeleteGlooInstance mocks base method.
func (m *MockGlooInstanceEventHandler) DeleteGlooInstance(obj *v1.GlooInstance) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteGlooInstance", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteGlooInstance indicates an expected call of DeleteGlooInstance.
func (mr *MockGlooInstanceEventHandlerMockRecorder) DeleteGlooInstance(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGlooInstance", reflect.TypeOf((*MockGlooInstanceEventHandler)(nil).DeleteGlooInstance), obj)
}

// GenericGlooInstance mocks base method.
func (m *MockGlooInstanceEventHandler) GenericGlooInstance(obj *v1.GlooInstance) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericGlooInstance", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericGlooInstance indicates an expected call of GenericGlooInstance.
func (mr *MockGlooInstanceEventHandlerMockRecorder) GenericGlooInstance(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericGlooInstance", reflect.TypeOf((*MockGlooInstanceEventHandler)(nil).GenericGlooInstance), obj)
}

// UpdateGlooInstance mocks base method.
func (m *MockGlooInstanceEventHandler) UpdateGlooInstance(old, new *v1.GlooInstance) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateGlooInstance", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateGlooInstance indicates an expected call of UpdateGlooInstance.
func (mr *MockGlooInstanceEventHandlerMockRecorder) UpdateGlooInstance(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGlooInstance", reflect.TypeOf((*MockGlooInstanceEventHandler)(nil).UpdateGlooInstance), old, new)
}

// MockGlooInstanceEventWatcher is a mock of GlooInstanceEventWatcher interface.
type MockGlooInstanceEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockGlooInstanceEventWatcherMockRecorder
}

// MockGlooInstanceEventWatcherMockRecorder is the mock recorder for MockGlooInstanceEventWatcher.
type MockGlooInstanceEventWatcherMockRecorder struct {
	mock *MockGlooInstanceEventWatcher
}

// NewMockGlooInstanceEventWatcher creates a new mock instance.
func NewMockGlooInstanceEventWatcher(ctrl *gomock.Controller) *MockGlooInstanceEventWatcher {
	mock := &MockGlooInstanceEventWatcher{ctrl: ctrl}
	mock.recorder = &MockGlooInstanceEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGlooInstanceEventWatcher) EXPECT() *MockGlooInstanceEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
func (m *MockGlooInstanceEventWatcher) AddEventHandler(ctx context.Context, h controller.GlooInstanceEventHandler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, h}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddEventHandler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEventHandler indicates an expected call of AddEventHandler.
func (mr *MockGlooInstanceEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockGlooInstanceEventWatcher)(nil).AddEventHandler), varargs...)
}

// MockFailoverSchemeEventHandler is a mock of FailoverSchemeEventHandler interface.
type MockFailoverSchemeEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockFailoverSchemeEventHandlerMockRecorder
}

// MockFailoverSchemeEventHandlerMockRecorder is the mock recorder for MockFailoverSchemeEventHandler.
type MockFailoverSchemeEventHandlerMockRecorder struct {
	mock *MockFailoverSchemeEventHandler
}

// NewMockFailoverSchemeEventHandler creates a new mock instance.
func NewMockFailoverSchemeEventHandler(ctrl *gomock.Controller) *MockFailoverSchemeEventHandler {
	mock := &MockFailoverSchemeEventHandler{ctrl: ctrl}
	mock.recorder = &MockFailoverSchemeEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFailoverSchemeEventHandler) EXPECT() *MockFailoverSchemeEventHandlerMockRecorder {
	return m.recorder
}

// CreateFailoverScheme mocks base method.
func (m *MockFailoverSchemeEventHandler) CreateFailoverScheme(obj *v1.FailoverScheme) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFailoverScheme", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateFailoverScheme indicates an expected call of CreateFailoverScheme.
func (mr *MockFailoverSchemeEventHandlerMockRecorder) CreateFailoverScheme(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFailoverScheme", reflect.TypeOf((*MockFailoverSchemeEventHandler)(nil).CreateFailoverScheme), obj)
}

// DeleteFailoverScheme mocks base method.
func (m *MockFailoverSchemeEventHandler) DeleteFailoverScheme(obj *v1.FailoverScheme) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFailoverScheme", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFailoverScheme indicates an expected call of DeleteFailoverScheme.
func (mr *MockFailoverSchemeEventHandlerMockRecorder) DeleteFailoverScheme(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFailoverScheme", reflect.TypeOf((*MockFailoverSchemeEventHandler)(nil).DeleteFailoverScheme), obj)
}

// GenericFailoverScheme mocks base method.
func (m *MockFailoverSchemeEventHandler) GenericFailoverScheme(obj *v1.FailoverScheme) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericFailoverScheme", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericFailoverScheme indicates an expected call of GenericFailoverScheme.
func (mr *MockFailoverSchemeEventHandlerMockRecorder) GenericFailoverScheme(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericFailoverScheme", reflect.TypeOf((*MockFailoverSchemeEventHandler)(nil).GenericFailoverScheme), obj)
}

// UpdateFailoverScheme mocks base method.
func (m *MockFailoverSchemeEventHandler) UpdateFailoverScheme(old, new *v1.FailoverScheme) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFailoverScheme", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateFailoverScheme indicates an expected call of UpdateFailoverScheme.
func (mr *MockFailoverSchemeEventHandlerMockRecorder) UpdateFailoverScheme(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFailoverScheme", reflect.TypeOf((*MockFailoverSchemeEventHandler)(nil).UpdateFailoverScheme), old, new)
}

// MockFailoverSchemeEventWatcher is a mock of FailoverSchemeEventWatcher interface.
type MockFailoverSchemeEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockFailoverSchemeEventWatcherMockRecorder
}

// MockFailoverSchemeEventWatcherMockRecorder is the mock recorder for MockFailoverSchemeEventWatcher.
type MockFailoverSchemeEventWatcherMockRecorder struct {
	mock *MockFailoverSchemeEventWatcher
}

// NewMockFailoverSchemeEventWatcher creates a new mock instance.
func NewMockFailoverSchemeEventWatcher(ctrl *gomock.Controller) *MockFailoverSchemeEventWatcher {
	mock := &MockFailoverSchemeEventWatcher{ctrl: ctrl}
	mock.recorder = &MockFailoverSchemeEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFailoverSchemeEventWatcher) EXPECT() *MockFailoverSchemeEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
func (m *MockFailoverSchemeEventWatcher) AddEventHandler(ctx context.Context, h controller.FailoverSchemeEventHandler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, h}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddEventHandler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEventHandler indicates an expected call of AddEventHandler.
func (mr *MockFailoverSchemeEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockFailoverSchemeEventWatcher)(nil).AddEventHandler), varargs...)
}
