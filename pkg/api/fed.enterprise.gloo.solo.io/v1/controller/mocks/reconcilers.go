// Code generated by MockGen. DO NOT EDIT.
// Source: ./reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	v1 "github.com/solo-io/solo-apis/pkg/api/fed.enterprise.gloo.solo.io/v1"
	controller "github.com/solo-io/solo-apis/pkg/api/fed.enterprise.gloo.solo.io/v1/controller"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockFederatedAuthConfigReconciler is a mock of FederatedAuthConfigReconciler interface
type MockFederatedAuthConfigReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedAuthConfigReconcilerMockRecorder
}

// MockFederatedAuthConfigReconcilerMockRecorder is the mock recorder for MockFederatedAuthConfigReconciler
type MockFederatedAuthConfigReconcilerMockRecorder struct {
	mock *MockFederatedAuthConfigReconciler
}

// NewMockFederatedAuthConfigReconciler creates a new mock instance
func NewMockFederatedAuthConfigReconciler(ctrl *gomock.Controller) *MockFederatedAuthConfigReconciler {
	mock := &MockFederatedAuthConfigReconciler{ctrl: ctrl}
	mock.recorder = &MockFederatedAuthConfigReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFederatedAuthConfigReconciler) EXPECT() *MockFederatedAuthConfigReconcilerMockRecorder {
	return m.recorder
}

// ReconcileFederatedAuthConfig mocks base method
func (m *MockFederatedAuthConfigReconciler) ReconcileFederatedAuthConfig(obj *v1.FederatedAuthConfig) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileFederatedAuthConfig", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileFederatedAuthConfig indicates an expected call of ReconcileFederatedAuthConfig
func (mr *MockFederatedAuthConfigReconcilerMockRecorder) ReconcileFederatedAuthConfig(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileFederatedAuthConfig", reflect.TypeOf((*MockFederatedAuthConfigReconciler)(nil).ReconcileFederatedAuthConfig), obj)
}

// MockFederatedAuthConfigDeletionReconciler is a mock of FederatedAuthConfigDeletionReconciler interface
type MockFederatedAuthConfigDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedAuthConfigDeletionReconcilerMockRecorder
}

// MockFederatedAuthConfigDeletionReconcilerMockRecorder is the mock recorder for MockFederatedAuthConfigDeletionReconciler
type MockFederatedAuthConfigDeletionReconcilerMockRecorder struct {
	mock *MockFederatedAuthConfigDeletionReconciler
}

// NewMockFederatedAuthConfigDeletionReconciler creates a new mock instance
func NewMockFederatedAuthConfigDeletionReconciler(ctrl *gomock.Controller) *MockFederatedAuthConfigDeletionReconciler {
	mock := &MockFederatedAuthConfigDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockFederatedAuthConfigDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFederatedAuthConfigDeletionReconciler) EXPECT() *MockFederatedAuthConfigDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileFederatedAuthConfigDeletion mocks base method
func (m *MockFederatedAuthConfigDeletionReconciler) ReconcileFederatedAuthConfigDeletion(req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileFederatedAuthConfigDeletion", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileFederatedAuthConfigDeletion indicates an expected call of ReconcileFederatedAuthConfigDeletion
func (mr *MockFederatedAuthConfigDeletionReconcilerMockRecorder) ReconcileFederatedAuthConfigDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileFederatedAuthConfigDeletion", reflect.TypeOf((*MockFederatedAuthConfigDeletionReconciler)(nil).ReconcileFederatedAuthConfigDeletion), req)
}

// MockFederatedAuthConfigFinalizer is a mock of FederatedAuthConfigFinalizer interface
type MockFederatedAuthConfigFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedAuthConfigFinalizerMockRecorder
}

// MockFederatedAuthConfigFinalizerMockRecorder is the mock recorder for MockFederatedAuthConfigFinalizer
type MockFederatedAuthConfigFinalizerMockRecorder struct {
	mock *MockFederatedAuthConfigFinalizer
}

// NewMockFederatedAuthConfigFinalizer creates a new mock instance
func NewMockFederatedAuthConfigFinalizer(ctrl *gomock.Controller) *MockFederatedAuthConfigFinalizer {
	mock := &MockFederatedAuthConfigFinalizer{ctrl: ctrl}
	mock.recorder = &MockFederatedAuthConfigFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFederatedAuthConfigFinalizer) EXPECT() *MockFederatedAuthConfigFinalizerMockRecorder {
	return m.recorder
}

// ReconcileFederatedAuthConfig mocks base method
func (m *MockFederatedAuthConfigFinalizer) ReconcileFederatedAuthConfig(obj *v1.FederatedAuthConfig) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileFederatedAuthConfig", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileFederatedAuthConfig indicates an expected call of ReconcileFederatedAuthConfig
func (mr *MockFederatedAuthConfigFinalizerMockRecorder) ReconcileFederatedAuthConfig(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileFederatedAuthConfig", reflect.TypeOf((*MockFederatedAuthConfigFinalizer)(nil).ReconcileFederatedAuthConfig), obj)
}

// FederatedAuthConfigFinalizerName mocks base method
func (m *MockFederatedAuthConfigFinalizer) FederatedAuthConfigFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FederatedAuthConfigFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// FederatedAuthConfigFinalizerName indicates an expected call of FederatedAuthConfigFinalizerName
func (mr *MockFederatedAuthConfigFinalizerMockRecorder) FederatedAuthConfigFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FederatedAuthConfigFinalizerName", reflect.TypeOf((*MockFederatedAuthConfigFinalizer)(nil).FederatedAuthConfigFinalizerName))
}

// FinalizeFederatedAuthConfig mocks base method
func (m *MockFederatedAuthConfigFinalizer) FinalizeFederatedAuthConfig(obj *v1.FederatedAuthConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeFederatedAuthConfig", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeFederatedAuthConfig indicates an expected call of FinalizeFederatedAuthConfig
func (mr *MockFederatedAuthConfigFinalizerMockRecorder) FinalizeFederatedAuthConfig(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeFederatedAuthConfig", reflect.TypeOf((*MockFederatedAuthConfigFinalizer)(nil).FinalizeFederatedAuthConfig), obj)
}

// MockFederatedAuthConfigReconcileLoop is a mock of FederatedAuthConfigReconcileLoop interface
type MockFederatedAuthConfigReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedAuthConfigReconcileLoopMockRecorder
}

// MockFederatedAuthConfigReconcileLoopMockRecorder is the mock recorder for MockFederatedAuthConfigReconcileLoop
type MockFederatedAuthConfigReconcileLoopMockRecorder struct {
	mock *MockFederatedAuthConfigReconcileLoop
}

// NewMockFederatedAuthConfigReconcileLoop creates a new mock instance
func NewMockFederatedAuthConfigReconcileLoop(ctrl *gomock.Controller) *MockFederatedAuthConfigReconcileLoop {
	mock := &MockFederatedAuthConfigReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockFederatedAuthConfigReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFederatedAuthConfigReconcileLoop) EXPECT() *MockFederatedAuthConfigReconcileLoopMockRecorder {
	return m.recorder
}

// RunFederatedAuthConfigReconciler mocks base method
func (m *MockFederatedAuthConfigReconcileLoop) RunFederatedAuthConfigReconciler(ctx context.Context, rec controller.FederatedAuthConfigReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunFederatedAuthConfigReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunFederatedAuthConfigReconciler indicates an expected call of RunFederatedAuthConfigReconciler
func (mr *MockFederatedAuthConfigReconcileLoopMockRecorder) RunFederatedAuthConfigReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunFederatedAuthConfigReconciler", reflect.TypeOf((*MockFederatedAuthConfigReconcileLoop)(nil).RunFederatedAuthConfigReconciler), varargs...)
}
