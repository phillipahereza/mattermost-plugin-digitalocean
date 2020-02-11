package main

import (
	"fmt"
	"strings"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
)

const commandHelp = `* |/do help| - Run 'test' to see if you're configured to run bamboo commands
* |/do connect <do account name>| - Points your current instance to work with Digital Ocean account whose name is provided
`

func getCommand() *model.Command {
	return &model.Command{
		Trigger:          "do",
		DisplayName:      "do",
		Description:      "Integration with Digital Ocean.",
		AutoComplete:     true,
		AutoCompleteDesc: "Available commands: help",
		AutoCompleteHint: "[command]",
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

// ExecuteCommand is
func (p *Plugin) ExecuteCommand(c *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	split := strings.Fields(args.Command)
	command := split[0]
	action := ""
	if len(split) > 1 {
		action = split[1]
	}

	if command != "/do" {
		return &model.CommandResponse{}, nil
	}

	switch action {
	case "help":
		return p.helpCommandFunc(args)
	case "connect":
		return p.connectCommandFunc(args)
	default:
		return p.responsef(args, fmt.Sprintf("Unknown action %v", action)), nil
	}
}

func (p *Plugin) isUserAuthorized(id string) (bool, *model.AppError) {
	user, appErr := p.API.GetUser(id)
	if appErr != nil {
		return false, appErr
	}
	if !strings.Contains(user.Roles, "system_admin") {
		return false, nil
	}

	return true, nil
}

func (p *Plugin) connectCommandFunc(args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	userID := args.UserId
	allowed, err := p.isUserAuthorized(userID)
	if err != nil {
		return &model.CommandResponse{}, err
	}

	if !allowed {
		return p.responsef(args, "You're not allowed to run this command"), &model.AppError{Message: "unauthorized"}
	}

	// TODO: Parse to pick the DO Acc URL and also return error when not provided as arg to /do connect
	url := args.Command

	instance := NewDOAccountInstance(p, url)
	stErr := p.store.StoreInstance(instance)
	if stErr != nil {
		return p.responsef(args, "Failed to store instance. Contact system admin"),
			&model.AppError{Message: stErr.Error()}
	}

	return p.responsef(args, fmt.Sprintf("Successfully connected to Digital Ocean at %s", url)), nil
}

func (p *Plugin) helpCommandFunc(args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	text := "###### Mattermost Digital Ocean Plugin - Slash Command Help\n" + strings.Replace(commandHelp, "|", "`", -1)
	return p.responsef(args, text), nil
}
