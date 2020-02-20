package main

import (
	"context"

	"github.com/digitalocean/godo"
	"github.com/mattermost/mattermost-server/v5/model"
	cron "github.com/robfig/cron/v3"
)

// RunPollingJobs is
func (p *Plugin) RunPollingJobs() {
	c := cron.New()
	c.AddFunc("*/1 * * * *", func() {
		p.PostToChannels()
	})

	c.Start()
}

// PollForDroplets is
func (p *Plugin) PollForDroplets() ([]godo.Droplet, error) {
	client, err := p.GetClient("pass_user_id")
	if err != nil {
		return []godo.Droplet{}, err
	}
	droplets, _, err := client.ListDroplets(context.Background(), &godo.ListOptions{})
	if err != nil {
		return []godo.Droplet{}, err
	}
	return droplets, nil
}

func checkDroplets(droplets []godo.Droplet) *model.Post {
	var inactiveDroplets []godo.Droplet

	if len(droplets) == 0 {
		// TODO: Compose better message
		return &model.Post{}
	}

	for _, droplet := range droplets {
		if droplet.Status == "active" {
			continue
		}
		inactiveDroplets = append(inactiveDroplets, droplet)
	}

	// TODO: Compose post with information after the polling
	return &model.Post{}
}

// PostToChannels is
func (p *Plugin) PostToChannels() {
	subcription, _ := p.store.LoadSubscription()
	channels := subcription.Channels
	for _, channel := range channels {
		post := &model.Post{
			UserId:    p.BotUserID,
			Message:   "I am bot NS",
			ChannelId: channel,
		}

		p.API.SendEphemeralPost("mijye5mke7dbdbi8fb3b3a8bwh", post)
	}
}
