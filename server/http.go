package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/digitalocean/godo"
	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
	"golang.org/x/net/context"
)

const (
	routeToDOApps        = "/api/v1/do-api-apps"
	routeToCreateDroplet = "/api/v1/create-droplet"
)

func (p *Plugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	switch path {
	case routeToDOApps:
		p.httpRouteToDOApps(w, r)
	case routeToCreateDroplet:
		p.httpRouteCreateDroplet(w, r)
	default:
		http.NotFound(w, r)
	}
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

	createDropletRequest := &godo.DropletCreateRequest{}
	err := json.NewDecoder(r.Body).Decode(&createDropletRequest)
	if err != nil {
		return
	}

	client, err := p.GetClient(mattermostUserID)
	if err != nil {
		return
	}

	droplet, _, err := client.CreateDroplet(context.Background(), createDropletRequest)

	post := &model.Post{
		Message: fmt.Sprintf("New droplet: %s created by NAME", droplet.Name),
	}

	p.API.CreatePost(post)
}
