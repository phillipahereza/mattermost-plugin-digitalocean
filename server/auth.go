package main

import (
	"github.com/mattermost/mattermost-server/v5/model"
	"strings"
)

func (p *Plugin) isUserAuthorized(id string) (bool, *model.AppError) {
	user, appErr := p.API.GetUser(id)
	if appErr != nil {
		p.API.LogError("User not authorized", "user_id", id, "Err", appErr.Error())
		return false, appErr
	}
	if !strings.Contains(user.Roles, "system_admin") {
		return false, nil
	}

	return true, nil
}

func (p *Plugin) connectCommandFunc(args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	userID := args.UserId

	token := extractTokenFromCommand(args.Command)
	stErr := p.store.StoreUserDOToken(token, userID)
	if stErr != nil {
		p.API.LogError("Failed to store token", "user_id", args.UserId, "Err", stErr.Error())
		return p.responsef(args, "Failed to store token. Contact system admin"),
			&model.AppError{Message: stErr.Error()}
	}

	return p.responsef(args, "Successfully added a connecting token"), nil
}

func (p *Plugin) disconnectCommandFunc(args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	err := p.store.DeleteUserDOToken(args.UserId)
	if err != nil {
		p.API.LogError("Failed to remove the user token", "user_id", args.UserId, "Err", err.Error())
		return p.responsef(args, "Failed to remove the user token"),
			&model.AppError{Message: err.Error()}
	}

	return p.responsef(args, "Successfully disconnected token"), nil
}
