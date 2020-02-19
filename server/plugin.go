package main

import (
	"context"
	"github.com/digitalocean/godo"
	"golang.org/x/oauth2"
	"path/filepath"
	"sync"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
	"github.com/phillipahereza/mattermost-plugin-digitalocean/server/client"
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
		p.API.LogError("Failed to ensure digitalOcean bot", "Err", err.Error())
		return errors.Wrap(err, "failed to ensure digitalOcean bot")
	}
	p.BotUserID = botID
	store := CreateStore(p)
	p.store = store

	// start jobs
	p.RunJobs()

	return nil
}

func getCommand() *model.Command {
	return &model.Command{
		Trigger:          "do",
		DisplayName:      "do",
		Description:      "Integration with DigitalOcean.",
		AutoComplete:     true,
		AutoCompleteDesc: "Available commands: help",
		AutoCompleteHint: "[command]",
	}
}

// GetClient returns a digital ocean client with configured token
func (p *Plugin) GetClient(mmUser string) (*client.DigitalOceanClient, error) {
	token, err := p.store.LoadUserDOToken(mmUser)
	if err != nil {
		p.API.LogError("Failed to load DO token", "user", mmUser, "Err", err.Error())
		return nil, err
	} else if token == "" {
		p.API.LogError("Failed to load DO token", "user", mmUser, "Err", err)
		return nil, errors.New("Missing DigitalOcean token: User `/do token` to get instructions on how to add a token")
	}

	tokenSource := &client.TokenSource{
		AccessToken: token,
	}

	oauthClient := oauth2.NewClient(context.Background(), tokenSource)
	godoClient := godo.NewClient(oauthClient)
	return &client.DigitalOceanClient{Client: godoClient}, nil
}
