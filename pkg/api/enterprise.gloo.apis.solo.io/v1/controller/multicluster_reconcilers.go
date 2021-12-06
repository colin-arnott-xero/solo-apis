// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./multicluster_reconcilers.go -destination mocks/multicluster_reconcilers.go

// Definitions for the multicluster Kubernetes Controllers
package controller



import (
	"context"

	enterprise_gloo_apis_solo_io_v1 "github.com/solo-io/solo-apis/pkg/api/enterprise.gloo.apis.solo.io/v1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/multicluster"
	mc_reconcile "github.com/solo-io/skv2/pkg/multicluster/reconcile"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the AuthConfig Resource across clusters.
// implemented by the user
type MulticlusterAuthConfigReconciler interface {
	ReconcileAuthConfig(clusterName string, obj *enterprise_gloo_apis_solo_io_v1.AuthConfig) (reconcile.Result, error)
}

// Reconcile deletion events for the AuthConfig Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterAuthConfigDeletionReconciler interface {
	ReconcileAuthConfigDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterAuthConfigReconcilerFuncs struct {
	OnReconcileAuthConfig         func(clusterName string, obj *enterprise_gloo_apis_solo_io_v1.AuthConfig) (reconcile.Result, error)
	OnReconcileAuthConfigDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterAuthConfigReconcilerFuncs) ReconcileAuthConfig(clusterName string, obj *enterprise_gloo_apis_solo_io_v1.AuthConfig) (reconcile.Result, error) {
	if f.OnReconcileAuthConfig == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileAuthConfig(clusterName, obj)
}

func (f *MulticlusterAuthConfigReconcilerFuncs) ReconcileAuthConfigDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileAuthConfigDeletion == nil {
		return nil
	}
	return f.OnReconcileAuthConfigDeletion(clusterName, req)
}

type MulticlusterAuthConfigReconcileLoop interface {
	// AddMulticlusterAuthConfigReconciler adds a MulticlusterAuthConfigReconciler to the MulticlusterAuthConfigReconcileLoop.
	AddMulticlusterAuthConfigReconciler(ctx context.Context, rec MulticlusterAuthConfigReconciler, predicates ...predicate.Predicate)
}

type multiclusterAuthConfigReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterAuthConfigReconcileLoop) AddMulticlusterAuthConfigReconciler(ctx context.Context, rec MulticlusterAuthConfigReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericAuthConfigMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterAuthConfigReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterAuthConfigReconcileLoop {
	return &multiclusterAuthConfigReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &enterprise_gloo_apis_solo_io_v1.AuthConfig{}, options)}
}

type genericAuthConfigMulticlusterReconciler struct {
	reconciler MulticlusterAuthConfigReconciler
}

func (g genericAuthConfigMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterAuthConfigDeletionReconciler); ok {
		return deletionReconciler.ReconcileAuthConfigDeletion(cluster, req)
	}
	return nil
}

func (g genericAuthConfigMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*enterprise_gloo_apis_solo_io_v1.AuthConfig)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: AuthConfig handler received event for %T", object)
	}
	return g.reconciler.ReconcileAuthConfig(cluster, obj)
}
