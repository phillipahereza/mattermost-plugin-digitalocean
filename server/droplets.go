package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/digitalocean/godo"
	"github.com/mattermost/mattermost-server/v5/model"
	"text/tabwriter"
)

func (p *Plugin) listDropletsFunc(args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	client, err := p.GetClient(args.UserId)
	if err != nil {
		p.API.LogError("Failed to get digitalOcean client", "Err", err.Error())
		return p.responsef(args, "Failed to get DigitalOcean client"),
			&model.AppError{Message: err.Error()}
	}

	opts := &godo.ListOptions{}

	droplets, response, err := client.Droplets.List(context.TODO(), opts)

	if err != nil {
		p.API.LogError("failed to fetch droplets", "response", response, "Err", err.Error())
		return p.responsef(args, "Error while fetching droplets list"),
			&model.AppError{Message: err.Error()}
	}

	if len(droplets) == 0 {
		return p.responsef(args, "You don't have any droplets configured. Use `/do droplet create` to provision one"), nil
	}

	buffer := new(bytes.Buffer)

	w := new(tabwriter.Writer)

	w.Init(buffer, 8, 8, 0, '\t', 0)
	fmt.Fprintf(w, "\n |%s|%s|%s|%s|%s|", "ID", "Name", "IP", "Region", "Image")
	fmt.Fprintf(w, "\n |%s|%s|%s|%s|%s|", "------", "----", "------", "----", "----")

	for _, droplet := range droplets {
		ip, _ := droplet.PublicIPv4()

		fmt.Fprintf(w, "\n |%d|%s|%s|%s|%s|", droplet.ID, droplet.Name, ip, droplet.Region.Name, fmt.Sprintf("%s %s", droplet.Image.Distribution, droplet.Image.Name))
	}

	w.Flush()

	return p.responsef(args, buffer.String()), nil
}

func (p *Plugin) rebootDropletFunc(args *model.CommandArgs, dropletID int) (*model.CommandResponse, *model.AppError) {
	client, err := p.GetClient(args.UserId)
	if err != nil {
		return p.responsef(args, "Failed to get DigitalOcean client"),
			&model.AppError{Message: err.Error()}
	}
	action, response, err := client.DropletActions.Reboot(context.TODO(), dropletID)

	if err != nil {
		p.API.LogError("Failed to reboot droplet", "dropletID", dropletID, "response", response, "Err", err.Error())
		return p.responsef(args, "Failed to reboot droplet %d", dropletID),
			&model.AppError{Message: err.Error()}
	}

	return p.responsef(args, "Rebooting started at: %s with status `%s`", action.StartedAt.String(), action.Status), nil
}
