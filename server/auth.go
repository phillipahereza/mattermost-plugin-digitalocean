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

func (p *Plugin) getSiteURL() string {
	url := p.API.GetConfig().ServiceSettings.SiteURL
	if url == nil {
		return ""
	}

	return *url
}

func (p *Plugin) connectCommandFunc(args *model.CommandArgs, token string) (*model.CommandResponse, *model.AppError) {
	siteURL := p.getSiteURL()
	if siteURL == "" {
		errMsg := "Please set a SITEURL"
		p.API.LogError(errMsg, "user_id", args.UserId, "Err", errMsg)
		return p.responsef(args, errMsg),
			&model.AppError{Message: errMsg}
	}
	userID := args.UserId
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
