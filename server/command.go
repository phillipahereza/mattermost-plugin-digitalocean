package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/digitalocean/godo"
	"strings"
	"text/tabwriter"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
)

const commandHelp = `* |/do help| - Run 'test' to see if you're configured to run bamboo commands
* |/do connect <access token>| - Associates your DO team personal token with your mattermost account
* |/do token| - Provides instructions on getting a personal access token for the configured Digital Ocean team
* |/do show-configured-token| - Display your configured access token
* |/do list-droplets| - List all Droplets in your account 
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
	case "token":
		return p.getPersonalTokenCommandFunc(args)
	case "show-configured-token":
		return p.showConnectTokenFunc(args)
	case "list-droplets":
		return p.listDropletsFunc(args)
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

func (p *Plugin) getPersonalTokenCommandFunc(args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	const responseMessage = `Click the link below and follow subsequent steps

1. [Digital Ocean Apps and APIs](/plugins/com.mattermost.digitalocean/%s)
2. Generate a token and add copy it
3. Run |/do connect <your-token>|
`
	tID := p.getConfiguration().DOTeamID
	if tID == "" {
		return p.responsef(args, "No team was set by system admin"),
			&model.AppError{Message: "No team id in config"}
	}

	return p.responsef(args, fmt.Sprintf(responseMessage, routeToDOApps)), nil
}

func (p *Plugin) showConnectTokenFunc(args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	tk, err := p.store.LoadUserDOToken(args.UserId)
	if err != nil {
		return p.responsef(args, "Failed to retrieve token."),
			&model.AppError{Message: err.Error()}
	}

	if tk == "" {
		return p.responsef(args, "No token was found"),
			&model.AppError{Message: "Empty token"}
	}

	return p.responsef(args, fmt.Sprintf("Your token: %s", tk)), nil
}

func extractTokenFromCommand(command string) string {
	tk := strings.Fields(command)[2]
	return strings.TrimSpace(tk)
}

func (p *Plugin) connectCommandFunc(args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	userID := args.UserId

	token := extractTokenFromCommand(args.Command)
	stErr := p.store.StoreUserDOToken(token, userID)
	if stErr != nil {
		return p.responsef(args, "Failed to store token. Contact system admin"),
			&model.AppError{Message: stErr.Error()}
	}

	return p.responsef(args, "Successfully added a connecting token"), nil
}

func (p *Plugin) listDropletsFunc(args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	client, err := p.GetClient(args.UserId)
	if err != nil {
		return p.responsef(args, "Failed to get DigitalOcean client"),
			&model.AppError{Message: err.Error()}
	}

	opts := &godo.ListOptions{}

	droplets, _, err := client.Droplets.List(context.TODO(), opts)

	if err != nil {
		return p.responsef(args, "Error while fetching droplets list"),
			&model.AppError{Message: err.Error()}
	}

	if len(droplets) == 0 {
		return p.responsef(args, "You don't have any droplets configured. Use `do droplet create` to provision on"), nil
	}

	buffer := new(bytes.Buffer)

	// initialize the tabwriter
	w := new(tabwriter.Writer)

	w.Init(buffer, 8, 8, 0, '\t', 0)
	fmt.Fprintf(w, "\n |%s|%s|%s|%s|%s|", "ID", "Name", "IP", "Region", "Image")
	fmt.Fprintf(w, "\n |%s|%s|%s|%s|%s|", "------", "----", "------", "----", "----")

	for _, droplet := range droplets {
		ip, _ := droplet.PublicIPv4()

		fmt.Fprintf(w, "\n |%d|%s|%s|%s|%s|", droplet.ID, droplet.Name, ip, droplet.Region.Name, fmt.Sprintf("%s %s", droplet.Image.Distribution, droplet.Image.Name))
	}

	w.Flush()

	return p.responsef(args, buffer.String()), nil
}

func (p *Plugin) helpCommandFunc(args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	text := "###### Mattermost Digital Ocean Plugin - Slash Command Help\n" + strings.Replace(commandHelp, "|", "`", -1)
	return p.responsef(args, text), nil
}
