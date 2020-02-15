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
	default:
		return p.responsef(args, fmt.Sprintf("Unknown action %v", action)), nil
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
