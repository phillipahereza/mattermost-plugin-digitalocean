package main

import (
	"github.com/mattermost/mattermost-server/v5/model"
)

type Subscription struct {
	name string
	channels []string
	events []string
}

func (p *Plugin) subscribeCommandFunc(args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	channelID := args.ChannelId

}