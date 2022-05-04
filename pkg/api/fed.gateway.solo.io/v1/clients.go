// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./clients.go -destination mocks/clients.go

package v1

import (
	"context"

	"github.com/solo-io/skv2/pkg/controllerutils"
	"github.com/solo-io/skv2/pkg/multicluster"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// MulticlusterClientset for the fed.gateway.solo.io/v1 APIs
type MulticlusterClientset interface {
	// Cluster returns a Clientset for the given cluster
	Cluster(cluster string) (Clientset, error)
}

type multiclusterClientset struct {
	client multicluster.Client
}

func NewMulticlusterClientset(client multicluster.Client) MulticlusterClientset {
	return &multiclusterClientset{client: client}
}

func (m *multiclusterClientset) Cluster(cluster string) (Clientset, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewClientset(client), nil
}

// clienset for the fed.gateway.solo.io/v1 APIs
type Clientset interface {
	// clienset for the fed.gateway.solo.io/v1/v1 APIs
	FederatedGateways() FederatedGatewayClient
	// clienset for the fed.gateway.solo.io/v1/v1 APIs
	FederatedMatchableHttpGateways() FederatedMatchableHttpGatewayClient
	// clienset for the fed.gateway.solo.io/v1/v1 APIs
	FederatedRouteTables() FederatedRouteTableClient
	// clienset for the fed.gateway.solo.io/v1/v1 APIs
	FederatedVirtualServices() FederatedVirtualServiceClient
}

type clientSet struct {
	client client.Client
}

func NewClientsetFromConfig(cfg *rest.Config) (Clientset, error) {
	scheme := scheme.Scheme
	if err := SchemeBuilder.AddToScheme(scheme); err != nil {
		return nil, err
	}
	client, err := client.New(cfg, client.Options{
		Scheme: scheme,
	})
	if err != nil {
		return nil, err
	}
	return NewClientset(client), nil
}

func NewClientset(client client.Client) Clientset {
	return &clientSet{client: client}
}

// clienset for the fed.gateway.solo.io/v1/v1 APIs
func (c *clientSet) FederatedGateways() FederatedGatewayClient {
	return NewFederatedGatewayClient(c.client)
}

// clienset for the fed.gateway.solo.io/v1/v1 APIs
func (c *clientSet) FederatedMatchableHttpGateways() FederatedMatchableHttpGatewayClient {
	return NewFederatedMatchableHttpGatewayClient(c.client)
}

// clienset for the fed.gateway.solo.io/v1/v1 APIs
func (c *clientSet) FederatedRouteTables() FederatedRouteTableClient {
	return NewFederatedRouteTableClient(c.client)
}

// clienset for the fed.gateway.solo.io/v1/v1 APIs
func (c *clientSet) FederatedVirtualServices() FederatedVirtualServiceClient {
	return NewFederatedVirtualServiceClient(c.client)
}

// Reader knows how to read and list FederatedGateways.
type FederatedGatewayReader interface {
	// Get retrieves a FederatedGateway for the given object key
	GetFederatedGateway(ctx context.Context, key client.ObjectKey) (*FederatedGateway, error)

	// List retrieves list of FederatedGateways for a given namespace and list options.
	ListFederatedGateway(ctx context.Context, opts ...client.ListOption) (*FederatedGatewayList, error)
}

// FederatedGatewayTransitionFunction instructs the FederatedGatewayWriter how to transition between an existing
// FederatedGateway object and a desired on an Upsert
type FederatedGatewayTransitionFunction func(existing, desired *FederatedGateway) error

// Writer knows how to create, delete, and update FederatedGateways.
type FederatedGatewayWriter interface {
	// Create saves the FederatedGateway object.
	CreateFederatedGateway(ctx context.Context, obj *FederatedGateway, opts ...client.CreateOption) error

	// Delete deletes the FederatedGateway object.
	DeleteFederatedGateway(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given FederatedGateway object.
	UpdateFederatedGateway(ctx context.Context, obj *FederatedGateway, opts ...client.UpdateOption) error

	// Patch patches the given FederatedGateway object.
	PatchFederatedGateway(ctx context.Context, obj *FederatedGateway, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all FederatedGateway objects matching the given options.
	DeleteAllOfFederatedGateway(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the FederatedGateway object.
	UpsertFederatedGateway(ctx context.Context, obj *FederatedGateway, transitionFuncs ...FederatedGatewayTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a FederatedGateway object.
type FederatedGatewayStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given FederatedGateway object.
	UpdateFederatedGatewayStatus(ctx context.Context, obj *FederatedGateway, opts ...client.UpdateOption) error

	// Patch patches the given FederatedGateway object's subresource.
	PatchFederatedGatewayStatus(ctx context.Context, obj *FederatedGateway, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on FederatedGateways.
type FederatedGatewayClient interface {
	FederatedGatewayReader
	FederatedGatewayWriter
	FederatedGatewayStatusWriter
}

type federatedGatewayClient struct {
	client client.Client
}

func NewFederatedGatewayClient(client client.Client) *federatedGatewayClient {
	return &federatedGatewayClient{client: client}
}

func (c *federatedGatewayClient) GetFederatedGateway(ctx context.Context, key client.ObjectKey) (*FederatedGateway, error) {
	obj := &FederatedGateway{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *federatedGatewayClient) ListFederatedGateway(ctx context.Context, opts ...client.ListOption) (*FederatedGatewayList, error) {
	list := &FederatedGatewayList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *federatedGatewayClient) CreateFederatedGateway(ctx context.Context, obj *FederatedGateway, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *federatedGatewayClient) DeleteFederatedGateway(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &FederatedGateway{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *federatedGatewayClient) UpdateFederatedGateway(ctx context.Context, obj *FederatedGateway, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *federatedGatewayClient) PatchFederatedGateway(ctx context.Context, obj *FederatedGateway, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *federatedGatewayClient) DeleteAllOfFederatedGateway(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &FederatedGateway{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *federatedGatewayClient) UpsertFederatedGateway(ctx context.Context, obj *FederatedGateway, transitionFuncs ...FederatedGatewayTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*FederatedGateway), desired.(*FederatedGateway)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *federatedGatewayClient) UpdateFederatedGatewayStatus(ctx context.Context, obj *FederatedGateway, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *federatedGatewayClient) PatchFederatedGatewayStatus(ctx context.Context, obj *FederatedGateway, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides FederatedGatewayClients for multiple clusters.
type MulticlusterFederatedGatewayClient interface {
	// Cluster returns a FederatedGatewayClient for the given cluster
	Cluster(cluster string) (FederatedGatewayClient, error)
}

type multiclusterFederatedGatewayClient struct {
	client multicluster.Client
}

func NewMulticlusterFederatedGatewayClient(client multicluster.Client) MulticlusterFederatedGatewayClient {
	return &multiclusterFederatedGatewayClient{client: client}
}

func (m *multiclusterFederatedGatewayClient) Cluster(cluster string) (FederatedGatewayClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewFederatedGatewayClient(client), nil
}

// Reader knows how to read and list FederatedMatchableHttpGateways.
type FederatedMatchableHttpGatewayReader interface {
	// Get retrieves a FederatedMatchableHttpGateway for the given object key
	GetFederatedMatchableHttpGateway(ctx context.Context, key client.ObjectKey) (*FederatedMatchableHttpGateway, error)

	// List retrieves list of FederatedMatchableHttpGateways for a given namespace and list options.
	ListFederatedMatchableHttpGateway(ctx context.Context, opts ...client.ListOption) (*FederatedMatchableHttpGatewayList, error)
}

// FederatedMatchableHttpGatewayTransitionFunction instructs the FederatedMatchableHttpGatewayWriter how to transition between an existing
// FederatedMatchableHttpGateway object and a desired on an Upsert
type FederatedMatchableHttpGatewayTransitionFunction func(existing, desired *FederatedMatchableHttpGateway) error

// Writer knows how to create, delete, and update FederatedMatchableHttpGateways.
type FederatedMatchableHttpGatewayWriter interface {
	// Create saves the FederatedMatchableHttpGateway object.
	CreateFederatedMatchableHttpGateway(ctx context.Context, obj *FederatedMatchableHttpGateway, opts ...client.CreateOption) error

	// Delete deletes the FederatedMatchableHttpGateway object.
	DeleteFederatedMatchableHttpGateway(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given FederatedMatchableHttpGateway object.
	UpdateFederatedMatchableHttpGateway(ctx context.Context, obj *FederatedMatchableHttpGateway, opts ...client.UpdateOption) error

	// Patch patches the given FederatedMatchableHttpGateway object.
	PatchFederatedMatchableHttpGateway(ctx context.Context, obj *FederatedMatchableHttpGateway, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all FederatedMatchableHttpGateway objects matching the given options.
	DeleteAllOfFederatedMatchableHttpGateway(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the FederatedMatchableHttpGateway object.
	UpsertFederatedMatchableHttpGateway(ctx context.Context, obj *FederatedMatchableHttpGateway, transitionFuncs ...FederatedMatchableHttpGatewayTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a FederatedMatchableHttpGateway object.
type FederatedMatchableHttpGatewayStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given FederatedMatchableHttpGateway object.
	UpdateFederatedMatchableHttpGatewayStatus(ctx context.Context, obj *FederatedMatchableHttpGateway, opts ...client.UpdateOption) error

	// Patch patches the given FederatedMatchableHttpGateway object's subresource.
	PatchFederatedMatchableHttpGatewayStatus(ctx context.Context, obj *FederatedMatchableHttpGateway, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on FederatedMatchableHttpGateways.
type FederatedMatchableHttpGatewayClient interface {
	FederatedMatchableHttpGatewayReader
	FederatedMatchableHttpGatewayWriter
	FederatedMatchableHttpGatewayStatusWriter
}

type federatedMatchableHttpGatewayClient struct {
	client client.Client
}

func NewFederatedMatchableHttpGatewayClient(client client.Client) *federatedMatchableHttpGatewayClient {
	return &federatedMatchableHttpGatewayClient{client: client}
}

func (c *federatedMatchableHttpGatewayClient) GetFederatedMatchableHttpGateway(ctx context.Context, key client.ObjectKey) (*FederatedMatchableHttpGateway, error) {
	obj := &FederatedMatchableHttpGateway{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *federatedMatchableHttpGatewayClient) ListFederatedMatchableHttpGateway(ctx context.Context, opts ...client.ListOption) (*FederatedMatchableHttpGatewayList, error) {
	list := &FederatedMatchableHttpGatewayList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *federatedMatchableHttpGatewayClient) CreateFederatedMatchableHttpGateway(ctx context.Context, obj *FederatedMatchableHttpGateway, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *federatedMatchableHttpGatewayClient) DeleteFederatedMatchableHttpGateway(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &FederatedMatchableHttpGateway{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *federatedMatchableHttpGatewayClient) UpdateFederatedMatchableHttpGateway(ctx context.Context, obj *FederatedMatchableHttpGateway, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *federatedMatchableHttpGatewayClient) PatchFederatedMatchableHttpGateway(ctx context.Context, obj *FederatedMatchableHttpGateway, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *federatedMatchableHttpGatewayClient) DeleteAllOfFederatedMatchableHttpGateway(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &FederatedMatchableHttpGateway{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *federatedMatchableHttpGatewayClient) UpsertFederatedMatchableHttpGateway(ctx context.Context, obj *FederatedMatchableHttpGateway, transitionFuncs ...FederatedMatchableHttpGatewayTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*FederatedMatchableHttpGateway), desired.(*FederatedMatchableHttpGateway)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *federatedMatchableHttpGatewayClient) UpdateFederatedMatchableHttpGatewayStatus(ctx context.Context, obj *FederatedMatchableHttpGateway, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *federatedMatchableHttpGatewayClient) PatchFederatedMatchableHttpGatewayStatus(ctx context.Context, obj *FederatedMatchableHttpGateway, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides FederatedMatchableHttpGatewayClients for multiple clusters.
type MulticlusterFederatedMatchableHttpGatewayClient interface {
	// Cluster returns a FederatedMatchableHttpGatewayClient for the given cluster
	Cluster(cluster string) (FederatedMatchableHttpGatewayClient, error)
}

type multiclusterFederatedMatchableHttpGatewayClient struct {
	client multicluster.Client
}

func NewMulticlusterFederatedMatchableHttpGatewayClient(client multicluster.Client) MulticlusterFederatedMatchableHttpGatewayClient {
	return &multiclusterFederatedMatchableHttpGatewayClient{client: client}
}

func (m *multiclusterFederatedMatchableHttpGatewayClient) Cluster(cluster string) (FederatedMatchableHttpGatewayClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewFederatedMatchableHttpGatewayClient(client), nil
}

// Reader knows how to read and list FederatedRouteTables.
type FederatedRouteTableReader interface {
	// Get retrieves a FederatedRouteTable for the given object key
	GetFederatedRouteTable(ctx context.Context, key client.ObjectKey) (*FederatedRouteTable, error)

	// List retrieves list of FederatedRouteTables for a given namespace and list options.
	ListFederatedRouteTable(ctx context.Context, opts ...client.ListOption) (*FederatedRouteTableList, error)
}

// FederatedRouteTableTransitionFunction instructs the FederatedRouteTableWriter how to transition between an existing
// FederatedRouteTable object and a desired on an Upsert
type FederatedRouteTableTransitionFunction func(existing, desired *FederatedRouteTable) error

// Writer knows how to create, delete, and update FederatedRouteTables.
type FederatedRouteTableWriter interface {
	// Create saves the FederatedRouteTable object.
	CreateFederatedRouteTable(ctx context.Context, obj *FederatedRouteTable, opts ...client.CreateOption) error

	// Delete deletes the FederatedRouteTable object.
	DeleteFederatedRouteTable(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given FederatedRouteTable object.
	UpdateFederatedRouteTable(ctx context.Context, obj *FederatedRouteTable, opts ...client.UpdateOption) error

	// Patch patches the given FederatedRouteTable object.
	PatchFederatedRouteTable(ctx context.Context, obj *FederatedRouteTable, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all FederatedRouteTable objects matching the given options.
	DeleteAllOfFederatedRouteTable(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the FederatedRouteTable object.
	UpsertFederatedRouteTable(ctx context.Context, obj *FederatedRouteTable, transitionFuncs ...FederatedRouteTableTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a FederatedRouteTable object.
type FederatedRouteTableStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given FederatedRouteTable object.
	UpdateFederatedRouteTableStatus(ctx context.Context, obj *FederatedRouteTable, opts ...client.UpdateOption) error

	// Patch patches the given FederatedRouteTable object's subresource.
	PatchFederatedRouteTableStatus(ctx context.Context, obj *FederatedRouteTable, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on FederatedRouteTables.
type FederatedRouteTableClient interface {
	FederatedRouteTableReader
	FederatedRouteTableWriter
	FederatedRouteTableStatusWriter
}

type federatedRouteTableClient struct {
	client client.Client
}

func NewFederatedRouteTableClient(client client.Client) *federatedRouteTableClient {
	return &federatedRouteTableClient{client: client}
}

func (c *federatedRouteTableClient) GetFederatedRouteTable(ctx context.Context, key client.ObjectKey) (*FederatedRouteTable, error) {
	obj := &FederatedRouteTable{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *federatedRouteTableClient) ListFederatedRouteTable(ctx context.Context, opts ...client.ListOption) (*FederatedRouteTableList, error) {
	list := &FederatedRouteTableList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *federatedRouteTableClient) CreateFederatedRouteTable(ctx context.Context, obj *FederatedRouteTable, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *federatedRouteTableClient) DeleteFederatedRouteTable(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &FederatedRouteTable{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *federatedRouteTableClient) UpdateFederatedRouteTable(ctx context.Context, obj *FederatedRouteTable, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *federatedRouteTableClient) PatchFederatedRouteTable(ctx context.Context, obj *FederatedRouteTable, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *federatedRouteTableClient) DeleteAllOfFederatedRouteTable(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &FederatedRouteTable{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *federatedRouteTableClient) UpsertFederatedRouteTable(ctx context.Context, obj *FederatedRouteTable, transitionFuncs ...FederatedRouteTableTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*FederatedRouteTable), desired.(*FederatedRouteTable)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *federatedRouteTableClient) UpdateFederatedRouteTableStatus(ctx context.Context, obj *FederatedRouteTable, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *federatedRouteTableClient) PatchFederatedRouteTableStatus(ctx context.Context, obj *FederatedRouteTable, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides FederatedRouteTableClients for multiple clusters.
type MulticlusterFederatedRouteTableClient interface {
	// Cluster returns a FederatedRouteTableClient for the given cluster
	Cluster(cluster string) (FederatedRouteTableClient, error)
}

type multiclusterFederatedRouteTableClient struct {
	client multicluster.Client
}

func NewMulticlusterFederatedRouteTableClient(client multicluster.Client) MulticlusterFederatedRouteTableClient {
	return &multiclusterFederatedRouteTableClient{client: client}
}

func (m *multiclusterFederatedRouteTableClient) Cluster(cluster string) (FederatedRouteTableClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewFederatedRouteTableClient(client), nil
}

// Reader knows how to read and list FederatedVirtualServices.
type FederatedVirtualServiceReader interface {
	// Get retrieves a FederatedVirtualService for the given object key
	GetFederatedVirtualService(ctx context.Context, key client.ObjectKey) (*FederatedVirtualService, error)

	// List retrieves list of FederatedVirtualServices for a given namespace and list options.
	ListFederatedVirtualService(ctx context.Context, opts ...client.ListOption) (*FederatedVirtualServiceList, error)
}

// FederatedVirtualServiceTransitionFunction instructs the FederatedVirtualServiceWriter how to transition between an existing
// FederatedVirtualService object and a desired on an Upsert
type FederatedVirtualServiceTransitionFunction func(existing, desired *FederatedVirtualService) error

// Writer knows how to create, delete, and update FederatedVirtualServices.
type FederatedVirtualServiceWriter interface {
	// Create saves the FederatedVirtualService object.
	CreateFederatedVirtualService(ctx context.Context, obj *FederatedVirtualService, opts ...client.CreateOption) error

	// Delete deletes the FederatedVirtualService object.
	DeleteFederatedVirtualService(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given FederatedVirtualService object.
	UpdateFederatedVirtualService(ctx context.Context, obj *FederatedVirtualService, opts ...client.UpdateOption) error

	// Patch patches the given FederatedVirtualService object.
	PatchFederatedVirtualService(ctx context.Context, obj *FederatedVirtualService, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all FederatedVirtualService objects matching the given options.
	DeleteAllOfFederatedVirtualService(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the FederatedVirtualService object.
	UpsertFederatedVirtualService(ctx context.Context, obj *FederatedVirtualService, transitionFuncs ...FederatedVirtualServiceTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a FederatedVirtualService object.
type FederatedVirtualServiceStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given FederatedVirtualService object.
	UpdateFederatedVirtualServiceStatus(ctx context.Context, obj *FederatedVirtualService, opts ...client.UpdateOption) error

	// Patch patches the given FederatedVirtualService object's subresource.
	PatchFederatedVirtualServiceStatus(ctx context.Context, obj *FederatedVirtualService, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on FederatedVirtualServices.
type FederatedVirtualServiceClient interface {
	FederatedVirtualServiceReader
	FederatedVirtualServiceWriter
	FederatedVirtualServiceStatusWriter
}

type federatedVirtualServiceClient struct {
	client client.Client
}

func NewFederatedVirtualServiceClient(client client.Client) *federatedVirtualServiceClient {
	return &federatedVirtualServiceClient{client: client}
}

func (c *federatedVirtualServiceClient) GetFederatedVirtualService(ctx context.Context, key client.ObjectKey) (*FederatedVirtualService, error) {
	obj := &FederatedVirtualService{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *federatedVirtualServiceClient) ListFederatedVirtualService(ctx context.Context, opts ...client.ListOption) (*FederatedVirtualServiceList, error) {
	list := &FederatedVirtualServiceList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *federatedVirtualServiceClient) CreateFederatedVirtualService(ctx context.Context, obj *FederatedVirtualService, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *federatedVirtualServiceClient) DeleteFederatedVirtualService(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &FederatedVirtualService{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *federatedVirtualServiceClient) UpdateFederatedVirtualService(ctx context.Context, obj *FederatedVirtualService, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *federatedVirtualServiceClient) PatchFederatedVirtualService(ctx context.Context, obj *FederatedVirtualService, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *federatedVirtualServiceClient) DeleteAllOfFederatedVirtualService(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &FederatedVirtualService{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *federatedVirtualServiceClient) UpsertFederatedVirtualService(ctx context.Context, obj *FederatedVirtualService, transitionFuncs ...FederatedVirtualServiceTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*FederatedVirtualService), desired.(*FederatedVirtualService)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *federatedVirtualServiceClient) UpdateFederatedVirtualServiceStatus(ctx context.Context, obj *FederatedVirtualService, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *federatedVirtualServiceClient) PatchFederatedVirtualServiceStatus(ctx context.Context, obj *FederatedVirtualService, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides FederatedVirtualServiceClients for multiple clusters.
type MulticlusterFederatedVirtualServiceClient interface {
	// Cluster returns a FederatedVirtualServiceClient for the given cluster
	Cluster(cluster string) (FederatedVirtualServiceClient, error)
}

type multiclusterFederatedVirtualServiceClient struct {
	client multicluster.Client
}

func NewMulticlusterFederatedVirtualServiceClient(client multicluster.Client) MulticlusterFederatedVirtualServiceClient {
	return &multiclusterFederatedVirtualServiceClient{client: client}
}

func (m *multiclusterFederatedVirtualServiceClient) Cluster(cluster string) (FederatedVirtualServiceClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewFederatedVirtualServiceClient(client), nil
}
