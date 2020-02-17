package main

import (
	"fmt"
	"github.com/mattermost/mattermost-server/v5/model"
	"strings"
)

func (p *Plugin) getPersonalTokenCommandFunc(args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	const responseMessage = `Click the link below and follow subsequent steps

1. [DigitalOcean Apps and APIs](/plugins/com.mattermost.digitalocean/%s)
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

func (p *Plugin) showConnectTokenCommandFunc(args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	tk, err := p.store.LoadUserDOToken(args.UserId)
	if err != nil {
		p.API.LogError("Failed to show connected token", "Err", err.Error())
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
