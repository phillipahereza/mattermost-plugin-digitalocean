package main

import (
	"context"
	"fmt"

	"github.com/digitalocean/godo"
	"github.com/mattermost/mattermost-server/v5/model"
	cron "github.com/robfig/cron/v3"
)

// DOCron is
type DOCron struct {
	cron *cron.Cron
}

func newCron() *cron.Cron {
	c := cron.New()
	return c
}

// RegisterPollingJobs is
func (p *Plugin) RegisterPollingJobs() {
	cronConfig := p.getConfiguration().CronConfig
	p.cron.AddFunc(cronConfig, func() {
		p.PostToChannels()
	})
}

// StartCronJobs starts our cron jobs
func (p *Plugin) StartCronJobs() {
	p.RegisterPollingJobs()
	p.cron.Start()
}

// StopCronJobs stops cron jobs
func (p *Plugin) StopCronJobs() {
	p.cron.Stop()
}

func getDropletURL(id int) string {
	return fmt.Sprintf("https://cloud.digitalocean.com/droplets/%d", id)
}

func (p *Plugin) getSlackAttachment(inactives []godo.Droplet) []*model.SlackAttachment {
	var fields []*model.SlackAttachmentField
	title := "Droplet regular check"
	var inactiveSummary string

	if len(inactives) > 0 {
		for _, droplet := range inactives {
			inactiveSummary = inactiveSummary + fmt.Sprintf(`- [%s](%s)`, droplet.Name, getDropletURL(droplet.ID)) + "\n"
		}

		fields = append(fields, &model.SlackAttachmentField{
			Title: "Inactive droplets report",
			Value: inactiveSummary,
			Short: true,
		})

	} else {
		fields = append(fields, &model.SlackAttachmentField{
			Title: title,
			Value: "All droplets alive and kickin'",
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
	client, err := p.GetClient(adminKVKey)
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

		p.API.CreatePost(post)
	}
}
