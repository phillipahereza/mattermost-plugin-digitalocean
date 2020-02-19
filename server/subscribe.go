package main

import "github.com/mattermost/mattermost-server/v5/model"

const subscriptionKVKey = "subs"

//Subscription is type that keeps a record of all channels
// that'll receive updates from the DO bot.
type Subscription struct {
	Channels []string
}

func (p *Plugin) subscribeCommandFunc(args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	channelID := args.ChannelId
	stErr := p.store.AddChannelToSubcription(channelID)
	if stErr != nil {
		p.API.LogError("Failed to add channel to subcriptions", "user_id", channelID, "Err", stErr.Error())
		return p.responsef(args, stErr.Error()),
			&model.AppError{Message: stErr.Error()}
	}

	return p.responsef(args, "Channel successfully subscribed!"), nil
}

func (p *Plugin) unsubscribeCommandFunc(args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	channelID := args.ChannelId
	stErr := p.store.RemoveChannelFromSubscription(channelID)
	if stErr != nil {
		p.API.LogError("Failed to remove channel from subcriptions", "user_id", channelID, "Err", stErr.Error())
		return p.responsef(args, "Failed to remove channel from subscription list."),
			&model.AppError{Message: stErr.Error()}
	}

	return p.responsef(args, "This channel will nolonger recieve updates from DO bot"), nil
}
