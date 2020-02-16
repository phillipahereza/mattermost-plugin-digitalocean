package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/digitalocean/godo"
	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/phillipahereza/mattermost-plugin-digitalocean/server/client"
	"text/tabwriter"
)

func (p *Plugin) listDomainsCommandFunc(client *client.DigitalOceanClient, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {

	opts := &godo.ListOptions{}

	domains, response, err := client.ListDomains(context.TODO(), opts)

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
