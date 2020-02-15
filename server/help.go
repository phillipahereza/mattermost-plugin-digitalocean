package main

import (
	"github.com/mattermost/mattermost-server/v5/model"
	"strings"
)

const commandHelp = `* |/do help| - Run 'test' to see if you're configured to run bamboo commands
* |/do connect <access token>| - Associates your DO team personal token with your mattermost account
* |/do token| - Provides instructions on getting a personal access token for the configured DigitalOcean team
* |/do show-configured-token| - Display your configured access token
* |/do list-droplets| - List all Droplets in your team
* |/do rename-droplet <dropletID> <name>| - Rename a droplet
`

func (p *Plugin) helpCommandFunc(args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	text := "###### Mattermost DigitalOcean Plugin - Slash Command Help\n" + strings.Replace(commandHelp, "|", "`", -1)
	return p.responsef(args, text), nil
}
