// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./reconcilers.go -destination mocks/reconcilers.go

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	settings_solo_io_v1 "github.com/solo-io/solo-apis/pkg/api/settings.solo.io/v1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the Settings Resource.
// implemented by the user
type SettingsReconciler interface {
	ReconcileSettings(obj *settings_solo_io_v1.Settings) (reconcile.Result, error)
}

// Reconcile deletion events for the Settings Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type SettingsDeletionReconciler interface {
	ReconcileSettingsDeletion(req reconcile.Request) error
}

type SettingsReconcilerFuncs struct {
	OnReconcileSettings         func(obj *settings_solo_io_v1.Settings) (reconcile.Result, error)
	OnReconcileSettingsDeletion func(req reconcile.Request) error
}

func (f *SettingsReconcilerFuncs) ReconcileSettings(obj *settings_solo_io_v1.Settings) (reconcile.Result, error) {
	if f.OnReconcileSettings == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileSettings(obj)
}

func (f *SettingsReconcilerFuncs) ReconcileSettingsDeletion(req reconcile.Request) error {
	if f.OnReconcileSettingsDeletion == nil {
		return nil
	}
	return f.OnReconcileSettingsDeletion(req)
}

// Reconcile and finalize the Settings Resource
// implemented by the user
type SettingsFinalizer interface {
	SettingsReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	SettingsFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeSettings(obj *settings_solo_io_v1.Settings) error
}

type SettingsReconcileLoop interface {
	RunSettingsReconciler(ctx context.Context, rec SettingsReconciler, predicates ...predicate.Predicate) error
}

type settingsReconcileLoop struct {
	loop reconcile.Loop
}

func NewSettingsReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) SettingsReconcileLoop {
	return &settingsReconcileLoop{
		// empty cluster indicates this reconciler is built for the local cluster
		loop: reconcile.NewLoop(name, "", mgr, &settings_solo_io_v1.Settings{}, options),
	}
}

func (c *settingsReconcileLoop) RunSettingsReconciler(ctx context.Context, reconciler SettingsReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericSettingsReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(SettingsFinalizer); ok {
		reconcilerWrapper = genericSettingsFinalizer{
			genericSettingsReconciler: genericReconciler,
			finalizingReconciler:      finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericSettingsHandler implements a generic reconcile.Reconciler
type genericSettingsReconciler struct {
	reconciler SettingsReconciler
}

func (r genericSettingsReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*settings_solo_io_v1.Settings)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: Settings handler received event for %T", object)
	}
	return r.reconciler.ReconcileSettings(obj)
}

func (r genericSettingsReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(SettingsDeletionReconciler); ok {
		return deletionReconciler.ReconcileSettingsDeletion(request)
	}
	return nil
}

// genericSettingsFinalizer implements a generic reconcile.FinalizingReconciler
type genericSettingsFinalizer struct {
	genericSettingsReconciler
	finalizingReconciler SettingsFinalizer
}

func (r genericSettingsFinalizer) FinalizerName() string {
	return r.finalizingReconciler.SettingsFinalizerName()
}

func (r genericSettingsFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*settings_solo_io_v1.Settings)
	if !ok {
		return errors.Errorf("internal error: Settings handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeSettings(obj)
}

// Reconcile Upsert events for the Dashboard Resource.
// implemented by the user
type DashboardReconciler interface {
	ReconcileDashboard(obj *settings_solo_io_v1.Dashboard) (reconcile.Result, error)
}

// Reconcile deletion events for the Dashboard Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type DashboardDeletionReconciler interface {
	ReconcileDashboardDeletion(req reconcile.Request) error
}

type DashboardReconcilerFuncs struct {
	OnReconcileDashboard         func(obj *settings_solo_io_v1.Dashboard) (reconcile.Result, error)
	OnReconcileDashboardDeletion func(req reconcile.Request) error
}

func (f *DashboardReconcilerFuncs) ReconcileDashboard(obj *settings_solo_io_v1.Dashboard) (reconcile.Result, error) {
	if f.OnReconcileDashboard == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileDashboard(obj)
}

func (f *DashboardReconcilerFuncs) ReconcileDashboardDeletion(req reconcile.Request) error {
	if f.OnReconcileDashboardDeletion == nil {
		return nil
	}
	return f.OnReconcileDashboardDeletion(req)
}

// Reconcile and finalize the Dashboard Resource
// implemented by the user
type DashboardFinalizer interface {
	DashboardReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	DashboardFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeDashboard(obj *settings_solo_io_v1.Dashboard) error
}

type DashboardReconcileLoop interface {
	RunDashboardReconciler(ctx context.Context, rec DashboardReconciler, predicates ...predicate.Predicate) error
}

type dashboardReconcileLoop struct {
	loop reconcile.Loop
}

func NewDashboardReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) DashboardReconcileLoop {
	return &dashboardReconcileLoop{
		// empty cluster indicates this reconciler is built for the local cluster
		loop: reconcile.NewLoop(name, "", mgr, &settings_solo_io_v1.Dashboard{}, options),
	}
}

func (c *dashboardReconcileLoop) RunDashboardReconciler(ctx context.Context, reconciler DashboardReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericDashboardReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(DashboardFinalizer); ok {
		reconcilerWrapper = genericDashboardFinalizer{
			genericDashboardReconciler: genericReconciler,
			finalizingReconciler:       finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericDashboardHandler implements a generic reconcile.Reconciler
type genericDashboardReconciler struct {
	reconciler DashboardReconciler
}

func (r genericDashboardReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*settings_solo_io_v1.Dashboard)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: Dashboard handler received event for %T", object)
	}
	return r.reconciler.ReconcileDashboard(obj)
}

func (r genericDashboardReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(DashboardDeletionReconciler); ok {
		return deletionReconciler.ReconcileDashboardDeletion(request)
	}
	return nil
}

// genericDashboardFinalizer implements a generic reconcile.FinalizingReconciler
type genericDashboardFinalizer struct {
	genericDashboardReconciler
	finalizingReconciler DashboardFinalizer
}

func (r genericDashboardFinalizer) FinalizerName() string {
	return r.finalizingReconciler.DashboardFinalizerName()
}

func (r genericDashboardFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*settings_solo_io_v1.Dashboard)
	if !ok {
		return errors.Errorf("internal error: Dashboard handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeDashboard(obj)
}