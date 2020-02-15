package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/digitalocean/godo"
	"github.com/mattermost/mattermost-server/v5/model"
	"text/tabwriter"
)

func (p *Plugin) listDomainsFunc(args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	client, err := p.GetClient(args.UserId)
	if err != nil {
		p.API.LogError("Failed to get digitalOcean client", "Err", err.Error())
		return p.responsef(args, "Failed to get DigitalOcean client"),
			&model.AppError{Message: err.Error()}
	}

	opts := &godo.ListOptions{}

	domains, response, err := client.Domains.List(context.TODO(), opts)

	if err != nil {
		p.API.LogError("failed to fetch domains", "response", response, "Err", err.Error())
		return p.responsef(args, "Error while fetching domains list"),
			&model.AppError{Message: err.Error()}
	}

	if len(domains) == 0 {
		return p.responsef(args, "You don't have any domains configured. Use `/do create-domain  <domainName> <ipAddress[optional]>` to provision one"), nil
	}

	buffer := new(bytes.Buffer)

	w := new(tabwriter.Writer)

	w.Init(buffer, 8, 8, 0, '\t', 0)
	fmt.Fprintf(w, "\n |%s|%s|", "Name", "TTL")
	fmt.Fprintf(w, "\n |%s|%s|", "------", "----")

	for _, domain := range domains {

		fmt.Fprintf(w, "\n |%s|%d|", domain.Name, domain.TTL)
	}

	w.Flush()

	return p.responsef(args, buffer.String()), nil
}
