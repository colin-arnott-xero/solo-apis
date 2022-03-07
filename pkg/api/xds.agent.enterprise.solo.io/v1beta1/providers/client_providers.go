// Code generated by skv2. DO NOT EDIT.

package v1beta1

import (
	xds_agent_enterprise_solo_io_v1beta1 "github.com/solo-io/solo-apis/pkg/api/xds.agent.enterprise.solo.io/v1beta1"

	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

/*
  The intention of these providers are to be used for Mocking.
  They expose the Clients as interfaces, as well as factories to provide mocked versions
  of the clients when they require building within a component.

  See package `github.com/solo-io/skv2/pkg/multicluster/register` for example
*/

// Provider for XdsConfigClient from Clientset
func XdsConfigClientFromClientsetProvider(clients xds_agent_enterprise_solo_io_v1beta1.Clientset) xds_agent_enterprise_solo_io_v1beta1.XdsConfigClient {
	return clients.XdsConfigs()
}

// Provider for XdsConfig Client from Client
func XdsConfigClientProvider(client client.Client) xds_agent_enterprise_solo_io_v1beta1.XdsConfigClient {
	return xds_agent_enterprise_solo_io_v1beta1.NewXdsConfigClient(client)
}

type XdsConfigClientFactory func(client client.Client) xds_agent_enterprise_solo_io_v1beta1.XdsConfigClient

func XdsConfigClientFactoryProvider() XdsConfigClientFactory {
	return XdsConfigClientProvider
}

type XdsConfigClientFromConfigFactory func(cfg *rest.Config) (xds_agent_enterprise_solo_io_v1beta1.XdsConfigClient, error)

func XdsConfigClientFromConfigFactoryProvider() XdsConfigClientFromConfigFactory {
	return func(cfg *rest.Config) (xds_agent_enterprise_solo_io_v1beta1.XdsConfigClient, error) {
		clients, err := xds_agent_enterprise_solo_io_v1beta1.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.XdsConfigs(), nil
	}
}