// Code generated by skv2. DO NOT EDIT.

package v1alpha1

import (
	fed_ratelimit_solo_io_v1alpha1 "github.com/solo-io/solo-apis/pkg/api/fed.ratelimit.solo.io/v1alpha1"

	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

/*
  The intention of these providers are to be used for Mocking.
  They expose the Clients as interfaces, as well as factories to provide mocked versions
  of the clients when they require building within a component.

  See package `github.com/solo-io/skv2/pkg/multicluster/register` for example
*/

// Provider for FederatedRateLimitConfigClient from Clientset
func FederatedRateLimitConfigClientFromClientsetProvider(clients fed_ratelimit_solo_io_v1alpha1.Clientset) fed_ratelimit_solo_io_v1alpha1.FederatedRateLimitConfigClient {
	return clients.FederatedRateLimitConfigs()
}

// Provider for FederatedRateLimitConfig Client from Client
func FederatedRateLimitConfigClientProvider(client client.Client) fed_ratelimit_solo_io_v1alpha1.FederatedRateLimitConfigClient {
	return fed_ratelimit_solo_io_v1alpha1.NewFederatedRateLimitConfigClient(client)
}

type FederatedRateLimitConfigClientFactory func(client client.Client) fed_ratelimit_solo_io_v1alpha1.FederatedRateLimitConfigClient

func FederatedRateLimitConfigClientFactoryProvider() FederatedRateLimitConfigClientFactory {
	return FederatedRateLimitConfigClientProvider
}

type FederatedRateLimitConfigClientFromConfigFactory func(cfg *rest.Config) (fed_ratelimit_solo_io_v1alpha1.FederatedRateLimitConfigClient, error)

func FederatedRateLimitConfigClientFromConfigFactoryProvider() FederatedRateLimitConfigClientFromConfigFactory {
	return func(cfg *rest.Config) (fed_ratelimit_solo_io_v1alpha1.FederatedRateLimitConfigClient, error) {
		clients, err := fed_ratelimit_solo_io_v1alpha1.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.FederatedRateLimitConfigs(), nil
	}
}
