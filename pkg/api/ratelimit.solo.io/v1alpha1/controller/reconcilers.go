// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./reconcilers.go -destination mocks/reconcilers.go

// Definitions for the Kubernetes Controllers
package controller



import (
	"context"

    ratelimit_solo_io_v1alpha1 "github.com/solo-io/solo-apis/pkg/api/ratelimit.solo.io/v1alpha1"

    "github.com/pkg/errors"
    "github.com/solo-io/skv2/pkg/ezkube"
    "github.com/solo-io/skv2/pkg/reconcile"
    "sigs.k8s.io/controller-runtime/pkg/manager"
    "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the RateLimitConfig Resource.
// implemented by the user
type RateLimitConfigReconciler interface {
    ReconcileRateLimitConfig(obj *ratelimit_solo_io_v1alpha1.RateLimitConfig) (reconcile.Result, error)
}

// Reconcile deletion events for the RateLimitConfig Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type RateLimitConfigDeletionReconciler interface {
    ReconcileRateLimitConfigDeletion(req reconcile.Request) error
}

type RateLimitConfigReconcilerFuncs struct {
    OnReconcileRateLimitConfig func(obj *ratelimit_solo_io_v1alpha1.RateLimitConfig) (reconcile.Result, error)
    OnReconcileRateLimitConfigDeletion func(req reconcile.Request) error
}

func (f *RateLimitConfigReconcilerFuncs) ReconcileRateLimitConfig(obj *ratelimit_solo_io_v1alpha1.RateLimitConfig) (reconcile.Result, error) {
    if f.OnReconcileRateLimitConfig == nil {
        return reconcile.Result{}, nil
    }
    return f.OnReconcileRateLimitConfig(obj)
}

func (f *RateLimitConfigReconcilerFuncs) ReconcileRateLimitConfigDeletion(req reconcile.Request) error {
    if f.OnReconcileRateLimitConfigDeletion == nil {
        return nil
    }
    return f.OnReconcileRateLimitConfigDeletion(req)
}

// Reconcile and finalize the RateLimitConfig Resource
// implemented by the user
type RateLimitConfigFinalizer interface {
    RateLimitConfigReconciler

    // name of the finalizer used by this handler.
    // finalizer names should be unique for a single task
    RateLimitConfigFinalizerName() string

    // finalize the object before it is deleted.
    // Watchers created with a finalizing handler will a
    FinalizeRateLimitConfig(obj *ratelimit_solo_io_v1alpha1.RateLimitConfig) error
}

type RateLimitConfigReconcileLoop interface {
    RunRateLimitConfigReconciler(ctx context.Context, rec RateLimitConfigReconciler, predicates ...predicate.Predicate) error
}

type rateLimitConfigReconcileLoop struct {
    loop reconcile.Loop
}

func NewRateLimitConfigReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) RateLimitConfigReconcileLoop {
    return &rateLimitConfigReconcileLoop{
    	// empty cluster indicates this reconciler is built for the local cluster
        loop: reconcile.NewLoop(name, "", mgr, &ratelimit_solo_io_v1alpha1.RateLimitConfig{}, options),
    }
}

func (c *rateLimitConfigReconcileLoop) RunRateLimitConfigReconciler(ctx context.Context, reconciler RateLimitConfigReconciler, predicates ...predicate.Predicate) error {
    genericReconciler := genericRateLimitConfigReconciler{
        reconciler: reconciler,
    }

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(RateLimitConfigFinalizer); ok {
        reconcilerWrapper = genericRateLimitConfigFinalizer{
            genericRateLimitConfigReconciler: genericReconciler,
            finalizingReconciler: finalizingReconciler,
        }
    } else {
        reconcilerWrapper = genericReconciler
    }
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericRateLimitConfigHandler implements a generic reconcile.Reconciler
type genericRateLimitConfigReconciler struct {
    reconciler RateLimitConfigReconciler
}

func (r genericRateLimitConfigReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
    obj, ok := object.(*ratelimit_solo_io_v1alpha1.RateLimitConfig)
    if !ok {
        return reconcile.Result{}, errors.Errorf("internal error: RateLimitConfig handler received event for %T", object)
    }
    return r.reconciler.ReconcileRateLimitConfig(obj)
}

func (r genericRateLimitConfigReconciler) ReconcileDeletion(request reconcile.Request) error {
    if deletionReconciler, ok := r.reconciler.(RateLimitConfigDeletionReconciler); ok {
        return deletionReconciler.ReconcileRateLimitConfigDeletion(request)
    }
    return nil
}

// genericRateLimitConfigFinalizer implements a generic reconcile.FinalizingReconciler
type genericRateLimitConfigFinalizer struct {
    genericRateLimitConfigReconciler
    finalizingReconciler RateLimitConfigFinalizer
}


func (r genericRateLimitConfigFinalizer) FinalizerName() string {
    return r.finalizingReconciler.RateLimitConfigFinalizerName()
}

func (r genericRateLimitConfigFinalizer) Finalize(object ezkube.Object) error {
    obj, ok := object.(*ratelimit_solo_io_v1alpha1.RateLimitConfig)
    if !ok {
        return errors.Errorf("internal error: RateLimitConfig handler received event for %T", object)
    }
    return r.finalizingReconciler.FinalizeRateLimitConfig(obj)
}
