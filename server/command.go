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
		return p.showConnectTokenFunc(args)
	case "list-droplets":
		return p.listDropletsFunc(args)
	case "reboot-droplet":
		if len(parameters) == 0 {
			return p.responsef(args, "Please specify the droplet ID"), nil
		} else if len(parameters) == 1 {
			dropletID, err := strconv.Atoi(parameters[0])
			if err != nil {
				return p.responsef(args, "Droplet ID must be an integer"), nil
			}
			return p.rebootDropletFunc(args, dropletID)
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
			return p.shutdownDropletFunc(args, dropletID)
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
			return p.powercycleDropletFunc(args, dropletID)
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
			return p.renameDropletFunc(args, dropletID, newName)
		} else {
			return p.responsef(args, "Too many arguments, command should be in the form `/do rename-droplet <dropletID> <name>`"), nil
		}
	case "list-domains":
		return p.listDomainsFunc(args)
	case "list-keys":
		return p.listSSHKeysFunc(args)
	case "create-key":
		p.API.LogInfo("%v", parameters)
		if len(parameters) < 2 {
			return p.responsef(args, "Please specify the key name or and the publicKey in the form `/do create-key <name> <publicKey>`"), nil
		} else if len(parameters) >= 2 {
			name := parameters[0]
			publicKey := strings.Join(parameters[1:], " ")
			return p.createSSHKeysFunc(args, name, publicKey)
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
			return p.retrieveSSHKeyFunc(args, dropletID)
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
			return p.deleteSSHKeyFunc(args, dropletID)
		} else {
			return p.responsef(args, "Too many arguments, command should be in the form `/do delete-key <keyID>`"), nil
		}
	case "list-clusters":
		return p.listClustersFunc(args)
	case "list-cluster-backups":
		if len(parameters) == 0 {
			return p.responsef(args, "Please specify the Cluster ID"), nil
		} else if len(parameters) == 1 {
			clusterID := parameters[0]
			return p.listClusterBackupsFunc(args, clusterID)
		} else {
			return p.responsef(args, "Too many arguments, command should be in the form `/do retrieve-key <keyID>`"), nil
		}
	case "add-cluster-user":
		if len(parameters) < 2 {
			return p.responsef(args, "Please specify the cluster ID or and the name in the form `/do add-cluster-user <clusterID> <userName>`"), nil
		} else if len(parameters) == 2 {
			clusterID := parameters[0]
			userName := parameters[1]
			return p.addUserToClusterFunc(args, clusterID, userName)
		} else {
			return p.responsef(args, "Too many arguments, command should be in the form `/do add-cluster-user <clusterID> <userName>`"), nil
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
