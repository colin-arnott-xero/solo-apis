// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./reconcilers.go -destination mocks/reconcilers.go

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	graphql_gloo_solo_io_v1alpha1 "github.com/solo-io/solo-apis/pkg/api/graphql.gloo.solo.io/v1alpha1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the GraphQLSchema Resource.
// implemented by the user
type GraphQLSchemaReconciler interface {
	ReconcileGraphQLSchema(obj *graphql_gloo_solo_io_v1alpha1.GraphQLSchema) (reconcile.Result, error)
}

// Reconcile deletion events for the GraphQLSchema Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type GraphQLSchemaDeletionReconciler interface {
	ReconcileGraphQLSchemaDeletion(req reconcile.Request) error
}

type GraphQLSchemaReconcilerFuncs struct {
	OnReconcileGraphQLSchema         func(obj *graphql_gloo_solo_io_v1alpha1.GraphQLSchema) (reconcile.Result, error)
	OnReconcileGraphQLSchemaDeletion func(req reconcile.Request) error
}

func (f *GraphQLSchemaReconcilerFuncs) ReconcileGraphQLSchema(obj *graphql_gloo_solo_io_v1alpha1.GraphQLSchema) (reconcile.Result, error) {
	if f.OnReconcileGraphQLSchema == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileGraphQLSchema(obj)
}

func (f *GraphQLSchemaReconcilerFuncs) ReconcileGraphQLSchemaDeletion(req reconcile.Request) error {
	if f.OnReconcileGraphQLSchemaDeletion == nil {
		return nil
	}
	return f.OnReconcileGraphQLSchemaDeletion(req)
}

// Reconcile and finalize the GraphQLSchema Resource
// implemented by the user
type GraphQLSchemaFinalizer interface {
	GraphQLSchemaReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	GraphQLSchemaFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeGraphQLSchema(obj *graphql_gloo_solo_io_v1alpha1.GraphQLSchema) error
}

type GraphQLSchemaReconcileLoop interface {
	RunGraphQLSchemaReconciler(ctx context.Context, rec GraphQLSchemaReconciler, predicates ...predicate.Predicate) error
}

type graphQLSchemaReconcileLoop struct {
	loop reconcile.Loop
}

func NewGraphQLSchemaReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) GraphQLSchemaReconcileLoop {
	return &graphQLSchemaReconcileLoop{
		// empty cluster indicates this reconciler is built for the local cluster
		loop: reconcile.NewLoop(name, "", mgr, &graphql_gloo_solo_io_v1alpha1.GraphQLSchema{}, options),
	}
}

func (c *graphQLSchemaReconcileLoop) RunGraphQLSchemaReconciler(ctx context.Context, reconciler GraphQLSchemaReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericGraphQLSchemaReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(GraphQLSchemaFinalizer); ok {
		reconcilerWrapper = genericGraphQLSchemaFinalizer{
			genericGraphQLSchemaReconciler: genericReconciler,
			finalizingReconciler:           finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericGraphQLSchemaHandler implements a generic reconcile.Reconciler
type genericGraphQLSchemaReconciler struct {
	reconciler GraphQLSchemaReconciler
}

func (r genericGraphQLSchemaReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*graphql_gloo_solo_io_v1alpha1.GraphQLSchema)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: GraphQLSchema handler received event for %T", object)
	}
	return r.reconciler.ReconcileGraphQLSchema(obj)
}

func (r genericGraphQLSchemaReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(GraphQLSchemaDeletionReconciler); ok {
		return deletionReconciler.ReconcileGraphQLSchemaDeletion(request)
	}
	return nil
}

// genericGraphQLSchemaFinalizer implements a generic reconcile.FinalizingReconciler
type genericGraphQLSchemaFinalizer struct {
	genericGraphQLSchemaReconciler
	finalizingReconciler GraphQLSchemaFinalizer
}

func (r genericGraphQLSchemaFinalizer) FinalizerName() string {
	return r.finalizingReconciler.GraphQLSchemaFinalizerName()
}

func (r genericGraphQLSchemaFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*graphql_gloo_solo_io_v1alpha1.GraphQLSchema)
	if !ok {
		return errors.Errorf("internal error: GraphQLSchema handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeGraphQLSchema(obj)
}