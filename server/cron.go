package main

import (
	"context"
	"fmt"

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

func (p *Plugin) getSlackAttachment(inactive []godo.Droplet) []*model.SlackAttachment {
	var fields []*model.SlackAttachmentField

	if len(inactive) > 0 {
		fields = append(fields, &model.SlackAttachmentField{
			Title: "Droplet regular check",
			Value: fmt.Sprintf("%d were found inactive", len(inactive)),
			Short: true,
		})

	} else {
		fields = append(fields, &model.SlackAttachmentField{
			Title: "Droplet regular check",
			Value: "Add a message here. All droplets alive and kickin'",
			Short: true,
		})
	}
	return []*model.SlackAttachment{
		{
			Color:  "#95b7d0",
			Text:   "",
			Fields: fields,
		},
	}
}

// PollForDroplets is
func (p *Plugin) PollForDroplets() ([]godo.Droplet, error) {
	client, err := p.GetClient("mijye5mke7dbdbi8fb3b3a8bwh")
	if err != nil {
		return []godo.Droplet{}, err
	}
	droplets, _, err := client.ListDroplets(context.Background(), &godo.ListOptions{})
	if err != nil {
		return []godo.Droplet{}, err
	}
	return droplets, nil
}

func (p *Plugin) checkDroplets(droplets []godo.Droplet) (ok bool, inactive []godo.Droplet) {
	var inactiveDroplets []godo.Droplet

	if len(droplets) == 0 {
		return true, []godo.Droplet{}
	}

	for _, droplet := range droplets {
		if droplet.Status == "active" {
			continue
		}
		inactiveDroplets = append(inactiveDroplets, droplet)
	}

	if len(inactiveDroplets) > 0 {
		return false, inactiveDroplets
	}

	return true, []godo.Droplet{}
}

// PostToChannels is
func (p *Plugin) PostToChannels() {
	subcription, _ := p.store.LoadSubscription()
	droplets, _ := p.PollForDroplets()
	_, inactive := p.checkDroplets(droplets)
	channels := subcription.Channels
	for _, channel := range channels {
		attachment := p.getSlackAttachment(inactive)
		post := &model.Post{
			UserId:    p.BotUserID,
			ChannelId: channel,
		}
		post.AddProp("attachments", attachment)

		p.API.SendEphemeralPost("mijye5mke7dbdbi8fb3b3a8bwh", post)
	}
}
