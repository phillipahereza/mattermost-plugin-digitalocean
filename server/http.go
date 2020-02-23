package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"context"

	"github.com/digitalocean/godo"
	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
)

const (
	routeToDOApps           = "/api/v1/do-api-apps"
	routeToCreateDroplet    = "/api/v1/create-droplet"
	routeToListRegions      = "/api/v1/get-do-regions"
	routeToListDropletSizes = "/api/v1/get-do-sizes"
	routeToListImages       = "/api/v1/get-do-images"
)

func (p *Plugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	switch path {
	case routeToDOApps:
		p.httpRouteToDOApps(w, r)
	case routeToCreateDroplet:
		p.httpRouteCreateDroplet(w, r)
	case routeToListRegions:
		p.httpRouteToListRegions(w, r)
	case routeToListDropletSizes:
		p.httpRouteToListDropletSizes(w, r)
	case routeToListImages:
		p.httpRouteToListImages(w, r)
	default:
		http.NotFound(w, r)
	}
}

func writeError(w http.ResponseWriter, err error) {
	errBytes, _ := json.Marshal(err)
	w.Write(errBytes)
}

func (p *Plugin) httpRouteToDOApps(w http.ResponseWriter, r *http.Request) {
	tID := p.getConfiguration().DOTeamID
	doAppsPage := fmt.Sprintf("https://cloud.digitalocean.com/account/api/tokens?i=%s", tID)
	http.Redirect(w, r, doAppsPage, http.StatusSeeOther)
}

func (p *Plugin) httpRouteCreateDroplet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}

	mattermostUserID := r.Header.Get("Mattermost-User-Id")
	if mattermostUserID == "" {
		return
	}

	createDropletRequest := godo.DropletCreateRequest{}
	err := json.NewDecoder(r.Body).Decode(&createDropletRequest)
	if err != nil {
		return
	}

	client, err := p.GetClient(mattermostUserID)
	if err != nil {
		return
	}

	droplet, _, e := client.CreateDroplet(context.Background(), &createDropletRequest)
	if e != nil {
		return
	}

	subcription, _ := p.store.LoadSubscription()
	channels := subcription.Channels

	user, _ := p.API.GetUser(mattermostUserID)

	msg := fmt.Sprintf("New droplet: [%s](%s) has been created by %s", droplet.Name, getDropletURL(droplet.ID), user.Username)

	for _, channel := range channels {
		post := &model.Post{
			UserId:    p.BotUserID,
			ChannelId: channel,
			Message:   msg,
		}

		p.API.CreatePost(post)
	}
}

func (p *Plugin) httpRouteToListRegions(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		return
	}

	mattermostUserID := r.Header.Get("Mattermost-User-Id")
	if mattermostUserID == "" {
		return
	}

	client, err := p.GetClient(mattermostUserID)
	if err != nil {
		return
	}

	regions, _, err := client.ListRegions(context.Background(), &godo.ListOptions{})
	if err != nil {
		writeError(w, err)
	}

	data, _ := json.Marshal(regions)
	w.Write(data)
}

func (p *Plugin) httpRouteToListDropletSizes(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		return
	}

	mattermostUserID := r.Header.Get("Mattermost-User-Id")
	if mattermostUserID == "" {
		return
	}

	client, err := p.GetClient(mattermostUserID)
	if err != nil {
		return
	}

	sizes, _, err := client.ListSizes(context.Background(), &godo.ListOptions{})
	if err != nil {
		writeError(w, err)
	}

	data, _ := json.Marshal(sizes)
	w.Write(data)
}

func (p *Plugin) httpRouteToListImages(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		return
	}

	mattermostUserID := r.Header.Get("Mattermost-User-Id")
	if mattermostUserID == "" {
		return
	}

	client, err := p.GetClient(mattermostUserID)
	if err != nil {
		return
	}

	images, _, err := client.ListImages(context.Background(), &godo.ListOptions{})
	if err != nil {
		writeError(w, err)
	}

	data, _ := json.Marshal(images)
	w.Write(data)
}
