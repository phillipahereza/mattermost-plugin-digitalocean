package main

import (
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

func (p *Plugin) getSlackAttachment(message string) []*model.SlackAttachment {
	var fields []*model.SlackAttachmentField

	fields = append(fields, &model.SlackAttachmentField{
		Title: "DigitalOcean monitoring triggered",
		Value: message,
		Short: true,
	})
	return []*model.SlackAttachment{
		{
			Color:  "#95b7d0",
			Text:   "",
			Fields: fields,
		},
	}
}

// PostToChannels is
func (p *Plugin) PostToChannels() {

	alerts, err := p.retrieveDOEmailAlerts()
	if err != nil {
		p.API.LogError("Failed to retrieve email alerts", "Err", err.Error())
	}

	if len(alerts) > 0 {
		subcription, _ := p.store.LoadSubscription()
		channels := subcription.Channels

		for _, alert := range alerts {
			for _, channel := range channels {
				attachment := p.getSlackAttachment(alert)
				post := &model.Post{
					UserId:    p.BotUserID,
					ChannelId: channel,
				}
				post.AddProp("attachments", attachment)

				p.API.CreatePost(post)
			}
		}
	}
}
