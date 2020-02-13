package main

import (
	"fmt"
	"net/http"

	"github.com/mattermost/mattermost-server/v5/plugin"
)

const (
	routeToDOApps = "/api/v1/do-api-apps"
)

func (p *Plugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	switch path {
	case routeToDOApps:
		p.httpRouteToDOApps(w, r)
	default:
		http.NotFound(w, r)
	}
}

func (p *Plugin) httpRouteToDOApps(w http.ResponseWriter, r *http.Request) {
	tID := p.getConfiguration().DOTeamID
	doAppsPage := fmt.Sprintf("https://cloud.digitalocean.com/account/api/tokens?i=%s", tID)
	http.Redirect(w, r, doAppsPage, http.StatusSeeOther)
}
