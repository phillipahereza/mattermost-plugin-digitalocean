package main

import (
	"github.com/mattermost/mattermost-server/v5/model"
	cron "github.com/robfig/cron/v3"
)

// RunJobs is
func (p *Plugin) RunJobs() {
	c := cron.New()

	post := &model.Post{
		UserId:    p.BotUserID,
		Message:   "I am bot NS",
		ChannelId: "43gik9wr17nzijea5h655jxxiy",
	}

	c.AddFunc("*/1 * * * *", func() {
		p.API.SendEphemeralPost("mijye5mke7dbdbi8fb3b3a8bwh", post)
	})

	c.Start()

}
