// Code generated by MockGen. DO NOT EDIT.
// Source: ./reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	v1 "github.com/solo-io/solo-apis/pkg/api/fed.gateway.solo.io/v1"
	controller "github.com/solo-io/solo-apis/pkg/api/fed.gateway.solo.io/v1/controller"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockFederatedGatewayReconciler is a mock of FederatedGatewayReconciler interface
type MockFederatedGatewayReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedGatewayReconcilerMockRecorder
}

// MockFederatedGatewayReconcilerMockRecorder is the mock recorder for MockFederatedGatewayReconciler
type MockFederatedGatewayReconcilerMockRecorder struct {
	mock *MockFederatedGatewayReconciler
}

// NewMockFederatedGatewayReconciler creates a new mock instance
func NewMockFederatedGatewayReconciler(ctrl *gomock.Controller) *MockFederatedGatewayReconciler {
	mock := &MockFederatedGatewayReconciler{ctrl: ctrl}
	mock.recorder = &MockFederatedGatewayReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFederatedGatewayReconciler) EXPECT() *MockFederatedGatewayReconcilerMockRecorder {
	return m.recorder
}

// ReconcileFederatedGateway mocks base method
func (m *MockFederatedGatewayReconciler) ReconcileFederatedGateway(obj *v1.FederatedGateway) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileFederatedGateway", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileFederatedGateway indicates an expected call of ReconcileFederatedGateway
func (mr *MockFederatedGatewayReconcilerMockRecorder) ReconcileFederatedGateway(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileFederatedGateway", reflect.TypeOf((*MockFederatedGatewayReconciler)(nil).ReconcileFederatedGateway), obj)
}

// MockFederatedGatewayDeletionReconciler is a mock of FederatedGatewayDeletionReconciler interface
type MockFederatedGatewayDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedGatewayDeletionReconcilerMockRecorder
}

// MockFederatedGatewayDeletionReconcilerMockRecorder is the mock recorder for MockFederatedGatewayDeletionReconciler
type MockFederatedGatewayDeletionReconcilerMockRecorder struct {
	mock *MockFederatedGatewayDeletionReconciler
}

// NewMockFederatedGatewayDeletionReconciler creates a new mock instance
func NewMockFederatedGatewayDeletionReconciler(ctrl *gomock.Controller) *MockFederatedGatewayDeletionReconciler {
	mock := &MockFederatedGatewayDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockFederatedGatewayDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFederatedGatewayDeletionReconciler) EXPECT() *MockFederatedGatewayDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileFederatedGatewayDeletion mocks base method
func (m *MockFederatedGatewayDeletionReconciler) ReconcileFederatedGatewayDeletion(req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileFederatedGatewayDeletion", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileFederatedGatewayDeletion indicates an expected call of ReconcileFederatedGatewayDeletion
func (mr *MockFederatedGatewayDeletionReconcilerMockRecorder) ReconcileFederatedGatewayDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileFederatedGatewayDeletion", reflect.TypeOf((*MockFederatedGatewayDeletionReconciler)(nil).ReconcileFederatedGatewayDeletion), req)
}

// MockFederatedGatewayFinalizer is a mock of FederatedGatewayFinalizer interface
type MockFederatedGatewayFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedGatewayFinalizerMockRecorder
}

// MockFederatedGatewayFinalizerMockRecorder is the mock recorder for MockFederatedGatewayFinalizer
type MockFederatedGatewayFinalizerMockRecorder struct {
	mock *MockFederatedGatewayFinalizer
}

// NewMockFederatedGatewayFinalizer creates a new mock instance
func NewMockFederatedGatewayFinalizer(ctrl *gomock.Controller) *MockFederatedGatewayFinalizer {
	mock := &MockFederatedGatewayFinalizer{ctrl: ctrl}
	mock.recorder = &MockFederatedGatewayFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFederatedGatewayFinalizer) EXPECT() *MockFederatedGatewayFinalizerMockRecorder {
	return m.recorder
}

// ReconcileFederatedGateway mocks base method
func (m *MockFederatedGatewayFinalizer) ReconcileFederatedGateway(obj *v1.FederatedGateway) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileFederatedGateway", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileFederatedGateway indicates an expected call of ReconcileFederatedGateway
func (mr *MockFederatedGatewayFinalizerMockRecorder) ReconcileFederatedGateway(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileFederatedGateway", reflect.TypeOf((*MockFederatedGatewayFinalizer)(nil).ReconcileFederatedGateway), obj)
}

// FederatedGatewayFinalizerName mocks base method
func (m *MockFederatedGatewayFinalizer) FederatedGatewayFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FederatedGatewayFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// FederatedGatewayFinalizerName indicates an expected call of FederatedGatewayFinalizerName
func (mr *MockFederatedGatewayFinalizerMockRecorder) FederatedGatewayFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FederatedGatewayFinalizerName", reflect.TypeOf((*MockFederatedGatewayFinalizer)(nil).FederatedGatewayFinalizerName))
}

// FinalizeFederatedGateway mocks base method
func (m *MockFederatedGatewayFinalizer) FinalizeFederatedGateway(obj *v1.FederatedGateway) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeFederatedGateway", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeFederatedGateway indicates an expected call of FinalizeFederatedGateway
func (mr *MockFederatedGatewayFinalizerMockRecorder) FinalizeFederatedGateway(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeFederatedGateway", reflect.TypeOf((*MockFederatedGatewayFinalizer)(nil).FinalizeFederatedGateway), obj)
}

// MockFederatedGatewayReconcileLoop is a mock of FederatedGatewayReconcileLoop interface
type MockFederatedGatewayReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedGatewayReconcileLoopMockRecorder
}

// MockFederatedGatewayReconcileLoopMockRecorder is the mock recorder for MockFederatedGatewayReconcileLoop
type MockFederatedGatewayReconcileLoopMockRecorder struct {
	mock *MockFederatedGatewayReconcileLoop
}

// NewMockFederatedGatewayReconcileLoop creates a new mock instance
func NewMockFederatedGatewayReconcileLoop(ctrl *gomock.Controller) *MockFederatedGatewayReconcileLoop {
	mock := &MockFederatedGatewayReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockFederatedGatewayReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFederatedGatewayReconcileLoop) EXPECT() *MockFederatedGatewayReconcileLoopMockRecorder {
	return m.recorder
}

// RunFederatedGatewayReconciler mocks base method
func (m *MockFederatedGatewayReconcileLoop) RunFederatedGatewayReconciler(ctx context.Context, rec controller.FederatedGatewayReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunFederatedGatewayReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunFederatedGatewayReconciler indicates an expected call of RunFederatedGatewayReconciler
func (mr *MockFederatedGatewayReconcileLoopMockRecorder) RunFederatedGatewayReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunFederatedGatewayReconciler", reflect.TypeOf((*MockFederatedGatewayReconcileLoop)(nil).RunFederatedGatewayReconciler), varargs...)
}

// MockFederatedRouteTableReconciler is a mock of FederatedRouteTableReconciler interface
type MockFederatedRouteTableReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedRouteTableReconcilerMockRecorder
}

// MockFederatedRouteTableReconcilerMockRecorder is the mock recorder for MockFederatedRouteTableReconciler
type MockFederatedRouteTableReconcilerMockRecorder struct {
	mock *MockFederatedRouteTableReconciler
}

// NewMockFederatedRouteTableReconciler creates a new mock instance
func NewMockFederatedRouteTableReconciler(ctrl *gomock.Controller) *MockFederatedRouteTableReconciler {
	mock := &MockFederatedRouteTableReconciler{ctrl: ctrl}
	mock.recorder = &MockFederatedRouteTableReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFederatedRouteTableReconciler) EXPECT() *MockFederatedRouteTableReconcilerMockRecorder {
	return m.recorder
}

// ReconcileFederatedRouteTable mocks base method
func (m *MockFederatedRouteTableReconciler) ReconcileFederatedRouteTable(obj *v1.FederatedRouteTable) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileFederatedRouteTable", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileFederatedRouteTable indicates an expected call of ReconcileFederatedRouteTable
func (mr *MockFederatedRouteTableReconcilerMockRecorder) ReconcileFederatedRouteTable(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileFederatedRouteTable", reflect.TypeOf((*MockFederatedRouteTableReconciler)(nil).ReconcileFederatedRouteTable), obj)
}

// MockFederatedRouteTableDeletionReconciler is a mock of FederatedRouteTableDeletionReconciler interface
type MockFederatedRouteTableDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedRouteTableDeletionReconcilerMockRecorder
}

// MockFederatedRouteTableDeletionReconcilerMockRecorder is the mock recorder for MockFederatedRouteTableDeletionReconciler
type MockFederatedRouteTableDeletionReconcilerMockRecorder struct {
	mock *MockFederatedRouteTableDeletionReconciler
}

// NewMockFederatedRouteTableDeletionReconciler creates a new mock instance
func NewMockFederatedRouteTableDeletionReconciler(ctrl *gomock.Controller) *MockFederatedRouteTableDeletionReconciler {
	mock := &MockFederatedRouteTableDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockFederatedRouteTableDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFederatedRouteTableDeletionReconciler) EXPECT() *MockFederatedRouteTableDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileFederatedRouteTableDeletion mocks base method
func (m *MockFederatedRouteTableDeletionReconciler) ReconcileFederatedRouteTableDeletion(req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileFederatedRouteTableDeletion", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileFederatedRouteTableDeletion indicates an expected call of ReconcileFederatedRouteTableDeletion
func (mr *MockFederatedRouteTableDeletionReconcilerMockRecorder) ReconcileFederatedRouteTableDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileFederatedRouteTableDeletion", reflect.TypeOf((*MockFederatedRouteTableDeletionReconciler)(nil).ReconcileFederatedRouteTableDeletion), req)
}

// MockFederatedRouteTableFinalizer is a mock of FederatedRouteTableFinalizer interface
type MockFederatedRouteTableFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedRouteTableFinalizerMockRecorder
}

// MockFederatedRouteTableFinalizerMockRecorder is the mock recorder for MockFederatedRouteTableFinalizer
type MockFederatedRouteTableFinalizerMockRecorder struct {
	mock *MockFederatedRouteTableFinalizer
}

// NewMockFederatedRouteTableFinalizer creates a new mock instance
func NewMockFederatedRouteTableFinalizer(ctrl *gomock.Controller) *MockFederatedRouteTableFinalizer {
	mock := &MockFederatedRouteTableFinalizer{ctrl: ctrl}
	mock.recorder = &MockFederatedRouteTableFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFederatedRouteTableFinalizer) EXPECT() *MockFederatedRouteTableFinalizerMockRecorder {
	return m.recorder
}

// ReconcileFederatedRouteTable mocks base method
func (m *MockFederatedRouteTableFinalizer) ReconcileFederatedRouteTable(obj *v1.FederatedRouteTable) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileFederatedRouteTable", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileFederatedRouteTable indicates an expected call of ReconcileFederatedRouteTable
func (mr *MockFederatedRouteTableFinalizerMockRecorder) ReconcileFederatedRouteTable(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileFederatedRouteTable", reflect.TypeOf((*MockFederatedRouteTableFinalizer)(nil).ReconcileFederatedRouteTable), obj)
}

// FederatedRouteTableFinalizerName mocks base method
func (m *MockFederatedRouteTableFinalizer) FederatedRouteTableFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FederatedRouteTableFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// FederatedRouteTableFinalizerName indicates an expected call of FederatedRouteTableFinalizerName
func (mr *MockFederatedRouteTableFinalizerMockRecorder) FederatedRouteTableFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FederatedRouteTableFinalizerName", reflect.TypeOf((*MockFederatedRouteTableFinalizer)(nil).FederatedRouteTableFinalizerName))
}

// FinalizeFederatedRouteTable mocks base method
func (m *MockFederatedRouteTableFinalizer) FinalizeFederatedRouteTable(obj *v1.FederatedRouteTable) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeFederatedRouteTable", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeFederatedRouteTable indicates an expected call of FinalizeFederatedRouteTable
func (mr *MockFederatedRouteTableFinalizerMockRecorder) FinalizeFederatedRouteTable(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeFederatedRouteTable", reflect.TypeOf((*MockFederatedRouteTableFinalizer)(nil).FinalizeFederatedRouteTable), obj)
}

// MockFederatedRouteTableReconcileLoop is a mock of FederatedRouteTableReconcileLoop interface
type MockFederatedRouteTableReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedRouteTableReconcileLoopMockRecorder
}

// MockFederatedRouteTableReconcileLoopMockRecorder is the mock recorder for MockFederatedRouteTableReconcileLoop
type MockFederatedRouteTableReconcileLoopMockRecorder struct {
	mock *MockFederatedRouteTableReconcileLoop
}

// NewMockFederatedRouteTableReconcileLoop creates a new mock instance
func NewMockFederatedRouteTableReconcileLoop(ctrl *gomock.Controller) *MockFederatedRouteTableReconcileLoop {
	mock := &MockFederatedRouteTableReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockFederatedRouteTableReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFederatedRouteTableReconcileLoop) EXPECT() *MockFederatedRouteTableReconcileLoopMockRecorder {
	return m.recorder
}

// RunFederatedRouteTableReconciler mocks base method
func (m *MockFederatedRouteTableReconcileLoop) RunFederatedRouteTableReconciler(ctx context.Context, rec controller.FederatedRouteTableReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunFederatedRouteTableReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunFederatedRouteTableReconciler indicates an expected call of RunFederatedRouteTableReconciler
func (mr *MockFederatedRouteTableReconcileLoopMockRecorder) RunFederatedRouteTableReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunFederatedRouteTableReconciler", reflect.TypeOf((*MockFederatedRouteTableReconcileLoop)(nil).RunFederatedRouteTableReconciler), varargs...)
}

// MockFederatedVirtualServiceReconciler is a mock of FederatedVirtualServiceReconciler interface
type MockFederatedVirtualServiceReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedVirtualServiceReconcilerMockRecorder
}

// MockFederatedVirtualServiceReconcilerMockRecorder is the mock recorder for MockFederatedVirtualServiceReconciler
type MockFederatedVirtualServiceReconcilerMockRecorder struct {
	mock *MockFederatedVirtualServiceReconciler
}

// NewMockFederatedVirtualServiceReconciler creates a new mock instance
func NewMockFederatedVirtualServiceReconciler(ctrl *gomock.Controller) *MockFederatedVirtualServiceReconciler {
	mock := &MockFederatedVirtualServiceReconciler{ctrl: ctrl}
	mock.recorder = &MockFederatedVirtualServiceReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFederatedVirtualServiceReconciler) EXPECT() *MockFederatedVirtualServiceReconcilerMockRecorder {
	return m.recorder
}

// ReconcileFederatedVirtualService mocks base method
func (m *MockFederatedVirtualServiceReconciler) ReconcileFederatedVirtualService(obj *v1.FederatedVirtualService) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileFederatedVirtualService", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileFederatedVirtualService indicates an expected call of ReconcileFederatedVirtualService
func (mr *MockFederatedVirtualServiceReconcilerMockRecorder) ReconcileFederatedVirtualService(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileFederatedVirtualService", reflect.TypeOf((*MockFederatedVirtualServiceReconciler)(nil).ReconcileFederatedVirtualService), obj)
}

// MockFederatedVirtualServiceDeletionReconciler is a mock of FederatedVirtualServiceDeletionReconciler interface
type MockFederatedVirtualServiceDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedVirtualServiceDeletionReconcilerMockRecorder
}

// MockFederatedVirtualServiceDeletionReconcilerMockRecorder is the mock recorder for MockFederatedVirtualServiceDeletionReconciler
type MockFederatedVirtualServiceDeletionReconcilerMockRecorder struct {
	mock *MockFederatedVirtualServiceDeletionReconciler
}

// NewMockFederatedVirtualServiceDeletionReconciler creates a new mock instance
func NewMockFederatedVirtualServiceDeletionReconciler(ctrl *gomock.Controller) *MockFederatedVirtualServiceDeletionReconciler {
	mock := &MockFederatedVirtualServiceDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockFederatedVirtualServiceDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFederatedVirtualServiceDeletionReconciler) EXPECT() *MockFederatedVirtualServiceDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileFederatedVirtualServiceDeletion mocks base method
func (m *MockFederatedVirtualServiceDeletionReconciler) ReconcileFederatedVirtualServiceDeletion(req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileFederatedVirtualServiceDeletion", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileFederatedVirtualServiceDeletion indicates an expected call of ReconcileFederatedVirtualServiceDeletion
func (mr *MockFederatedVirtualServiceDeletionReconcilerMockRecorder) ReconcileFederatedVirtualServiceDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileFederatedVirtualServiceDeletion", reflect.TypeOf((*MockFederatedVirtualServiceDeletionReconciler)(nil).ReconcileFederatedVirtualServiceDeletion), req)
}

// MockFederatedVirtualServiceFinalizer is a mock of FederatedVirtualServiceFinalizer interface
type MockFederatedVirtualServiceFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedVirtualServiceFinalizerMockRecorder
}

// MockFederatedVirtualServiceFinalizerMockRecorder is the mock recorder for MockFederatedVirtualServiceFinalizer
type MockFederatedVirtualServiceFinalizerMockRecorder struct {
	mock *MockFederatedVirtualServiceFinalizer
}

// NewMockFederatedVirtualServiceFinalizer creates a new mock instance
func NewMockFederatedVirtualServiceFinalizer(ctrl *gomock.Controller) *MockFederatedVirtualServiceFinalizer {
	mock := &MockFederatedVirtualServiceFinalizer{ctrl: ctrl}
	mock.recorder = &MockFederatedVirtualServiceFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFederatedVirtualServiceFinalizer) EXPECT() *MockFederatedVirtualServiceFinalizerMockRecorder {
	return m.recorder
}

// ReconcileFederatedVirtualService mocks base method
func (m *MockFederatedVirtualServiceFinalizer) ReconcileFederatedVirtualService(obj *v1.FederatedVirtualService) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileFederatedVirtualService", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileFederatedVirtualService indicates an expected call of ReconcileFederatedVirtualService
func (mr *MockFederatedVirtualServiceFinalizerMockRecorder) ReconcileFederatedVirtualService(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileFederatedVirtualService", reflect.TypeOf((*MockFederatedVirtualServiceFinalizer)(nil).ReconcileFederatedVirtualService), obj)
}

// FederatedVirtualServiceFinalizerName mocks base method
func (m *MockFederatedVirtualServiceFinalizer) FederatedVirtualServiceFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FederatedVirtualServiceFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// FederatedVirtualServiceFinalizerName indicates an expected call of FederatedVirtualServiceFinalizerName
func (mr *MockFederatedVirtualServiceFinalizerMockRecorder) FederatedVirtualServiceFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FederatedVirtualServiceFinalizerName", reflect.TypeOf((*MockFederatedVirtualServiceFinalizer)(nil).FederatedVirtualServiceFinalizerName))
}

// FinalizeFederatedVirtualService mocks base method
func (m *MockFederatedVirtualServiceFinalizer) FinalizeFederatedVirtualService(obj *v1.FederatedVirtualService) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeFederatedVirtualService", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeFederatedVirtualService indicates an expected call of FinalizeFederatedVirtualService
func (mr *MockFederatedVirtualServiceFinalizerMockRecorder) FinalizeFederatedVirtualService(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeFederatedVirtualService", reflect.TypeOf((*MockFederatedVirtualServiceFinalizer)(nil).FinalizeFederatedVirtualService), obj)
}

// MockFederatedVirtualServiceReconcileLoop is a mock of FederatedVirtualServiceReconcileLoop interface
type MockFederatedVirtualServiceReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedVirtualServiceReconcileLoopMockRecorder
}

// MockFederatedVirtualServiceReconcileLoopMockRecorder is the mock recorder for MockFederatedVirtualServiceReconcileLoop
type MockFederatedVirtualServiceReconcileLoopMockRecorder struct {
	mock *MockFederatedVirtualServiceReconcileLoop
}

// NewMockFederatedVirtualServiceReconcileLoop creates a new mock instance
func NewMockFederatedVirtualServiceReconcileLoop(ctrl *gomock.Controller) *MockFederatedVirtualServiceReconcileLoop {
	mock := &MockFederatedVirtualServiceReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockFederatedVirtualServiceReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFederatedVirtualServiceReconcileLoop) EXPECT() *MockFederatedVirtualServiceReconcileLoopMockRecorder {
	return m.recorder
}

// RunFederatedVirtualServiceReconciler mocks base method
func (m *MockFederatedVirtualServiceReconcileLoop) RunFederatedVirtualServiceReconciler(ctx context.Context, rec controller.FederatedVirtualServiceReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunFederatedVirtualServiceReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunFederatedVirtualServiceReconciler indicates an expected call of RunFederatedVirtualServiceReconciler
func (mr *MockFederatedVirtualServiceReconcileLoopMockRecorder) RunFederatedVirtualServiceReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunFederatedVirtualServiceReconciler", reflect.TypeOf((*MockFederatedVirtualServiceReconcileLoop)(nil).RunFederatedVirtualServiceReconciler), varargs...)
}
