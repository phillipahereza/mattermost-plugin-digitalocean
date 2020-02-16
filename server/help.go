package main

import (
	"fmt"
	"github.com/mattermost/mattermost-server/v5/model"
	"strings"
)

const commandHelp = `* |/do help| - Run 'test' to see if you're configured to run do commands
* |/do connect <access token>| - Associates your DO team personal token with your mattermost account
* |/do disconnect| - Remove your DO team personal token with your mattermost account
* |/do token| - Provides instructions on getting a personal access token for the configured DigitalOcean team
* |/do show-configured-token| - Display your configured access token
* |/do list-droplets| - List all Droplets in your team
* |/do rename-droplet <dropletID> <name>| - Rename a droplet
* |/do reboot-droplet <dropletID>| - Reboot a droplet
* |/do shutdown-droplet <dropletID>| - Shutdown a droplet
* |/do powercycle-droplet <dropletID>| - action is similar to pushing the reset button on a physical machine, it's similar to booting from scratch
* |/do list-domains| - Retrieve a list of all of the domains in your team
* |/do list-keys| - Retrieve a list of all of SSH keys in your team
* |/do retrieve-key <keyID>| - Retrieve a single key by its ID
* |/do delete-key <keyID>| - Delete single key by its ID
* |/do create-key <name> <publicKey>| - Add an SSH key to your team. PublicKey is in double quotes
* |/do list-clusters| - Retrieve a list of all Database Clusters set up in your team
* |/do list-cluster-backups <id>| - Retrieve a list of all backups of a Database Cluster
* |/do add-cluster-user <clusterID> <userName>| - Add a database user to a cluster
* |/do list-cluster-users <clusterID>| - List database cluster users
* |/do delete-cluster-user <clusterID> <userName>| - Delete a database user to a cluster
* |/do list-k8s-clusters| - List all Kubernetes Clusters in your team
* |/do list-k8s-cluster-nodepools <clusterID>| - List Nodepools in a Kubernetes cluster
* |/do list-k8s-cluster-nodes <clusterID>| - List Nodes in a Kubernetes cluster
* |/do get-k8s-cluster-upgrades <clusterID>| - Retrieve a list of available upgrades for a Kubernetes cluster
* |/do get-k8s-kubeconfig <clusterID>| - Retrieve kubeconfig file in YAML format
* |/do upgrade-k8s-cluster <clusterID> <versionSlug>| - Upgrade a Kubernetes cluster to a newer patch release of Kubernetes
`

func (p *Plugin) helpCommandFunc(args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	text := "###### Mattermost DigitalOcean Plugin - Slash Command Help\n" + strings.Replace(commandHelp, "|", "`", -1)
	return p.responsef(args, text), nil
}

func (p *Plugin) defaultCommandFunc(args *model.CommandArgs, action string) (*model.CommandResponse, *model.AppError) {
	text := fmt.Sprintf("Unknown action %s. The following commands might help\n", action) + strings.Replace(commandHelp, "|", "`", -1)
	return p.responsef(args, text), nil
}
