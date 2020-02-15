package main

import (
	"context"

	digitalocean "github.com/digitalocean/godo"
	"golang.org/x/oauth2"
)

// TokenSource contains the access token used for authentication
type TokenSource struct {
	AccessToken string
}

// Token returns the access token
func (t *TokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}

// GetClient returns a digital ocean client with configured token
func (p *Plugin) GetClient(mmUser string) (*digitalocean.Client, error) {
	token, err := p.store.LoadUserDOToken(mmUser)
	if err != nil {
		p.API.LogError("Failed to load DO token", "user", mmUser, "Err", err.Error())
		return nil, err
	}

	tokenSource := &TokenSource{
		AccessToken: token,
	}

	oauthClient := oauth2.NewClient(context.Background(), tokenSource)
	client := digitalocean.NewClient(oauthClient)
	return client, nil
}
