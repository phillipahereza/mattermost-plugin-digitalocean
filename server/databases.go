package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/digitalocean/godo"
	"github.com/mattermost/mattermost-server/v5/model"
	"text/tabwriter"
)

func (p *Plugin) listClustersFunc(args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	client, err := p.GetClient(args.UserId)
	if err != nil {
		p.API.LogError("Failed to get digitalOcean client", "Err", err.Error())
		return p.responsef(args, "Failed to get DigitalOcean client"),
			&model.AppError{Message: err.Error()}
	}

	opts := &godo.ListOptions{}

	databases, response, err := client.Databases.List(context.TODO(), opts)

	if err != nil {
		p.API.LogError("failed to fetch databases", "response", response, "Err", err.Error())
		return p.responsef(args, "Error while fetching databases list"),
			&model.AppError{Message: err.Error()}
	}

	if len(databases) == 0 {
		return p.responsef(args, "You don't have any databases configured"), nil
	}

	buffer := new(bytes.Buffer)

	w := new(tabwriter.Writer)

	w.Init(buffer, 8, 8, 0, '\t', 0)
	fmt.Fprintf(w, "\n |%s|%s|%s|%s|%s|%s|", "ID", "Name", "Engine", "Status", "Size", "Region")
	fmt.Fprintf(w, "\n |%s|%s|%s|%s|%s|%s|", "------", "----", "----", "----", "----", "----")

	for _, db := range databases {

		fmt.Fprintf(w, "\n |%s|%s|%s|%s|%s|%s|", db.ID, db.Name, db.EngineSlug, db.Status, db.SizeSlug, db.RegionSlug)
	}

	w.Flush()

	return p.responsef(args, buffer.String()), nil
}
