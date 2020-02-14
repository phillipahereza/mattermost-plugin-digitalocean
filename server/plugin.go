package main

import (
	"path/filepath"
	"sync"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
	"github.com/pkg/errors"
)

// Plugin implements the interface expected by the Mattermost server to communicate between the server and plugin processes.
type Plugin struct {
	plugin.MattermostPlugin

	// configurationLock synchronizes access to the configuration.
	configurationLock sync.RWMutex

	// configuration is the active plugin configuration. Consult getConfiguration and
	// setConfiguration for usage.
	configuration *configuration

	store Store

	BotUserID string
}

// OnActivate is
func (p *Plugin) OnActivate() error {
	p.API.RegisterCommand(getCommand())

	profileImage := filepath.Join("assets", "do.png")

	botID, err := p.Helpers.EnsureBot(&model.Bot{
		Username:    "do",
		DisplayName: "DO",
		Description: "Created by the DigitalOcean plugin.",
	}, plugin.ProfileImagePath(profileImage))

	if err != nil {
		return errors.Wrap(err, "failed to ensure digital ocean bot")
	}
	p.BotUserID = botID
	store := CreateStore(p)
	p.store = store

	return nil
}
