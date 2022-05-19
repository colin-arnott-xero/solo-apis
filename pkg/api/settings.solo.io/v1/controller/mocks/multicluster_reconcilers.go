// Code generated by MockGen. DO NOT EDIT.
// Source: ./multicluster_reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	v1 "github.com/solo-io/solo-apis/pkg/api/settings.solo.io/v1"
	controller "github.com/solo-io/solo-apis/pkg/api/settings.solo.io/v1/controller"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockMulticlusterSettingsReconciler is a mock of MulticlusterSettingsReconciler interface
type MockMulticlusterSettingsReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterSettingsReconcilerMockRecorder
}

// MockMulticlusterSettingsReconcilerMockRecorder is the mock recorder for MockMulticlusterSettingsReconciler
type MockMulticlusterSettingsReconcilerMockRecorder struct {
	mock *MockMulticlusterSettingsReconciler
}

// NewMockMulticlusterSettingsReconciler creates a new mock instance
func NewMockMulticlusterSettingsReconciler(ctrl *gomock.Controller) *MockMulticlusterSettingsReconciler {
	mock := &MockMulticlusterSettingsReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterSettingsReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterSettingsReconciler) EXPECT() *MockMulticlusterSettingsReconcilerMockRecorder {
	return m.recorder
}

// ReconcileSettings mocks base method
func (m *MockMulticlusterSettingsReconciler) ReconcileSettings(clusterName string, obj *v1.Settings) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileSettings", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileSettings indicates an expected call of ReconcileSettings
func (mr *MockMulticlusterSettingsReconcilerMockRecorder) ReconcileSettings(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileSettings", reflect.TypeOf((*MockMulticlusterSettingsReconciler)(nil).ReconcileSettings), clusterName, obj)
}

// MockMulticlusterSettingsDeletionReconciler is a mock of MulticlusterSettingsDeletionReconciler interface
type MockMulticlusterSettingsDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterSettingsDeletionReconcilerMockRecorder
}

// MockMulticlusterSettingsDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterSettingsDeletionReconciler
type MockMulticlusterSettingsDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterSettingsDeletionReconciler
}

// NewMockMulticlusterSettingsDeletionReconciler creates a new mock instance
func NewMockMulticlusterSettingsDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterSettingsDeletionReconciler {
	mock := &MockMulticlusterSettingsDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterSettingsDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterSettingsDeletionReconciler) EXPECT() *MockMulticlusterSettingsDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileSettingsDeletion mocks base method
func (m *MockMulticlusterSettingsDeletionReconciler) ReconcileSettingsDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileSettingsDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileSettingsDeletion indicates an expected call of ReconcileSettingsDeletion
func (mr *MockMulticlusterSettingsDeletionReconcilerMockRecorder) ReconcileSettingsDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileSettingsDeletion", reflect.TypeOf((*MockMulticlusterSettingsDeletionReconciler)(nil).ReconcileSettingsDeletion), clusterName, req)
}

// MockMulticlusterSettingsReconcileLoop is a mock of MulticlusterSettingsReconcileLoop interface
type MockMulticlusterSettingsReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterSettingsReconcileLoopMockRecorder
}

// MockMulticlusterSettingsReconcileLoopMockRecorder is the mock recorder for MockMulticlusterSettingsReconcileLoop
type MockMulticlusterSettingsReconcileLoopMockRecorder struct {
	mock *MockMulticlusterSettingsReconcileLoop
}

// NewMockMulticlusterSettingsReconcileLoop creates a new mock instance
func NewMockMulticlusterSettingsReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterSettingsReconcileLoop {
	mock := &MockMulticlusterSettingsReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterSettingsReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterSettingsReconcileLoop) EXPECT() *MockMulticlusterSettingsReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterSettingsReconciler mocks base method
func (m *MockMulticlusterSettingsReconcileLoop) AddMulticlusterSettingsReconciler(ctx context.Context, rec controller.MulticlusterSettingsReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterSettingsReconciler", varargs...)
}

// AddMulticlusterSettingsReconciler indicates an expected call of AddMulticlusterSettingsReconciler
func (mr *MockMulticlusterSettingsReconcileLoopMockRecorder) AddMulticlusterSettingsReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterSettingsReconciler", reflect.TypeOf((*MockMulticlusterSettingsReconcileLoop)(nil).AddMulticlusterSettingsReconciler), varargs...)
}

// MockMulticlusterDashboardReconciler is a mock of MulticlusterDashboardReconciler interface
type MockMulticlusterDashboardReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterDashboardReconcilerMockRecorder
}

// MockMulticlusterDashboardReconcilerMockRecorder is the mock recorder for MockMulticlusterDashboardReconciler
type MockMulticlusterDashboardReconcilerMockRecorder struct {
	mock *MockMulticlusterDashboardReconciler
}

// NewMockMulticlusterDashboardReconciler creates a new mock instance
func NewMockMulticlusterDashboardReconciler(ctrl *gomock.Controller) *MockMulticlusterDashboardReconciler {
	mock := &MockMulticlusterDashboardReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterDashboardReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterDashboardReconciler) EXPECT() *MockMulticlusterDashboardReconcilerMockRecorder {
	return m.recorder
}

// ReconcileDashboard mocks base method
func (m *MockMulticlusterDashboardReconciler) ReconcileDashboard(clusterName string, obj *v1.Dashboard) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileDashboard", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileDashboard indicates an expected call of ReconcileDashboard
func (mr *MockMulticlusterDashboardReconcilerMockRecorder) ReconcileDashboard(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileDashboard", reflect.TypeOf((*MockMulticlusterDashboardReconciler)(nil).ReconcileDashboard), clusterName, obj)
}

// MockMulticlusterDashboardDeletionReconciler is a mock of MulticlusterDashboardDeletionReconciler interface
type MockMulticlusterDashboardDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterDashboardDeletionReconcilerMockRecorder
}

// MockMulticlusterDashboardDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterDashboardDeletionReconciler
type MockMulticlusterDashboardDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterDashboardDeletionReconciler
}

// NewMockMulticlusterDashboardDeletionReconciler creates a new mock instance
func NewMockMulticlusterDashboardDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterDashboardDeletionReconciler {
	mock := &MockMulticlusterDashboardDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterDashboardDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterDashboardDeletionReconciler) EXPECT() *MockMulticlusterDashboardDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileDashboardDeletion mocks base method
func (m *MockMulticlusterDashboardDeletionReconciler) ReconcileDashboardDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileDashboardDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileDashboardDeletion indicates an expected call of ReconcileDashboardDeletion
func (mr *MockMulticlusterDashboardDeletionReconcilerMockRecorder) ReconcileDashboardDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileDashboardDeletion", reflect.TypeOf((*MockMulticlusterDashboardDeletionReconciler)(nil).ReconcileDashboardDeletion), clusterName, req)
}

// MockMulticlusterDashboardReconcileLoop is a mock of MulticlusterDashboardReconcileLoop interface
type MockMulticlusterDashboardReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterDashboardReconcileLoopMockRecorder
}

// MockMulticlusterDashboardReconcileLoopMockRecorder is the mock recorder for MockMulticlusterDashboardReconcileLoop
type MockMulticlusterDashboardReconcileLoopMockRecorder struct {
	mock *MockMulticlusterDashboardReconcileLoop
}

// NewMockMulticlusterDashboardReconcileLoop creates a new mock instance
func NewMockMulticlusterDashboardReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterDashboardReconcileLoop {
	mock := &MockMulticlusterDashboardReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterDashboardReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterDashboardReconcileLoop) EXPECT() *MockMulticlusterDashboardReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterDashboardReconciler mocks base method
func (m *MockMulticlusterDashboardReconcileLoop) AddMulticlusterDashboardReconciler(ctx context.Context, rec controller.MulticlusterDashboardReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterDashboardReconciler", varargs...)
}

// AddMulticlusterDashboardReconciler indicates an expected call of AddMulticlusterDashboardReconciler
func (mr *MockMulticlusterDashboardReconcileLoopMockRecorder) AddMulticlusterDashboardReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterDashboardReconciler", reflect.TypeOf((*MockMulticlusterDashboardReconcileLoop)(nil).AddMulticlusterDashboardReconciler), varargs...)
}