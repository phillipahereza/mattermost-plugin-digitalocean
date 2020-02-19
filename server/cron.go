package main

import (
	"github.com/mattermost/mattermost-server/v5/model"
	cron "github.com/robfig/cron/v3"
)

// RunJobs is
func (p *Plugin) RunJobs() {
	c := cron.New()
	c.AddFunc("*/1 * * * *", func() {
		p.PostToChannels()
	})

	c.Start()
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
