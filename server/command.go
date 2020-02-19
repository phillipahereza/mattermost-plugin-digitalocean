package main

import (
	"fmt"
	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
	"strconv"
	"strings"
)

// ExecuteCommand is
func (p *Plugin) ExecuteCommand(c *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {

	split := strings.Fields(args.Command)
	command := split[0]
	parameters := []string{}
	action := ""
	if len(split) > 1 {
		action = split[1]
	}

	if len(split) > 2 {
		parameters = split[2:]
	}

	if command != "/do" {
		return &model.CommandResponse{}, nil
	}

	// actions that don't make calls to the DigitalOcean API
	switch action {
	case "":
		return p.helpCommandFunc(args)
	case "help":
		return p.helpCommandFunc(args)
	case "connect":
		return p.connectCommandFunc(args)
	case "disconnect":
		return p.disconnectCommandFunc(args)
	case "token":
		return p.getPersonalTokenCommandFunc(args)
	case "show-configured-token":
		return p.showConnectTokenCommandFunc(args)
	case "subscribe":
		return p.subscribeCommandFunc(args)
	}

	client, err := p.GetClient(args.UserId)
	if err != nil {
		p.API.LogError("Failed to get digitalOcean client", "Err", err.Error())
		return p.responsef(args, "Failed to get DigitalOcean client: %s", err.Error()),
			&model.AppError{Message: err.Error()}
	}

	switch action {
	case "list-droplets":
		return p.listDropletsCommandFunc(client, args)
	case "reboot-droplet":
		if len(parameters) == 0 {
			return p.responsef(args, "Please specify the droplet ID"), nil
		} else if len(parameters) == 1 {
			dropletID, err := strconv.Atoi(parameters[0])
			if err != nil {
				return p.responsef(args, "Droplet ID must be an integer"), nil
			}
			return p.rebootDropletCommandFunc(client, args, dropletID)
		} else {
			return p.responsef(args, "Too many arguments, command should be in the form `/do reboot-droplet <dropletID>`"), nil
		}
	case "shutdown-droplet":
		if len(parameters) == 0 {
			return p.responsef(args, "Please specify the droplet ID"), nil
		} else if len(parameters) == 1 {
			dropletID, err := strconv.Atoi(parameters[0])
			if err != nil {
				return p.responsef(args, "Droplet ID must be an integer"), nil
			}
			return p.shutdownDropletCommandFunc(client, args, dropletID)
		} else {
			return p.responsef(args, "Too many arguments, command should be in the form `/do shutdown-droplet <dropletID>`"), nil
		}
	case "powercycle-droplet":
		if len(parameters) == 0 {
			return p.responsef(args, "Please specify the droplet ID"), nil
		} else if len(parameters) == 1 {
			dropletID, err := strconv.Atoi(parameters[0])
			if err != nil {
				return p.responsef(args, "Droplet ID must be an integer"), nil
			}
			return p.powercycleDropletCommandFunc(client, args, dropletID)
		} else {
			return p.responsef(args, "Too many arguments, command should be in the form `/do powercycle-droplet <dropletID>`"), nil
		}
	case "rename-droplet":
		if len(parameters) < 2 {
			return p.responsef(args, "Please specify the droplet ID or and the new name in the form `/do rename-droplet <dropletID> <name>`"), nil
		} else if len(parameters) == 2 {
			dropletID, err := strconv.Atoi(parameters[0])
			if err != nil {
				return p.responsef(args, "Droplet ID must be an integer"), nil
			}
			newName := parameters[1]
			return p.renameDropletCommandFunc(client, args, dropletID, newName)
		} else {
			return p.responsef(args, "Too many arguments, command should be in the form `/do rename-droplet <dropletID> <name>`"), nil
		}
	case "list-domains":
		return p.listDomainsCommandFunc(client, args)
	case "list-keys":
		return p.listSSHKeysCommandFunc(client, args)
	case "create-key":
		p.API.LogInfo("%v", parameters)
		if len(parameters) < 2 {
			return p.responsef(args, "Please specify the key name or and the publicKey in the form `/do create-key <name> <publicKey>`"), nil
		} else if len(parameters) >= 2 {
			name := parameters[0]
			publicKey := strings.Join(parameters[1:], " ")
			return p.createSSHKeysCommandFunc(client, args, name, publicKey)
		} else {
			return p.responsef(args, "Too many arguments, command should be in the form `/do create-key <name> <publicKey>`"), nil
		}
	case "retrieve-key":
		if len(parameters) == 0 {
			return p.responsef(args, "Please specify the SSH key ID"), nil
		} else if len(parameters) == 1 {
			dropletID, err := strconv.Atoi(parameters[0])
			if err != nil {
				return p.responsef(args, "SSH key ID must be an integer"), nil
			}
			return p.retrieveSSHKeyCommandFunc(client, args, dropletID)
		} else {
			return p.responsef(args, "Too many arguments, command should be in the form `/do retrieve-key <keyID>`"), nil
		}
	case "delete-key":
		if len(parameters) == 0 {
			return p.responsef(args, "Please specify the SSH key ID"), nil
		} else if len(parameters) == 1 {
			dropletID, err := strconv.Atoi(parameters[0])
			if err != nil {
				return p.responsef(args, "SSH key ID must be an integer"), nil
			}
			return p.deleteSSHKeyCommandFunc(client, args, dropletID)
		} else {
			return p.responsef(args, "Too many arguments, command should be in the form `/do delete-key <keyID>`"), nil
		}
	case "list-clusters":
		return p.listDatabaseClustersCommandFunc(client, args)
	case "list-cluster-backups":
		if len(parameters) == 0 {
			return p.responsef(args, "Please specify the Cluster ID"), nil
		} else if len(parameters) == 1 {
			clusterID := parameters[0]
			return p.listDatabaseClusterBackupsCommandFunc(client, args, clusterID)
		} else {
			return p.responsef(args, "Too many arguments, command should be in the form `/do retrieve-key <keyID>`"), nil
		}
	case "add-cluster-user":
		if len(parameters) < 2 {
			return p.responsef(args, "Please specify the cluster ID or and the name in the form `/do add-cluster-user <clusterID> <userName>`"), nil
		} else if len(parameters) == 2 {
			clusterID := parameters[0]
			userName := parameters[1]
			return p.addUserToDatabaseClusterCommandFunc(client, args, clusterID, userName)
		} else {
			return p.responsef(args, "Too many arguments, command should be in the form `/do add-cluster-user <clusterID> <userName>`"), nil
		}
	case "list-cluster-users":
		if len(parameters) == 0 {
			return p.responsef(args, "Please specify the database cluster ID"), nil
		} else if len(parameters) == 1 {
			clusterID := parameters[0]
			return p.listDatabaseClusterUsersCommandFunc(client, args, clusterID)
		} else {
			return p.responsef(args, "Too many arguments, command should be in the form `/do list-cluster-users <clusterID>`"), nil
		}
	case "delete-cluster-user":
		if len(parameters) < 2 {
			return p.responsef(args, "Please specify the cluster ID or and the name in the form `/do delete-cluster-user <clusterID> <userName>`"), nil
		} else if len(parameters) == 2 {
			clusterID := parameters[0]
			userName := parameters[1]
			return p.deleteDatabaseClusterUserCommandFunc(client, args, clusterID, userName)
		} else {
			return p.responsef(args, "Too many arguments, command should be in the form `/do delete-cluster-user <clusterID> <userName>`"), nil
		}
	case "list-cluster-dbs":
		if len(parameters) == 0 {
			return p.responsef(args, "Please specify the database cluster ID"), nil
		} else if len(parameters) == 1 {
			clusterID := parameters[0]
			return p.listDatabasesInClusterCommandFunc(client, args, clusterID)
		} else {
			return p.responsef(args, "Too many arguments, command should be in the form `/do list-cluster-dbs <clusterID>`"), nil
		}

	case "list-k8s-clusters":
		return p.listKubernetesClustersCommandFunc(client, args)

	case "list-k8s-cluster-nodepools":
		if len(parameters) == 0 {
			return p.responsef(args, "Please specify the Kubernetes cluster ID"), nil
		} else if len(parameters) == 1 {
			clusterID := parameters[0]
			return p.listKubernetesClusterNodePoolsCommandFunc(client, args, clusterID)
		} else {
			return p.responsef(args, "Too many arguments, command should be in the form `/do list-k8s-cluster-nodepools <clusterID>`"), nil
		}
	case "list-k8s-cluster-nodes":
		if len(parameters) == 0 {
			return p.responsef(args, "Please specify the Kubernetes cluster ID"), nil
		} else if len(parameters) == 1 {
			clusterID := parameters[0]
			return p.listKubernetesClusterNodesCommandFunc(client, args, clusterID)
		} else {
			return p.responsef(args, "Too many arguments, command should be in the form `/do list-k8s-cluster-node <clusterID>`"), nil
		}
	case "list-k8s-cluster-upgrades":
		if len(parameters) == 0 {
			return p.responsef(args, "Please specify the Kubernetes cluster ID"), nil
		} else if len(parameters) == 1 {
			clusterID := parameters[0]
			return p.retrieveAvailableUpgradesForKubernetesClusterCommandFunc(client, args, clusterID)
		} else {
			return p.responsef(args, "Too many arguments, command should be in the form `/do list-k8s-cluster-upgrades <clusterID>`"), nil
		}
	case "upgrade-k8s-cluster":
		if len(parameters) < 2 {
			return p.responsef(args, "Please specify the cluster ID or/and the versionSlug in the form `/do upgrade-k8s-cluster <clusterID> <versionSlug>`"), nil
		} else if len(parameters) == 2 {
			clusterID := parameters[0]
			versionSlug := parameters[1]
			return p.upgradeKubernetesClusterCommandFunc(client, args, clusterID, versionSlug)
		} else {
			return p.responsef(args, "Too many arguments, command should be in the form `/do upgrade-k8s-cluster <clusterID> <versionSlug>`"), nil
		}
	case "get-k8s-config":
		if len(parameters) == 0 {
			return p.responsef(args, "Please specify the Kubernetes cluster ID"), nil
		} else if len(parameters) == 1 {
			clusterID := parameters[0]
			return p.retrieveKubeconfigCommandFunc(client, args, clusterID)
		} else {
			return p.responsef(args, "Too many arguments, command should be in the form `/do get-k8s-config <clusterID>`"), nil
		}

	default:
		return p.defaultCommandFunc(args, action)
	}
}

func (p *Plugin) postCommandResponse(args *model.CommandArgs, message string) {
	post := &model.Post{
		UserId:    p.BotUserID,
		ChannelId: args.ChannelId,
		Message:   message,
	}

	_ = p.API.SendEphemeralPost(args.UserId, post)
}

func (p *Plugin) responsef(commandArgs *model.CommandArgs, format string, args ...interface{}) *model.CommandResponse {
	p.postCommandResponse(commandArgs, fmt.Sprintf(format, args...))
	return &model.CommandResponse{}
}
