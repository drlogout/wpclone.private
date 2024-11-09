package dock

import (
	"fmt"

	docker "github.com/fsouza/go-dockerclient"
)

func EnsureNetwork(client *docker.Client, name string) (*docker.Network, error) {
	network, err := GetNetwork(client, name)
	if err != nil {
		return nil, err
	}

	if network != nil {
		return network, nil
	}

	network, err = client.CreateNetwork(docker.CreateNetworkOptions{
		Name: name,
		Labels: map[string]string{
			name: name,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create network: %w", err)
	}

	return network, nil
}

func GetNetwork(client *docker.Client, networkName string) (*docker.Network, error) {
	networks, err := client.ListNetworks()
	if err != nil {
		return nil, fmt.Errorf("failed to list networks: %w", err)
	}

	for _, network := range networks {
		if network.Name == networkName {
			return &network, nil
		}
	}

	return nil, nil
}
